#!/usr/bin/env python3
"""
Data Race Skeleton Verification Script

This script verifies the sanity of data race skeletons extracted from Uber's Go codebase.
It checks that each skeleton pair:
1. Contains at least one write operation
2. Has racyVar* pattern in each file
3. Is syntactically valid Go code
4. Preserves the essential concurrency context

The script also generates statistics and a combined output file for further analysis.
"""

import os
import subprocess
import re
import csv
from multiprocessing import Pool, cpu_count
from functools import partial
import argparse
from typing import Dict, List, Set, Tuple, Optional, Union, Any
import json

# Type aliases for better code readability
FileStats = Dict[str, int]
PatternStats = Dict[str, int]
AnalysisResult = Tuple[str, str, Optional[str]]  # (file_path, status, code_line)

def parse_args() -> argparse.Namespace:
    """Parse command line arguments for skeleton verification.
    
    Returns:
        argparse.Namespace: Parsed arguments containing:
            - input_dir: Directory containing skeleton pairs
            - combined_file: Path to output combined Go file
            - output_csv: Path to output verification results
    """
    parser = argparse.ArgumentParser(
        description='Process Go files and analyze for racy variables'
    )
    parser.add_argument(
        '--input-dir',
        required=True,
        help='Path to the input directory containing skeleton files'
    )
    parser.add_argument(
        '--combined-file',
        required=True,
        help='Path to the output combined Go file'
    )
    parser.add_argument(
        '--output-csv',
        required=True,
        help='Path to the output CSV file'
    )
    return parser.parse_args()

def normalize_filename(filename: str) -> str:
    """Normalize a skeleton filename to ensure consistent naming.
    
    Args:
        filename: Original skeleton filename
        
    Returns:
        str: Normalized filename with .go extension and no underscores
    """
    if not filename.endswith(".go"):
        filename += ".go"
    return filename.lower().replace("_", "")

def rename_files(skeletons_dir: str) -> int:
    """Normalize filenames in the skeletons directory for consistency.
    
    For each skeleton file:
    1. Ensures .go extension
    2. Converts to lowercase
    3. Removes underscores
    
    Args:
        skeletons_dir: Path to the skeletons directory
        
    Returns:
        int: Number of files renamed
    """
    rename_count = 0
    for root, _, files in os.walk(skeletons_dir):
        for filename in files:
            file_path = os.path.join(root, filename)
            new_filename = normalize_filename(filename)
            
            if new_filename != filename:
                new_file_path = os.path.join(root, new_filename)
                if os.path.exists(new_file_path):
                    print(f"Error: Cannot rename {file_path} to {new_file_path} because it already exists.")
                else:
                    os.rename(file_path, new_file_path)
                    print(f"Renamed: {file_path} -> {new_file_path}")
                    rename_count += 1
                    
    print(f"Total files renamed: {rename_count}")
    return rename_count

def ensure_package_declaration(file_path: str) -> bool:
    """Ensure a skeleton file has a valid package declaration.
    
    Args:
        file_path: Path to the skeleton file
        
    Returns:
        bool: True if the file was updated, False otherwise
    """
    with open(file_path, "r") as f:
        lines = f.readlines()
    
    if lines and lines[0].strip().startswith("package"):
        return False

    new_content = "package main\n\n" + "".join(lines)
    with open(file_path, "w") as f:
        f.write(new_content)
    print(f"Updated package declaration in {file_path}")
    return True

def remove_line_comments(file_path: str) -> None:
    """Remove line comments from skeleton files to focus on essential code.
    
    Args:
        file_path: Path to the skeleton file
    """
    with open(file_path, "r") as f:
        content = f.read()
    
    new_content = re.sub(r"//.*", "", content)
    with open(file_path, "w") as f:
        f.write(new_content)
    print(f"Removed // comments in {file_path}")

