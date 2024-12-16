package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"embed"

	datastar "github.com/starfederation/datastar/code/go/sdk"

	"github.com/delaneyj/toolbelt"
	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/samber/lo"
	"github.com/zangster300/northstar/helpers"
	"github.com/zangster300/northstar/web/components"
	"github.com/zangster300/northstar/web/layouts"
)

func setupIndexRoute(router chi.Router, store sessions.Store, ns *embeddednats.Server, customElements embed.FS) error {
	nc, err := ns.Client()
	if err != nil {
		return fmt.Errorf("error creating nats client: %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return fmt.Errorf("error creating jetstream client: %w", err)
	}

	kv, err := js.CreateOrUpdateKeyValue(context.Background(), jetstream.KeyValueConfig{
		Bucket:      "todos",
		Description: "Datastar Todos",
		Compression: true,
		TTL:         time.Hour,
		MaxBytes:    16 * 1024 * 1024,
	})

	if err != nil {
		return fmt.Errorf("error creating key value: %w", err)
	}

	saveMVC := func(ctx context.Context, sessionID string, mvc *components.TodoMVC) error {
		b, err := json.Marshal(mvc)
		if err != nil {
			return fmt.Errorf("failed to marshal mvc: %w", err)
		}
		if _, err := kv.Put(ctx, sessionID, b); err != nil {
			return fmt.Errorf("failed to put key value: %w", err)
		}
		return nil
	}

	resetMVC := func(mvc *components.TodoMVC) {
		mvc.Mode = components.TodoViewModeAll
		mvc.Todos = []*components.Todo{
			{Text: "Learn a backend language", Completed: true},
			{Text: "Learn Datastar", Completed: false},
			{Text: "Create Hypermedia", Completed: false},
			{Text: "???", Completed: false},
			{Text: "Profit", Completed: false},
		}
		mvc.EditingIdx = -1
	}

	mvcSession := func(w http.ResponseWriter, r *http.Request) (string, *components.TodoMVC, error) {
		ctx := r.Context()
		sessionID, err := upsertSessionID(store, r, w)
		if err != nil {
			return "", nil, fmt.Errorf("failed to get session id: %w", err)
		}

		mvc := &components.TodoMVC{}
		if entry, err := kv.Get(ctx, sessionID); err != nil {
			if err != jetstream.ErrKeyNotFound {
				return "", nil, fmt.Errorf("failed to get key value: %w", err)
			}
			resetMVC(mvc)

			if err := saveMVC(ctx, sessionID, mvc); err != nil {
				return "", nil, fmt.Errorf("failed to save mvc: %w", err)
			}
		} else {
			if err := json.Unmarshal(entry.Value(), mvc); err != nil {
				return "", nil, fmt.Errorf("failed to unmarshal mvc: %w", err)
			}
		}
		return sessionID, mvc, nil
	}

	// router.Get("/rss", func(w http.ResponseWriter, r *http.Request) {
	// 	// Check if client accepts brotli compression
	// 	acceptEncoding := r.Header.Get("Accept-Encoding")
	// 	returnCompressed := strings.Contains(acceptEncoding, "br")

	// 	// Set content type header with charset
	// 	w.Header().Set("Content-Type", "application/rss+xml; charset=UTF-8")

	// 	if returnCompressed {
	// 		// Serve compressed .br file
	// 		w.Header().Set("Content-Encoding", "br")
	// 		brContent, err := os.ReadFile("web/static/rss.br")
	// 		if err != nil {
	// 			http.Error(w, "Failed to read RSS feed", http.StatusInternalServerError)
	// 			return
	// 		}
	// 		w.Write(brContent)
	// 	} else {
	// 		// Serve uncompressed .xml file
	// 		xmlContent, err := os.ReadFile("web/static/rss.xml")
	// 		if err != nil {
	// 			http.Error(w, "Failed to read RSS feed", http.StatusInternalServerError)
	// 			return
	// 		}
			
	// 		// Ensure XML has UTF-8 declaration
	// 		if !strings.HasPrefix(string(xmlContent), "<?xml") {
	// 			xmlDeclaration := []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	// 			w.Write(xmlDeclaration)
	// 		}
			
	// 		w.Write(xmlContent)
	// 	}
	// })	

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		    w.Header().Set("Cache-Control", helpers.GetCacheControl())

                // Get pagination parameters
                limit := 20
                if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
                    if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
                        limit = l
                    }
                }

                offset := 0
                if offsetStr := r.URL.Query().Get("offset"); offsetStr != "" {
                    if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
                        offset = o
                    }
                }


                paginatedPosts, total, err := helpers.GetPaginatedPosts(limit, offset)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                posts := make([]map[string]interface{}, len(paginatedPosts))
                for i, post := range paginatedPosts {
                    posts[i] = map[string]interface{}{
                        "href": post.Href,
                        "frontmatter": map[string]interface{}{
                            "id":          post.Frontmatter.ID,
                            "title":       post.Frontmatter.Title,
                            "published":   post.Frontmatter.Published,
                            "slug":        post.Frontmatter.Slug,
                            "description": post.Frontmatter.Description,
                            "categories":  post.Frontmatter.Categories,
                            "author":      post.Frontmatter.Author,
                            "authorImage": post.Frontmatter.AuthorImage,
                            "type":        post.Frontmatter.Type,
                        },
                    }
                }

                data := map[string]interface{}{
                    "posts":  posts,
                    "limit":  limit,
                    "offset": offset,
                    "total":  total,
                }

				ssr_elements := helpers.RenderSSR(customElements, "<index-page></index-page>", data)
				if ssr_elements.Error != nil {
					http.Error(w, ssr_elements.Error.Error(), http.StatusInternalServerError)
					return
				}

			components.SiteLayout(ssr_elements.Body, layouts.MetaData{
				Title:       "Home",
				PageType:    layouts.PageTypeDefault,
			}).Render(r.Context(), w)

				
	})

	// router.Get("/sw", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/javascript")
	// 	w.Write([]byte(helpers.GetSWJS()))
	// })

	router.Get("/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		    w.Header().Set("Cache-Control", helpers.GetCacheControl())

            id := chi.URLParam(r, "id")

            
            post, err := helpers.GetPostById(id)
            if err != nil {
                http.Error(w, err.Error(), http.StatusNotFound)
                return
            }

            data := map[string]interface{}{
                "post": map[string]interface{}{
                    "href": post.Href,
                    "frontmatter": map[string]interface{}{
                        "id":          post.Frontmatter.ID,
                        "title":       post.Frontmatter.Title,
                        "published":   post.Frontmatter.Published,
                        "slug":        post.Frontmatter.Slug,
                        "description": post.Frontmatter.Description,
                        "categories":  post.Frontmatter.Categories,
                        "author":      post.Frontmatter.Author,
                        "authorImage": post.Frontmatter.AuthorImage,
                        "type":        post.Frontmatter.Type,
						"components":  post.Frontmatter.CustomElementKeys,
                    },
                    "html": post.Html,
                },
            }

			site_url := helpers.GetSiteURL()

            ssr_elements := helpers.RenderSSR(customElements, "<post-page></post-page>", data)
            if ssr_elements.Error != nil {
                http.Error(w, ssr_elements.Error.Error(), http.StatusInternalServerError)
                return
            }

			components.SiteLayout(ssr_elements.Body, layouts.MetaData{
				Title:       post.Frontmatter.Title,
				Description: post.Frontmatter.Description,
				CoverImage:  "/static/images/outdoorworking.png",
				CoverWidth:  "16",
				CoverHeight: "9",
				SiteURL:     site_url,
				PageType:    layouts.PageTypePost,
			}).Render(r.Context(), w)
            // if err := sse.MergeFragmentTempl(c); err != nil {
            //     sse.ConsoleError(err)
            //     return
            // }
    })

	router.Route("/api", func(apiRouter chi.Router) {

		apiRouter.Route("/posts", func(postsRouter chi.Router) {
			postsRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {

				
			})
		})

		apiRouter.Route("/todos", func(todosRouter chi.Router) {
			todosRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {

				sessionID, mvc, err := mvcSession(w, r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				sse := datastar.NewSSE(w, r)

				// Watch for updates
				ctx := r.Context()
				watcher, err := kv.Watch(ctx, sessionID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				defer watcher.Stop()

				for {
					select {
					case <-ctx.Done():
						return
					case entry := <-watcher.Updates():
						if entry == nil {
							continue
						}
						if err := json.Unmarshal(entry.Value(), mvc); err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}
						c := components.TodosMVCView(mvc)
						if err := sse.MergeFragmentTempl(c); err != nil {
							sse.ConsoleError(err)
							return
						}
					}
				}
			})

			todosRouter.Put("/reset", func(w http.ResponseWriter, r *http.Request) {
				sessionID, mvc, err := mvcSession(w, r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				resetMVC(mvc)
				if err := saveMVC(r.Context(), sessionID, mvc); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			})

			todosRouter.Put("/cancel", func(w http.ResponseWriter, r *http.Request) {

				sessionID, mvc, err := mvcSession(w, r)
				sse := datastar.NewSSE(w, r)
				if err != nil {
					sse.ConsoleError(err)
					return
				}

				mvc.EditingIdx = -1
				if err := saveMVC(r.Context(), sessionID, mvc); err != nil {
					sse.ConsoleError(err)
					return
				}
			})

			todosRouter.Put("/mode/{mode}", func(w http.ResponseWriter, r *http.Request) {

				sessionID, mvc, err := mvcSession(w, r)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				modeStr := chi.URLParam(r, "mode")
				modeRaw, err := strconv.Atoi(modeStr)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				mode := components.TodoViewMode(modeRaw)
				if mode < components.TodoViewModeAll || mode > components.TodoViewModeCompleted {
					http.Error(w, "invalid mode", http.StatusBadRequest)
					return
				}

				mvc.Mode = mode
				if err := saveMVC(r.Context(), sessionID, mvc); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			})

			todosRouter.Route("/{idx}", func(todoRouter chi.Router) {
				routeIndex := func(w http.ResponseWriter, r *http.Request) (int, error) {
					idx := chi.URLParam(r, "idx")
					i, err := strconv.Atoi(idx)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return 0, err
					}
					return i, nil
				}

				todoRouter.Post("/toggle", func(w http.ResponseWriter, r *http.Request) {
					sessionID, mvc, err := mvcSession(w, r)

					sse := datastar.NewSSE(w, r)
					if err != nil {
						sse.ConsoleError(err)
						return
					}

					i, err := routeIndex(w, r)
					if err != nil {
						sse.ConsoleError(err)
						return
					}

					if i < 0 {
						setCompletedTo := false
						for _, todo := range mvc.Todos {
							if !todo.Completed {
								setCompletedTo = true
								break
							}
						}
						for _, todo := range mvc.Todos {
							todo.Completed = setCompletedTo
						}
					} else {
						todo := mvc.Todos[i]
						todo.Completed = !todo.Completed
					}

					saveMVC(r.Context(), sessionID, mvc)
				})

				todoRouter.Route("/edit", func(editRouter chi.Router) {
					editRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
						sessionID, mvc, err := mvcSession(w, r)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}

						i, err := routeIndex(w, r)
						if err != nil {
							return
						}

						mvc.EditingIdx = i
						saveMVC(r.Context(), sessionID, mvc)
					})

					editRouter.Put("/", func(w http.ResponseWriter, r *http.Request) {
						type Store struct {
							Input string `json:"input"`
						}
						store := &Store{}

						if err := datastar.ReadSignals(r, store); err != nil {
							http.Error(w, err.Error(), http.StatusBadRequest)
							return
						}

						if store.Input == "" {
							return
						}

						sessionID, mvc, err := mvcSession(w, r)
						if err != nil {
							http.Error(w, err.Error(), http.StatusInternalServerError)
							return
						}

						i, err := routeIndex(w, r)
						if err != nil {
							return
						}

						if i >= 0 {
							mvc.Todos[i].Text = store.Input
						} else {
							mvc.Todos = append(mvc.Todos, &components.Todo{
								Text:      store.Input,
								Completed: false,
							})
						}
						mvc.EditingIdx = -1

						saveMVC(r.Context(), sessionID, mvc)

					})
				})

				todoRouter.Delete("/", func(w http.ResponseWriter, r *http.Request) {
					i, err := routeIndex(w, r)
					if err != nil {
						return
					}

					sessionID, mvc, err := mvcSession(w, r)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					if i >= 0 {
						mvc.Todos = append(mvc.Todos[:i], mvc.Todos[i+1:]...)
					} else {
						mvc.Todos = lo.Filter(mvc.Todos, func(todo *components.Todo, i int) bool {
							return !todo.Completed
						})
					}
					saveMVC(r.Context(), sessionID, mvc)
				})
			})
		})
	})

	return nil
}

func MustJSONMarshal(v any) string {
	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func upsertSessionID(store sessions.Store, r *http.Request, w http.ResponseWriter) (string, error) {

	sess, err := store.Get(r, "connections")
	if err != nil {
		return "", fmt.Errorf("failed to get session: %w", err)
	}
	id, ok := sess.Values["id"].(string)
	if !ok {
		id = toolbelt.NextEncodedID()
		sess.Values["id"] = id
		if err := sess.Save(r, w); err != nil {
			return "", fmt.Errorf("failed to save session: %w", err)
		}
	}
	return id, nil
}
