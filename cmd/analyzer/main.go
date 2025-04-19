package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/uber/data-race-skeletons/internal/analyzer"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	inputFile := flag.String("i", "", "Input Go file to analyze")
	flag.Parse()

	if *inputFile == "" {
		fmt.Fprintf(os.Stderr, "Error: Input file is required\n")
		flag.Usage()
		os.Exit(1)
	}

	analyzer.SetDebugMode(*debug)
	result, err := analyzer.AnalyzeFile(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error analyzing file: %v\n", err)
		os.Exit(1)
	}

	// Output result as JSON
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling result: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
