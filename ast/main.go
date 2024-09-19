package main

import (
	"encoding/hex"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// doc
func funcion(n int) int64 {
	one := 1
	n_plus_one := n + one
	gauss_sum := n * n_plus_one / 2
	return int64(gauss_sum)
}

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

func randIdPrefix(prefix string) string {
	return prefix + "_" + randId()
}

func randId() string {
	b := make([]byte, 2)

	if _, err := src.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}

type DotTranspiler struct {
	dot *strings.Builder
}

func NewDotTranspiler() *DotTranspiler {
	dot := strings.Builder{}
	dot.WriteString("digraph G {\n")
	return &DotTranspiler{dot: &dot}
}

func (t *DotTranspiler) String() string {
	t.dot.WriteRune('}') // FIXME
	return t.dot.String()
}

func (t *DotTranspiler) trans2dot(e interface{ ast.Node }) string {
	switch expr := e.(type) {
	case *ast.Ident:
		return t.transIdent2dot(expr)
	case *ast.CallExpr:
		return t.transCallExpr2dot(expr)
	case *ast.BinaryExpr:
		return t.transBinaryExp2dot(expr)
	case *ast.AssignStmt:
		return t.transAssign2dot(expr)
	case *ast.BasicLit:
		return t.transBasicLit2dot(expr)
	case *ast.ReturnStmt:
		return t.transReturnStmt2dot(expr)
	default:
		log.Printf("WARNING: Don't know how to transpile %v\n", expr)
	}
	return ""
}

func (t *DotTranspiler) connect(from, to string) {
	t.dot.WriteString(fmt.Sprintf("%s -> %s;\n", from, to))
}

func (t *DotTranspiler) connectNoDir(from, to string) {
	t.dot.WriteString(fmt.Sprintf("%s -> %s [dir=none];\n", from, to))
}

func (t *DotTranspiler) transFunc2dot(f *ast.FuncDecl) string {
	// write function name
	funcId := randIdPrefix(f.Name.String())
	t.dot.WriteString(fmt.Sprintf("%s [label=\"%s\", style=filled, fontcolor=white, fillcolor=palevioletred4", funcId, f.Name.String()))

	// write function documentation
	if f.Doc != nil {
		t.dot.WriteString(fmt.Sprintf(", xlabel=\"%s\"", strings.ReplaceAll(f.Doc.Text(), "\"", "\\\"")))
	}
	t.dot.WriteString("];\n")

	// write function parameters
	if f.Type.Params != nil {
		for _, param := range f.Type.Params.List {
			// write param name
			paramName := param.Names[0].String()
			paramNameId := randIdPrefix(paramName)
			t.dot.WriteString(fmt.Sprintf("%s [label=\"%s\", fontsize=10, shape=invtriangle, width=0.5, height=0.5];\n", paramNameId, param.Names[0].String()))

			// write param type
			paramType := fmt.Sprintf("%s", param.Type)
			paramTypeId := randIdPrefix(paramType)
			t.dot.WriteString(fmt.Sprintf("%s [label=\"%s\", fontsize=8, width=0.3, height=0.3];\n", paramTypeId, paramType))

			// connect to param type to param name
			t.connectNoDir(paramTypeId, paramNameId)

			// connect param name to function
			t.connectNoDir(paramNameId, funcId)
		}
	}

	// write function return values
	if f.Type.Results != nil {
		for _, ret := range f.Type.Results.List {
			retType := fmt.Sprintf("%s", ret.Type)
			retId := randIdPrefix(retType)
			t.dot.WriteString(fmt.Sprintf("%s [label=\"%s\", shape=triangle, fontsize=8, width=0.5, height=0.5];\n", retId, retType))

			// connect param name to function
			t.connectNoDir(retId, funcId)
		}
	}

	// write function body
	for _, statement := range f.Body.List {
		t.connect(funcId, t.trans2dot(statement))
	}

	return funcId
}

func (t *DotTranspiler) transReturnStmt2dot(stmt *ast.ReturnStmt) string {
	returnId := randIdPrefix("return")
	t.dot.WriteString(fmt.Sprintf("%s [label=\"return\"];\n", returnId))

	for _, result := range stmt.Results {
		t.connect(returnId, t.trans2dot(result))
	}

	return returnId
}

func (t *DotTranspiler) transBasicLit2dot(lit *ast.BasicLit) string {
	litId := randIdPrefix("lit_" + lit.Value)
	t.dot.WriteString(fmt.Sprintf("%s [label=\"%s\"];\n", litId, lit.Value))

	return litId
}

func (t *DotTranspiler) transCallExpr2dot(callExp *ast.CallExpr) string {
	funcName := fmt.Sprintf("%s", callExp.Fun)
	funcNameId := randIdPrefix(funcName)
	t.dot.WriteString(fmt.Sprintf("%s [label=\"call %s\"];\n", funcNameId, funcName))

	for _, arg := range callExp.Args {
		nodeId := t.trans2dot(arg)
		t.connect(funcNameId, nodeId)
	}

	return funcNameId
}

func (t *DotTranspiler) transBinaryExp2dot(binExp *ast.BinaryExpr) string {
	binOpName := binExp.Op.String()
	binOpId := randIdPrefix("binary_operator")
	t.dot.WriteString(fmt.Sprintf("%s [label=\"%s\"];\n", binOpId, binOpName))

	t.connect(binOpId, t.trans2dot(binExp.X))
	t.connect(binOpId, t.trans2dot(binExp.Y))

	return binOpId
}

func (t *DotTranspiler) transIdent2dot(ident *ast.Ident) string {
	identId := randIdPrefix("id")
	t.dot.WriteString(fmt.Sprintf("%s [label=\"%s\"];\n", identId, ident.String()))

	return identId
}

func (t *DotTranspiler) transAssign2dot(assign *ast.AssignStmt) string {
	assignId := randIdPrefix("assignment")
	t.dot.WriteString(fmt.Sprintf("%s [label=\"=\"];\n", assignId))

	t.connect(assignId, t.trans2dot(assign.Lhs[0]))
	t.connect(assignId, t.trans2dot(assign.Rhs[0]))

	return assignId
}

// transpileFunc2dot transpiles a function with the given name in the given file,
// returns the dot representation
func transpileFunc2dot(file *ast.File, funcName string) string {
	transpiler := NewDotTranspiler()

	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.File: // of course, we need to traverse the file to find the function we're looking for
			return true
		case *ast.FuncDecl:
			if node.Name.Name == funcName {
				transpiler.transFunc2dot(node)
			}
		}
		return false
	})

	return transpiler.String()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ast <filename>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fileBytes, err := os.ReadFile(filePath)
	if err != nil { // always check for errors!
		log.Fatalf("failed to read file %s, underlying error: %w", filePath, err)
	}

	fileSet := token.NewFileSet() // IRL we have to parse multiple files, thus a file set is needed
	fileSet.AddFile(filePath, fileSet.Base(), len(fileBytes))
	parsedFile, err := parser.ParseFile(fileSet, filepath.Base(filePath), fileBytes, parser.ParseComments|parser.Trace)
	if err != nil {
		log.Fatalf("failed to parse file %s, underlying error: %w", filePath, err)
	}

	// print the AST
	// _ = ast.Print(fileSet, parsedFile)

	// traverse the AST and transpile to dot code (only the "funcion" function)
	fmt.Println(transpileFunc2dot(parsedFile, "funcion"))
}
