package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// PrintMax prints the maximum of two numbers.
// The two values must be integers.
func PrintMax(x, y int) {
	if x > y {
		fmt.Printf("%d is maximum\n", x)
	} else {
		fmt.Printf("%d is maximum\n", y)
	}
}

// Get the name and path of a func
func funcPathAndName(f interface{}) string {
	s := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fmt.Println("+++ funcPathAndName:", s)
	return s
}

// Get the name of a func (with package path)
func funcName(f interface{}) string {
	splitFuncName := strings.Split(funcPathAndName(f), ".")
	s := splitFuncName[len(splitFuncName)-1]
	fmt.Println("+++ funcName:", s)
	return s
}

// Get description of a func
func funcDescription(f interface{}) string {
	fileName, _ := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).FileLine(0)
	funcName := funcName(f)
	fset := token.NewFileSet()

	// Parse src
	parsedAst, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	pkg := &ast.Package{
		Name:  "Any",
		Files: make(map[string]*ast.File),
	}
	pkg.Files[fileName] = parsedAst

	importPath, _ := filepath.Abs("/")
	myDoc := doc.New(pkg, importPath, doc.AllDecls)
	for _, theFunc := range myDoc.Funcs {
		if theFunc.Name == funcName {
			return theFunc.Doc
		}
	}
	return ""
}

func main() {
	fmt.Println("--- calling PrintMax with 3 and 5 as arguments")
	PrintMax(3, 5)
	fmt.Println("--- print PrintMax docstring")
	fmt.Println(funcDescription(PrintMax))
}
