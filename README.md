# Data Race Skeletons from Uber's Go Codebase

[![CI](https://github.com/uber-research/drfix/actions/workflows/ci.yml/badge.svg)](https://github.com/uber-research/drfix/actions/workflows/ci.yml)

## About

DR.FIX is a research project that provides a collection of data race skeletons extracted from Uber's Go codebase. These skeletons are created through program slicing to retain key elements such as racy variables and concurrency constructs, making them valuable for research in concurrency bug detection and analysis.

This repository is part of our research on automated data race detection and fixing. For more details, please refer to our paper: [DR.FIX: Automatically Fixing Data Races at Industry Scale](https://doi.org/10.1145/3729265).

## Research Use Cases

These data race skeletons are valuable for:
1. Training data race detectors
2. Analyzing common patterns in Go data races
3. Studying the relationship between write operations and race conditions
4. Developing new techniques for race detection and prevention

## Project Structure

- `cmd/analyzer/`: Contains the Go analyzer that detects write operations on racy variables
- `internal/analyzer/`: Core analysis logic for detecting write operations
- `scripts/`: Python scripts for processing and verifying the skeletons
- `data/skeletons/`: Directory containing the data race skeletons
- `data/examples/`: Directory containing real code examples showing data races and their fixes
- `output/`: Directory for analysis results and combined files

## How to Use

### Prerequisites

- Go 1.16 or later
- Python 3.8 or later
- Make

### Setting Up the Environment

1. Clone the repository:
   ```bash
   git clone https://github.com/uber-research/drfix.git
   cd drfix
   ```

2. Create and activate the Python virtual environment:
   ```bash
   python3 -m venv venv
   source venv/bin/activate  # On Windows, use: venv\Scripts\activate
   pip install -r requirements.txt
   ```

   Note: You need to activate the virtual environment in each new terminal session where you want to run Python commands.

### Building the Analyzer

Build the Go analyzer:
```bash
make build
```
This will compile the analyzer binary and place it in the `bin/` directory.

### Analyzing Skeletons

Run the analysis:
```bash
make analyze
```
This will:
- Process all skeleton pairs in `data/skeletons/`
- Generate analysis results in `output/analyzer_results.csv`
- Create a combined file in `output/combined.go`

### Running Tests

Run all tests:
```bash
make test
```
This will run both Go and Python tests.

### Debug Mode

To enable debug output from both the Go analyzer and Python script:
```bash
DEBUG=1 make analyze
```

## Verification Tools

### Go Analyzer

The Go analyzer (`cmd/analyzer/main.go`) performs AST-based analysis to detect write operations on racy variables. It outputs results in JSON format:

```json
{
    "file": "path/to/file.go",
    "has_write": true,
    "line_number": 42,
    "line_content": "racyVar0 = 42",
    "error": null
}
```

### Python Processing Script

The Python script (`scripts/process.py`) verifies the skeletons and generates a comprehensive CSV report. It checks:
1. Write operations on racy variables
2. Syntactical validity of Go code
3. Presence of required concurrency constructs
4. File characteristics and statistics

## Output Format

### CSV Output

The verification script generates a CSV file with the following columns:

#### Basic Information
- `case_id`: Directory name (e.g., D7959843)
- `file1_name`: Name of the first file
- `file2_name`: Name of the second file

#### Analysis Results
- `file1_status`: Boolean indicating if first file contains a write operation (True/False)
- `file1_line`: Line containing write operation in first file (if True)
- `file2_status`: Boolean indicating if second file contains a write operation (True/False)
- `file2_line`: Line containing write operation in second file (if True)
- `has_write`: TRUE if either file contains a write to a racy variable
- `race_type`: Type of race condition
  - `read-write`: One file reads while the other writes to the same variable
  - `write-write`: Both files write to the same variable

#### File Statistics
- `file1_size`: Size of first file in bytes
- `file2_size`: Size of second file in bytes
- `file1_line_count`: Number of lines in first file
- `file2_line_count`: Number of lines in second file
- `file1_has_package`: Whether first file has package declaration (True/False)
- `file2_has_package`: Whether second file has package declaration (True/False)
- `file1_has_comments`: Whether first file has comments (True/False)
- `file2_has_comments`: Whether second file has comments (True/False)

#### Error Information
- `error_message`: Any error messages encountered during analysis

Example CSV output:
```csv
case_id,file1_name,file1_status,file1_line,file2_name,file2_status,file2_line,has_write,race_type,file1_size,file2_size,file1_line_count,file2_line_count,file1_has_package,file2_has_package,file1_has_comments,file2_has_comments,error_message
D13766163,read2.go,True,"v22, racyVar0 = v2.v37.Func45(v13)",write1.go,True,"v22, racyVar0 = v2.v37.Func45(v13)",TRUE,write-write,1239,1239,60,60,True,True,False,False,
D15314495,read1.go,True,racyVar0 += int64(v10),write2.go,True,racyVar0 += int64(v10),TRUE,read-write,980,980,25,25,True,True,False,False,
D12646115,read2.go,True,"for _, racyVar0 := range v1 {",write1.go,True,"for _, racyVar0 := range v1 {",TRUE,read-write,273,273,13,13,True,True,False,False,
```


## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Citation

If you use this work in your research, please cite:

```
@article{behrang2025drfix,
  author = {Behrang, Farnaz and Zhang, Zhizhou and Saioc, Georgian-Vlad and Liu, Peng and Chabbi, Milind},
  title = {{DR.FIX: Automatically Fixing Data Races at Industry Scale}},
  year = {2025},
  issue_date = {June 2025},
  publisher = {Association for Computing Machinery},
  address = {New York, NY, USA},
  volume = {9},
  number = {Programming Language Design and Implementation},
  articleno = {166},
  numpages = {37},
  url = {https://doi.org/10.1145/3729265},
  doi = {10.1145/3729265},
  journal = {Proceedings of the ACM on Programming Languages},
  month = {jun},
  keywords = {Data Races, Program Analysis, Static Analysis, Dynamic Analysis, Go}
}
```

## Contributing

We welcome contributions in the following areas:
1. New data race skeletons
2. Improvements to the verification tools
3. Analysis of race patterns and write operations
4. Documentation and examples

Please read our [Code of Conduct](CODE_OF_CONDUCT.md) and [Contributing Guidelines](CONTRIBUTING.md) before contributing.
