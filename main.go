package main

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"

	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"

	"bytes"
	"encoding/json"
	"path/filepath"

	extism "github.com/extism/go-sdk"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zangster300/northstar/routes"
	"golang.org/x/sync/errgroup"
)

const port = 8080

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info(fmt.Sprintf("Starting Server @:%d", port))
	defer logger.Info("Stopping Server")

	// ReadReportsAndFindSafeReportsBasedOnLevels()

	markup := "<my-header>Hello World</my-header>"
	elementPath := "./web/custom-elements"
	elements := readElements(elementPath)
	initialState := make(map[string]interface{})
	data := map[string]interface{}{
		"markup":       markup,
		"elements":     elements,
		"initialState": initialState,
	}
	payload, _ := Marshal(data)

	rendered, _ := render(payload)

	var result map[string]interface{}
	if err := json.Unmarshal(rendered, &result); err != nil {
		fmt.Println("Failed to parse rendered output", err)
		return
	}

	fmt.Println("result", result)

	element_body, ok := result["body"].(string)
	if !ok {
		fmt.Println("Rendered document is not a string")
		return
	}

	fmt.Println("element_body", element_body)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := run(ctx, logger); err != nil {
		logger.Error("Error running server", slog.Any("err", err))
		os.Exit(1)
	}
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


// Day 2 of Advent of Code challenge

func ReadReportsAndFindSafeReportsBasedOnLevels() {
	file, error := os.Open("AoC/day2-input.txt")

	if error != nil {
		fmt.Println("Error opening file:", error)
		return
	}
	defer file.Close()
	validListCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers, err := convertLineToNumbers(line)
		if err != nil {
			fmt.Println("Error converting line to numbers:", err)
			continue
		}

		if isValidList(numbers) {
			fmt.Println("Valid list:", numbers)
			validListCount++
		} else {
			fmt.Println("Invalid list:", numbers)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Valid list count: %d\n", validListCount)
}

func convertLineToNumbers(line string) ([]int, error) {
	values := strings.Fields(line)

	numbers := make([]int, len(values))
	for i, value := range values {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("error converting value to integer: %v", err)
		}
		numbers[i] = num
	}

	return numbers, nil
}

func isValidList(numbers []int) bool {
	// First check if the list is valid as-is
	if isStrictlyValid(numbers) {
		return true
	}

	// Try removing each number one at a time
	for i := 0; i < len(numbers); i++ {
		// Create a new slice without the current number
		tempList := make([]int, 0, len(numbers)-1)
		tempList = append(tempList, numbers[:i]...)
		tempList = append(tempList, numbers[i+1:]...)
		
		// Check if removing this number makes the list valid
		if isStrictlyValid(tempList) {
			return true
		}
	}

	return false
}

func isStrictlyValid(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}

	increasing := true
	decreasing := true

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		absDiff := abs(diff)
		
		// Check for valid differences (1-3)
		if absDiff < 1 || absDiff > 3 {
			return false
		}
		
		// Check for strictly increasing/decreasing
		if diff <= 0 {  // Changed from < to <= to handle duplicates
			increasing = false
		}
		if diff >= 0 {  // Changed from > to >= to handle duplicates
			decreasing = false
		}
	}

	return increasing || decreasing
}

// Helper function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


// Day 1 of Advent of Code challenge
func ReadListAndSplit() {
	file, err := os.Open("AoC/day1-list.txt")
if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize slices for left and right columns
	var leftColumn []int
	var rightColumn []int

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into two parts
		values := strings.Fields(line) // Splits by spaces or tabs

		// Ensure there are exactly 2 columns per line
		if len(values) != 2 {
			fmt.Printf("Invalid line format: %s\n", line)
			return
		}

		// Convert the two parts into integers
		leftValue, err1 := strconv.Atoi(values[0])
		rightValue, err2 := strconv.Atoi(values[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error converting values to integers: %v, %v\n", err1, err2)
			return
		}

		// Append to respective columns
		leftColumn = append(leftColumn, leftValue)
		rightColumn = append(rightColumn, rightValue)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort both columns
	sort.Ints(leftColumn)
	sort.Ints(rightColumn)

	// Step 1: Calculate the similarity scores
	similarityScores := calculateSimilarityScores(leftColumn, rightColumn)

	// Step 2: Calculate the total similarity score
	totalSimilarity := 0
	for _, score := range similarityScores {
		totalSimilarity += score
	}

	// Print the similarity scores and total
	fmt.Println("Similarity Scores:", similarityScores)
	fmt.Println("Total Similarity Score:", totalSimilarity)
}

// Helper function to calculate similarity scores
func calculateSimilarityScores(left, right []int) []int {
	// Map to count occurrences of each number in the right list
	rightCount := make(map[int]int)
	for _, num := range right {
		rightCount[num]++
	}

	// Calculate similarity scores for the left list
	var similarityScores []int
	for _, num := range left {
		count := rightCount[num] // Get the count of this number in the right list
		similarityScore := num * count
		similarityScores = append(similarityScores, similarityScore)
	}

	return similarityScores
}