package helpers

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"

	extism "github.com/extism/go-sdk"
)

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

// func readElements(directory string) map[string]string {
// 	elements := make(map[string]string)
// 	files, err := os.ReadDir("web/custom-elements")
// 	if err != nil {
// 		fmt.Printf("Error reading directory: %s\n", err)
// 		return elements
// 	}

// 	for _, file := range files {
// 		if !file.IsDir() {
// 			filePath := filepath.Join(directory, file.Name())
// 			content, err := os.ReadFile(filePath)
// 			if err != nil {
// 				fmt.Printf("Error reading file %s: %s\n", file.Name(), err)
// 				continue
// 			}
// 			key := filepath.Base(filePath)
// 			ext := filepath.Ext(key)
// 			keyWithoutExt := key[:len(key)-len(ext)]
// 			elements[keyWithoutExt] = string(content)
// 		}
// 	}
// 	return elements
// }


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

// EnhanceSSRResult represents the processed SSR output
type EnhanceSSRResult struct {
    Body    string
    Head    string            // for future meta/head content
    Scripts map[string]string // for future hydration scripts
    Error   error
}

// RenderSSR processes HTML components with the Enhance SSR engine
func RenderSSR(customElements embed.FS, markup string, initialState map[string]interface{}) EnhanceSSRResult {
    elements := readElementsFromEmbed(customElements)
    
    data := map[string]interface{}{
        "markup":       markup,
        "elements":     elements,
        "initialState": initialState,
    }
    
    payload, err := Marshal(data)
    if err != nil {
        return EnhanceSSRResult{Error: fmt.Errorf("marshal error: %w", err)}
    }
    
    rendered, err := render(payload)
    if err != nil {
        return EnhanceSSRResult{Error: fmt.Errorf("render error: %w", err)}
    }
    
    var result map[string]interface{}
    if err := json.Unmarshal(rendered, &result); err != nil {
        return EnhanceSSRResult{Error: fmt.Errorf("unmarshal error: %w", err)}
    }
    
    body, ok := result["body"].(string)
    if !ok {
        return EnhanceSSRResult{Error: fmt.Errorf("rendered document body is not a string")}
    }
    
    return EnhanceSSRResult{Body: body}
}