package main

import (
	"context"
	"fmt"
	"log/slog"

	"embed"

	"net/http"
	"os"

	"bytes"
	"encoding/json"
	"path/filepath"

	extism "github.com/extism/go-sdk"
	"github.com/zangster300/northstar/web/components"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zangster300/northstar/routes"
	"golang.org/x/sync/errgroup"
)

const port = 8080

//go:embed web/custom-elements
var customElements embed.FS

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(fmt.Sprintf("Starting Server @:%d", port))
	// defer logger.Info("Stopping Server")

	markup := "<site-layout></site-layout>"
	elements := readElementsFromEmbed(customElements)
	fmt.Println("elements", elements["site-layout"])
	initialState := make(map[string]interface{})
	data := map[string]interface{}{
		"markup":       markup,
		"elements":     elements,
		"initialState": initialState,
	}
	payload, _ := Marshal(data)
	rendered, _ := render(payload)
	fmt.Println("rendered", rendered)
	var result map[string]interface{}
	if err := json.Unmarshal(rendered, &result); err != nil {
		fmt.Println("Failed to parse rendered output", err)
		return
	}

	element_body, ok := result["body"].(string)
	if !ok {
		fmt.Println("Rendered document is not a string")
		return
	}

	// fmt.Println("element_body", element_body)

	router := chi.NewMux()
	router.Use(middleware.Logger)

	router.Handle("/static/*", http.StripPrefix("/static/", static(logger)))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		components.DummySiteLayout(element_body).Render(r.Context(), w)
	})

	http.ListenAndServe(":8080", router)

	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer stop()

	// if err := run(ctx, logger); err != nil {
	// 	logger.Error("Error running server", slog.Any("err", err))
	// 	os.Exit(1)
	// }
}

func readElementsFromEmbed(fs embed.FS) map[string]string {
    elements := make(map[string]string)
    entries, err := fs.ReadDir("web/custom-elements")
    if err != nil {
        fmt.Printf("Error reading embedded directory: %s\n", err)
        return elements
    }

    for _, entry := range entries {
        if !entry.IsDir() {
            content, err := fs.ReadFile("web/custom-elements/" + entry.Name())
            if err != nil {
                fmt.Printf("Error reading embedded file %s: %s\n", entry.Name(), err)
                continue
            }
            key := filepath.Base(entry.Name())
            ext := filepath.Ext(key)
            keyWithoutExt := key[:len(key)-len(ext)]
            elements[keyWithoutExt] = string(content)
        }
    }
    return elements
}

func run(ctx context.Context, logger *slog.Logger) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(startServer(ctx, logger, port))

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error running server: %w", err)
	}

	return nil
}

func startServer(ctx context.Context, logger *slog.Logger, port int) func() error {
	return func() error {
		router := chi.NewMux()

		router.Use(
			middleware.Logger,
			middleware.Recoverer,
		)

		router.Handle("/static/*", http.StripPrefix("/static/", static(logger)))

		cleanup, err := routes.SetupRoutes(ctx, logger, router)
		defer cleanup()
		if err != nil {
			return fmt.Errorf("error setting up routes: %w", err)
		}

		srv := &http.Server{
			Addr:    fmt.Sprintf("localhost:%d", port),
			Handler: router,
		}

		go func() {
			<-ctx.Done()
			srv.Shutdown(context.Background())
		}()

		return srv.ListenAndServe()
	}
}

// Enhance Element Helpers

func Marshal(i interface{}) ([]byte, error) {
    buffer := &bytes.Buffer{}
    encoder := json.NewEncoder(buffer)
    encoder.SetEscapeHTML(false)
    encoder.SetIndent("", "    ") 
    err := encoder.Encode(i)
    if err != nil {
        return nil, err 
    }
    // Trim the trailing newline added by Encode
    return bytes.TrimRight(buffer.Bytes(), "\n"), nil
}

func readElements(directory string) map[string]string {
	elements := make(map[string]string)
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err)
		return elements
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(directory, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading file %s: %s\n", file.Name(), err)
				continue
			}
			key := filepath.Base(filePath)
			ext := filepath.Ext(key)
			keyWithoutExt := key[:len(key)-len(ext)]
			elements[keyWithoutExt] = string(content)
		}
	}
	return elements
}

func render(payload []byte) ([]byte, error) {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
      extism.WasmFile{
				Path: "./wasm/enhance-ssr.wasm",
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{
    EnableWasi: true,
  }
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize plugin: %v", err)
	}

	exit, out, err := plugin.Call("ssr", payload)
	if err != nil {
		return nil, fmt.Errorf("plugin call failed: %v, exit code: %d", err, exit)
	}

	return out, nil
}