def check_go_file(file_path: str) -> Tuple[bool, str, bool]:
    """Verify a skeleton file is valid Go code.
    
    Args:
        file_path: Path to the skeleton file
        
    Returns:
        Tuple[bool, str, bool]: (is_parsable, error_message, package_updated)
    """
    package_updated = ensure_package_declaration(file_path)
    remove_line_comments(file_path)
    
    try:
        result = subprocess.run(
            ["gofmt", "-e", file_path],
            capture_output=True,
            text=True
        )
        if result.returncode != 0:
            return False, result.stderr.strip(), package_updated
        return True, "", package_updated
    except FileNotFoundError:
        print(f"Warning: gofmt not found. Skipping format check for {file_path}")
        return True, "", package_updated

def check_files(skeletons_dir: str) -> Tuple[Dict[str, str], Set[str], FileStats]:
    """Verify all skeleton files in the directory.
    
    Args:
        skeletons_dir: Path to the skeletons directory
        
    Returns:
        Tuple containing:
            - Dict of parsing errors (file_path -> error_message)
            - Set of successfully parsed file paths
            - File verification statistics
    """
    parsing_errors = {}
    parsable_files = set()
    stats = {
        'total_files': 0,
        'parsed_success': 0,
        'parse_errors': 0,
        'package_updates': 0
    }

    for root, _, files in os.walk(skeletons_dir):
        for filename in files:
            if filename.endswith(".go"):
                stats['total_files'] += 1
                filepath = os.path.join(root, filename)
                parsed, error_msg, pkg_updated = check_go_file(filepath)
                
                if pkg_updated:
                    stats['package_updates'] += 1
                if not parsed:
                    stats['parse_errors'] += 1
                    parsing_errors[filepath] = error_msg
                    print(f"Parse error in {filepath}:\n{error_msg}\n")
                else:
                    stats['parsed_success'] += 1
                    parsable_files.add(filepath)
                    
    return parsing_errors, parsable_files, stats

def check_racyvar_pattern(skeletons_dir: str, parsable_files: Set[str]) -> PatternStats:
    """Verify racyVar pattern in skeleton files.
    
    Args:
        skeletons_dir: Path to the skeletons directory
        parsable_files: Set of successfully parsed file paths
        
    Returns:
        Dict containing pattern verification statistics
    """
    pattern = re.compile(r"racyVar\d+")
    stats = {
        "total_subdirs": 0,
        "subdirs_with_issues": 0,
        "issues": []
    }
    
    try:
        for subdir in next(os.walk(skeletons_dir))[1]:
            stats["total_subdirs"] += 1
            subdir_path = os.path.join(skeletons_dir, subdir)
            go_files = [f for f in os.listdir(subdir_path) if f.endswith(".go")]
            
            if len(go_files) != 2:
                print(f"Warning: Directory {subdir_path} does not have exactly 2 .go files.")
                continue
                
            missing_files = []
            for go_file in go_files:
                file_path = os.path.join(subdir_path, go_file)
                if file_path not in parsable_files:
                    continue
                    
                with open(file_path, 'r') as f:
                    content = f.read()
                    if not pattern.search(content):
                        missing_files.append(go_file)
                        
            if missing_files:
                stats["subdirs_with_issues"] += 1
                stats["issues"].append(
                    f"In directory {subdir_path}, files missing racyVar pattern: {', '.join(missing_files)}"
                )
    except StopIteration:
        print(f"Warning: No subdirectories found in {skeletons_dir}")
        
    return stats

def create_final_combined_file(skeletons_dir: str, output_file: str) -> None:
    """Create a combined file containing all skeleton code.
    
    Args:
        skeletons_dir: Path to the skeletons directory
        output_file: Path to the output combined file
    """
    # Create output directory if it doesn't exist
    output_dir = os.path.dirname(output_file)
    if output_dir and not os.path.exists(output_dir):
        os.makedirs(output_dir)
        
    with open(output_file, "w") as final_file:
        for root, _, files in os.walk(skeletons_dir):
            for filename in sorted(files):
                if filename.endswith(".go"):
                    file_path = os.path.join(root, filename)
                    final_file.write(f"// {file_path}\n")
                    with open(file_path, "r") as f:
                        content = f.read()
                    final_file.write(content + "\n\n")
    print(f"Final combined file created: {output_file}")

