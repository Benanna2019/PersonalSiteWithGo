package helpers

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"

	// "github.com/alecthomas/chroma/formatters/html" // for syntax highlighting
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

// FrontMatter represents your markdown metadata
type FrontMatter struct {
    ID          string   `yaml:"_id"`
    Title       string   `yaml:"title"`
    Published   string   `yaml:"published"`
    Slug        string   `yaml:"slug"`
    Description string   `yaml:"description"`
    Categories  []string `yaml:"categories"`
    Author      string   `yaml:"author"`
    AuthorImage string   `yaml:"authorImage"`
    Type        string   `yaml:"type"`
    CustomElementKeys []string `yaml:"customElementKeys"`
}

type PostData struct {
    Href        string      `json:"href"`
    Frontmatter FrontMatter `json:"frontmatter"`
}

type PostWithContent struct {
    Href        string      `json:"href"`
    Frontmatter FrontMatter `json:"frontmatter"`
    Html        string      `json:"html"`
}

func GetPostById(id string) (*PostWithContent, error) {
    // Ensure the id has .md extension
    if !strings.HasSuffix(id, ".md") {
        id = id + ".md"
    }

    // Try to read the specific file directly
    filePath := filepath.Join("content", id)
    meta, content, err := ParseMarkdownFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("post not found: %w", err)
    }

    // Create href without .md extension
    baseName := strings.TrimSuffix(id, ".md")
    href := fmt.Sprintf("/content/%s", baseName)
    
    return &PostWithContent{
        Href:        href,
        Frontmatter: *meta,
        Html:        content,
    }, nil
}

func ParseMarkdownFile(filename string) (*FrontMatter, string, error) {
    // Read the markdown file
    content, err := os.ReadFile(filename)
    if err != nil {
        return nil, "", fmt.Errorf("reading file: %w", err)
    }

    // Parse frontmatter
    var meta FrontMatter
    err = frontmatter.YAML.Unmarshal(content, &meta)
    if err != nil {
        return nil, "", fmt.Errorf("parsing frontmatter: %w", err)
    }

    // Setup Markdown parser with all the features
    md := goldmark.New(
        goldmark.WithExtensions(
            extension.GFM, 
            extension.Typographer,
            highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
			&frontmatter.Extender{
				Mode: frontmatter.SetMetadata,
			},
        ),
        goldmark.WithParserOptions(
            parser.WithAutoHeadingID(),
        ),
        goldmark.WithRendererOptions(
            html.WithHardWraps(),
        ),
    )

     // Create a new parser context to store metadata
    context := parser.NewContext()
    var buf bytes.Buffer
    if err := md.Convert(content, &buf, parser.WithContext(context)); err != nil {
        return nil, "", fmt.Errorf("converting markdown: %w", err)
    }

    // Extract metadata from context
    if metadata := frontmatter.Get(context); metadata != nil {
        if err := metadata.Decode(&meta); err != nil {
            return nil, "", fmt.Errorf("decoding frontmatter: %w", err)
        }
    }

    return &meta, buf.String(), nil
}

// Usage in a route handler
func GetMarkdownPosts() []struct {
    Meta    FrontMatter
    Content string
} {
    posts := make([]struct {
        Meta    FrontMatter
        Content string
    }, 0)

    // Read all markdown files from a directory
    files, err := os.ReadDir("content")
    if err != nil {
        return posts
    }

    for _, file := range files {
        if !strings.HasSuffix(file.Name(), ".md") {
            continue
        }

        meta, content, err := ParseMarkdownFile(filepath.Join("content", file.Name()))
        if err != nil {
            continue
        }

        posts = append(posts, struct {
            Meta    FrontMatter
            Content string
        }{
            Meta:    *meta,
            Content: content,
        })
    }

    // Sort posts by date if needed
    sort.Slice(posts, func(i, j int) bool {
		date1, _ := time.Parse("January 2, 2006", posts[i].Meta.Published)
		date2, _ := time.Parse("January 2, 2006", posts[j].Meta.Published)

        return date1.Unix() > date2.Unix()
    })

	fmt.Println("posts", posts)

    return posts
}

func GetPostMetadata() ([]PostData, error) {
  var posts []PostData
    
    // Read all markdown files from the directory
    files, err := os.ReadDir("content")
    if err != nil {
        return nil, fmt.Errorf("reading directory: %w", err)
    }

    for _, file := range files {
        if !strings.HasSuffix(file.Name(), ".md") {
            continue
        }

        meta, _, err := ParseMarkdownFile(filepath.Join("content", file.Name()))
        if err != nil {
            continue
        }

        // Create the href by removing .md extension
        baseName := strings.TrimSuffix(file.Name(), ".md")
        href := fmt.Sprintf("/posts/%s", baseName)

        posts = append(posts, PostData{
            Href:        href,
            Frontmatter: *meta,
        })
    }

    // Sort by date with newest first
    sort.Slice(posts, func(i, j int) bool {
        date1, _ := time.Parse("January 2, 2006", posts[i].Frontmatter.Published)
		date2, _ := time.Parse("January 2, 2006", posts[j].Frontmatter.Published)

        return date1.Unix() > date2.Unix()
    })

    return posts, nil
}

func GetPaginatedPosts(limit, offset int) ([]PostData, int, error) {
    allPosts, err := GetPostMetadata()
    if err != nil {
        return nil, 0, err
    }

    return allPosts, len(allPosts), nil
}
