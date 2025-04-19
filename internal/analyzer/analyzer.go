package analyzer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

var debugMode bool

// AnalysisResult represents the result of analyzing a Go file
type AnalysisResult struct {
	File        string `json:"file"`
	HasWrite    bool   `json:"has_write"` // Whether the file contains a write to a racy variable
	LineNumber  int    `json:"line_number,omitempty"`
	LineContent string `json:"line_content,omitempty"`
	Error       string `json:"error,omitempty"`
}

// SetDebugMode sets the debug mode for the analyzer
func SetDebugMode(debug bool) {
	debugMode = debug
}

func debug(format string, args ...interface{}) {
	if debugMode {
		fmt.Printf(format, args...)
	}
}

// isRacyVarWrite checks if an identifier is a racy variable write
func isRacyVarWrite(ident *ast.Ident) bool {
	debug("Checking if %s is a racy var\n", ident.Name)
	return strings.HasPrefix(ident.Name, "racyVar")
}

// isReadOnlyContext checks if a node puts its child in a read-only context
func isReadOnlyContext(node ast.Node, child ast.Expr) bool {
	debug("Checking read-only context for %T\n", node)
	switch n := node.(type) {
	case *ast.SelectorExpr:
		return n.X == child
	case *ast.IndexExpr:
		return n.Index == child
	case *ast.SliceExpr:
		return n.Low == child || n.High == child || n.Max == child
	case *ast.CallExpr:
		for _, arg := range n.Args {
			if arg == child {
				debug("Found read in CallExpr\n")
				return true
			}
		}
	case *ast.KeyValueExpr:
		return n.Key == child || n.Value == child
	case *ast.TypeAssertExpr:
		return true
	case *ast.UnaryExpr:
		return n.X == child
	case *ast.BinaryExpr:
		return n.X == child || n.Y == child
	// case *ast.ParenExpr:
	// 	return n.X == child
	case *ast.ReturnStmt:
		for _, result := range n.Results {
			if result == child {
				debug("Found read in ReturnStmt\n")
				return true
			}
		}
	case *ast.SendStmt:
		return n.Chan == child || n.Value == child
	case *ast.CommClause:
		return true
	case *ast.TypeSpec, *ast.StructType, *ast.InterfaceType, *ast.FuncType, *ast.MapType, *ast.ChanType, *ast.ArrayType:
		return true
	case *ast.BlockStmt, *ast.SwitchStmt, *ast.IfStmt, *ast.ForStmt, *ast.DeferStmt, *ast.GoStmt, *ast.SelectStmt, *ast.CaseClause:
		return true
	}
	return false
}

// isWriteContext checks if a node puts its child in a write context
func isWriteContext(node ast.Node, child ast.Expr) bool {
	debug("Checking write context for %T\n", node)
	switch n := node.(type) {
	case *ast.AssignStmt:
		debug("AssignStmt with token %v\n", n.Tok)
		// All except :=
		if n.Tok != token.DEFINE {
			for _, lhs := range n.Lhs {
				if lhs == child {
					debug("Found write in AssignStmt\n")
					return true
				}
			}
		}
	case *ast.IncDecStmt:
		debug("IncDecStmt\n")
		return n.X == child
	case *ast.RangeStmt:
		debug("RangeStmt with token %v\n", n.Tok)
		if n.Tok == token.DEFINE || n.Tok == token.ASSIGN {
			if n.Key == child || n.Value == child {
				debug("Found write in RangeStmt\n")
				return true
			}
		}
	case *ast.FuncDecl:
		debug("FuncDecl\n")
		if n.Type.Results != nil {
			for _, field := range n.Type.Results.List {
				for _, name := range field.Names {
					if name == child {
						debug("Found write in FuncDecl\n")
						return true
					}
				}
			}
		}
	}
	return false
}

// isWritten checks if an identifier is being written to by walking up the AST
func isWritten(ident *ast.Ident, stack []ast.Node) bool {
	debug("Checking if %s is written\n", ident.Name)
	// Walk up the stack to find if any parent node writes to this expression
	var child ast.Expr = ident
	for i := len(stack) - 1; i >= 0; i-- {
		parent := stack[i]
		debug("Parent type: %T\n", parent)

		// If we hit a read-only context, we can stop
		if isReadOnlyContext(parent, child) {
			debug("Found read-only context\n")
			return false
		}

		// If we hit a write context, we found our write
		if isWriteContext(parent, child) {
			debug("Found write context\n")
			return true
		}

		// Only continue if the parent is an expression
		if expr, ok := parent.(ast.Expr); ok {
			child = expr
		} else {
			break
		}
	}
	return false
}

// visitor implements the ast.Visitor interface
type visitor struct {
	hasRacyWrite bool
	fset         *token.FileSet
	stack        []ast.Node     // Keep track of parent nodes
	writePos     token.Position // Track position of racy write
}

// RacyWrite represents a found racy variable write
type RacyWrite struct {
	LineNumber int
	LineText   string
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		if len(v.stack) > 0 {
			v.stack = v.stack[:len(v.stack)-1]
		}
		return v
	}

	v.stack = append(v.stack, node)

	switch n := node.(type) {
	case *ast.Ident:
		debug("Visiting identifier: %s\n", n.Name)
		if isRacyVarWrite(n) && isWritten(n, v.stack) {
			debug("Found racy write for %s\n", n.Name)
			v.hasRacyWrite = true
			v.writePos = v.fset.Position(n.Pos())
		}
	}

	return v
}

// AnalyzeFile analyzes a Go source file and returns whether it contains writes to racy variables
func AnalyzeFile(filename string) (*AnalysisResult, error) {
	result := &AnalysisResult{
		File: filename,
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		result.Error = fmt.Sprintf("error parsing file: %v", err)
		return result, err
	}

	v := &visitor{
		fset:  fset,
		stack: make([]ast.Node, 0),
	}
	ast.Walk(v, node)

	if v.hasRacyWrite {
		result.HasWrite = true
		result.LineNumber = v.writePos.Line

		// Read the file to get the line content
		content, err := os.ReadFile(filename)
		if err != nil {
			result.Error = fmt.Sprintf("error reading file: %v", err)
			return result, err
		}
		lines := strings.Split(string(content), "\n")
		if v.writePos.Line-1 < len(lines) {
			result.LineContent = strings.TrimSpace(lines[v.writePos.Line-1])
		}
	}

	return result, nil
}
