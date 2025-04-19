"""
Tests for the data race skeleton analyzer processing script.

This module contains tests for the main functionality of process.py,
including argument parsing and file processing.
"""

import os
import tempfile
import pytest
import sys
from process import (
    parse_args,
    create_final_combined_file,
    normalize_filename,
    ensure_package_declaration,
    remove_line_comments
)

def test_parse_args(monkeypatch):
    """Test command line argument parsing.
    
    Verifies that the script correctly parses input directory,
    combined file, and output CSV paths.
    """
    # Mock sys.argv with test arguments
    test_args = [
        'process.py',
        '--input-dir', 'test_input',
        '--combined-file', 'test.go',
        '--output-csv', 'test.csv'
    ]
    monkeypatch.setattr(sys, 'argv', test_args)

    # Parse arguments and verify values
    args = parse_args()
    assert args.input_dir == 'test_input'
    assert args.combined_file == 'test.go'
    assert args.output_csv == 'test.csv'

def test_normalize_filename():
    """Test filename normalization.
    
    Verifies that filenames are correctly normalized by:
    1. Adding .go extension if missing
    2. Converting to lowercase
    3. Removing underscores
    """
    # Test cases: (input, expected_output)
    test_cases = [
        ('test.go', 'test.go'),
        ('TEST.go', 'test.go'),
        ('test_file.go', 'testfile.go'),
        ('test', 'test.go'),
        ('TEST_FILE', 'testfile.go')
    ]
    
    for input_name, expected in test_cases:
        assert normalize_filename(input_name) == expected

def test_ensure_package_declaration():
    """Test package declaration handling.
    
    Verifies that:
    1. Files with existing package declarations are unchanged
    2. Files without package declarations get 'package main' added
    """
    with tempfile.TemporaryDirectory() as tmpdir:
        # Test case 1: File with existing package declaration
        file1 = os.path.join(tmpdir, 'test1.go')
        with open(file1, 'w') as f:
            f.write('package test\n\nfunc main() {}\n')
        
        assert not ensure_package_declaration(file1)
        with open(file1, 'r') as f:
            assert f.read().startswith('package test')
        
        # Test case 2: File without package declaration
        file2 = os.path.join(tmpdir, 'test2.go')
        with open(file2, 'w') as f:
            f.write('func main() {}\n')
        
        assert ensure_package_declaration(file2)
        with open(file2, 'r') as f:
            content = f.read()
            assert content.startswith('package main\n\n')
            assert 'func main() {}' in content

def test_remove_line_comments():
    """Test line comment removal.
    
    Verifies that // comments are correctly removed from Go files.
    """
    with tempfile.TemporaryDirectory() as tmpdir:
        test_file = os.path.join(tmpdir, 'test.go')
        with open(test_file, 'w') as f:
            f.write('''package main

// This is a comment
func main() {
    // Another comment
    x := 1 // Inline comment
}
''')
        
        remove_line_comments(test_file)
        
        with open(test_file, 'r') as f:
            content = f.read()
            assert '//' not in content
            assert 'func main()' in content
            assert 'x := 1' in content

def test_create_final_combined_file():
    """Test combined file creation.
    
    Verifies that:
    1. Multiple Go files are correctly combined
    2. File paths are preserved as comments
    3. Files are processed in sorted order
    """
    with tempfile.TemporaryDirectory() as tmpdir:
        # Create input directory with test files
        input_dir = os.path.join(tmpdir, 'input')
        os.makedirs(input_dir)
        
        # Create test files
        files = [
            ('a.go', 'package main\n\nfunc a() {}\n'),
            ('b.go', 'package main\n\nfunc b() {}\n'),
            ('c.go', 'package main\n\nfunc c() {}\n')
        ]
        
        for filename, content in files:
            with open(os.path.join(input_dir, filename), 'w') as f:
                f.write(content)
        
        # Create combined file
        output_file = os.path.join(tmpdir, 'combined.go')
        create_final_combined_file(input_dir, output_file)
        
        # Verify output
        assert os.path.exists(output_file)
        with open(output_file, 'r') as f:
            content = f.read()
            
            # Check that all files are included
            for filename, _ in files:
                assert f'// {os.path.join(input_dir, filename)}' in content
            
            # Check that content is in sorted order
            assert content.find('func a()') < content.find('func b()') < content.find('func c()') 