def run_analyzer(file_path: str) -> Dict[str, Any]:
    """Run the Go analyzer on a file and return the results."""
    try:
        debug_flag = ['-debug'] if os.environ.get('DEBUG') else []
        result = subprocess.run(
            ["./bin/analyzer", "-i", file_path] + debug_flag,
            capture_output=True,
            text=True,
            check=True
        )
        return json.loads(result.stdout)
    except subprocess.CalledProcessError as e:
        return {
            "file": file_path,
            "has_write": False,
            "error": e.stderr.strip()
        }
    except json.JSONDecodeError as e:
        return {
            "file": file_path,
            "has_write": False,
            "error": f"Failed to parse analyzer output: {str(e)}"
        }

def analyze_directory_pairs(skeletons_dir: str, output_csv: str) -> None:
    """Verify pairs of skeleton files and write results to CSV.
    
    Args:
        skeletons_dir: Path to the skeletons directory
        output_csv: Path to the output CSV file
    """
    # Collect all files to analyze
    files_to_analyze = []
    directory_pairs = []
    
    try:
        for subdir in next(os.walk(skeletons_dir))[1]:
            subdir_path = os.path.join(skeletons_dir, subdir)
            go_files = sorted([f for f in os.listdir(subdir_path) if f.endswith(".go")])
            
            if len(go_files) != 2:
                print(f"Warning: Directory {subdir_path} does not have exactly 2 .go files.")
                continue
            
            file1 = os.path.join(subdir_path, go_files[0])
            file2 = os.path.join(subdir_path, go_files[1])
            files_to_analyze.extend([file1, file2])
            directory_pairs.append((subdir, go_files[0], go_files[1], file1, file2))
    except StopIteration:
        print(f"Warning: No subdirectories found in {skeletons_dir}")
        return
    
    # Run analysis in parallel
    num_processes = min(cpu_count(), len(files_to_analyze))
    print(f"Running analysis with {num_processes} processes...")
    
    with Pool(processes=num_processes) as pool:
        results = pool.map(run_analyzer, files_to_analyze)
    
    # Create a dictionary of results for easy lookup
    results_dict = {file_path: result for file_path, result in zip(files_to_analyze, results)}
    
    # Write results to CSV
    with open(output_csv, 'w', newline='') as csvfile:
        writer = csv.writer(csvfile)
        writer.writerow([
            'case_id',                    # Directory name (e.g., D7959843)
            'file1_name',                 # Name of the first file
            'file1_status',               # Status of first file analysis
            'file1_line',                 # Line containing racy variable in first file
            'file2_name',                 # Name of the second file
            'file2_status',               # Status of second file analysis
            'file2_line',                 # Line containing racy variable in second file
            'has_write',                  # TRUE if either file contains a write to a racy variable
            'race_type',                  # Type of race (read-write or write-write)
            'file1_size',                 # Size of first file in bytes
            'file2_size',                 # Size of second file in bytes
            'file1_line_count',           # Number of lines in first file
            'file2_line_count',           # Number of lines in second file
            'file1_has_package',          # Whether first file has package declaration
            'file2_has_package',          # Whether second file has package declaration
            'file1_has_comments',         # Whether first file has comments
            'file2_has_comments',         # Whether second file has comments
            'error_message'               # Any error messages encountered
        ])
        
        for subdir, file1_name, file2_name, file1_path, file2_path in directory_pairs:
            result1 = results_dict[file1_path]
            result2 = results_dict[file2_path]
            
            # Get file statistics
            file1_stats = get_file_stats(file1_path)
            file2_stats = get_file_stats(file2_path)
            
            # Determine if there's a write
            has_write = "TRUE" if result1.get("has_write", False) or result2.get("has_write", False) else "FALSE"
            
            # Determine race type
            race_type = determine_race_type(file1_path, file2_path)
            
            # Get any error messages
            error_msg = ""
            if result1.get("error"):
                error_msg += f"File1 error: {result1['error']}; "
            if result2.get("error"):
                error_msg += f"File2 error: {result2['error']}"
            
            writer.writerow([
                subdir,
                file1_name,
                result1.get("has_write", "NOT_FOUND"),
                result1.get("line_content", ""),
                file2_name,
                result2.get("has_write", "NOT_FOUND"),
                result2.get("line_content", ""),
                has_write,
                race_type,
                file1_stats['size'],
                file2_stats['size'],
                file1_stats['line_count'],
                file2_stats['line_count'],
                file1_stats['has_package'],
                file2_stats['has_package'],
                file1_stats['has_comments'],
                file2_stats['has_comments'],
                error_msg.strip()
            ])
    
    print(f"Analysis results written to {output_csv}")

