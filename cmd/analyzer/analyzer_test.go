package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestAnalyzer(t *testing.T) {
	// Build the analyzer binary
	cmd := exec.Command("go", "build", "-o", "analyzer")
	cmd.Dir = "."
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to build analyzer: %v", err)
	}
	defer os.Remove("analyzer")

	// Test cases
	tests := []struct {
		name      string
		args      []string
		wantErr   bool
		debugFlag bool
	}{
		{
			name:      "missing input file",
			args:      []string{"./analyzer"},
			wantErr:   true,
			debugFlag: false,
		},
		{
			name:      "with debug flag",
			args:      []string{"./analyzer", "-debug", "-i", "testdata/test.go"},
			wantErr:   false,
			debugFlag: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(tt.args[0], tt.args[1:]...)
			output, err := cmd.CombinedOutput()

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v\nOutput: %s", err, output)
				}
			}
		})
	}
}
