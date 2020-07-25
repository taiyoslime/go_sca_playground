package main

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"

	"fmt"
	"log"
	"os"
)

const example = `
package main

import "fmt"

func main() {
	a, b, c, d, e := 2, 3, 10, 5, 6
	f := a + b
	g := c + d + e
	h := a * b * c * d * e
	fmt.Println(f, g, h)
}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", example, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	files := []*ast.File{f}

	pkg := types.NewPackage("example", "")
	ssa, _, err := ssautil.BuildPackage(&types.Config{Importer: importer.Default()}, fset, pkg, files, ssa.SanityCheckFunctions)
	if err != nil {
		log.Fatal(err)
	}
	ssa.Func("main").WriteTo(os.Stdout)
	fmt.Println("ok")
}
