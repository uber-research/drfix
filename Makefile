.PHONY: all build test analyze clean

all: build test

build:
	go build -o bin/analyzer cmd/analyzer/main.go

test:
	go test ./cmd/... ./internal/...
	#python3 -m pytest scripts/test_*.py

analyze: build
	DEBUG=$(DEBUG) python3 scripts/process.py --input-dir data/skeletons/ --combined-file output/final_combined.go --output-csv output/analyzer_results.csv

clean:
	rm -rf bin/
	rm -rf output/