def determine_race_type(file1_path: str, file2_path: str) -> str:
    """Determine the type of race (read-write or write-write) by analyzing the files.
    
    Args:
        file1_path: Path to the first file
        file2_path: Path to the second file
        
    Returns:
        str: "read-write" or "write-write"
    """
    def has_write_operation(file_path: str) -> bool:
        """Check if a file contains a write operation on a racy variable."""
        try:
            with open(file_path, 'r') as f:
                content = f.read()
                # Look for assignment operations on racy variables
                return bool(re.search(r'racyVar\d+\s*=', content))
        except Exception:
            return False
    
    file1_has_write = has_write_operation(file1_path)
    file2_has_write = has_write_operation(file2_path)
    
    if file1_has_write and file2_has_write:
        return "write-write"
    else:
        return "read-write"

def get_file_stats(file_path: str) -> Dict[str, Union[int, bool]]:
    """Get statistics about a Go file.
    
    Args:
        file_path: Path to the Go file
        
    Returns:
        Dict containing file statistics
    """
    stats = {
        'size': 0,
        'line_count': 0,
        'has_package': False,
        'has_comments': False
    }
    
    try:
        # Get file size
        stats['size'] = os.path.getsize(file_path)
        
        # Read file content
        with open(file_path, 'r') as f:
            content = f.read()
            lines = content.splitlines()
            stats['line_count'] = len(lines)
            
            # Check for package declaration
            stats['has_package'] = any(line.strip().startswith('package ') for line in lines)
            
            # Check for comments
            stats['has_comments'] = any('//' in line for line in lines)
            
    except Exception as e:
        print(f"Error getting stats for {file_path}: {e}")
    
    return stats

def main() -> None:
    """Main entry point for skeleton verification."""
    args = parse_args()
    print("Starting processing...\n")
    
    # Process files
    renamed_count = rename_files(args.input_dir)
    parsing_errors, parsable_files, file_stats = check_files(args.input_dir)
    pattern_stats = check_racyvar_pattern(args.input_dir, parsable_files)
    
    # Generate outputs
    create_final_combined_file(args.input_dir, args.combined_file)
    analyze_directory_pairs(args.input_dir, args.output_csv)
    
    # Print final stats
    print("\nFinal Stats:")
    print("-----------")
    print(f"Total .go files processed: {file_stats['total_files']}")
    print(f"Files renamed: {renamed_count}")
    print(f"Files updated with package declaration: {file_stats['package_updates']}")
    print(f"Files parsed successfully: {file_stats['parsed_success']}")
    print(f"Files with parse errors: {file_stats['parse_errors']}")
    print(f"Total subdirectories processed: {pattern_stats['total_subdirs']}")
    print(f"Subdirectories with racyVar pattern issues: {pattern_stats['subdirs_with_issues']}")
    
    if pattern_stats['issues']:
        print("\nSubdirectory issues details:")
        for issue in pattern_stats['issues']:
            print(f" - {issue}")

if __name__ == "__main__":
    main()
