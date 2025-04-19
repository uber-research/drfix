package analyzer

import (
	"testing"
)

func TestAnalyzeFile(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "write to racyVar",
			filename: "testdata/write.go",
			want:     true,
		},
		{
			name:     "read from racyVar",
			filename: "testdata/read.go",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AnalyzeFile(tt.filename)
			if err != nil {
				t.Fatalf("AnalyzeFile() error = %v", err)
			}
			if result.HasWrite != tt.want {
				t.Errorf("AnalyzeFile() = %v, want %v", result.HasWrite, tt.want)
			}
		})
	}
}
