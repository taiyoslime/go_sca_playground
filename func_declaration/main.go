package main

import (
	"go/ast"
	"go/parser"
	"go/token"

	"fmt"
	"log"
)

const example = `
package main

import "fmt"

func ex1(x int) int {
	return x
}

func ex2(x, y int) int {
	return x + y
}

type Type int

type Interface interface {
	ex3() int
}

func (a Type) ex3() int {
	return 1
}

func ex4(x int) (int, int) {
	return x, x
}

func ex5() (x int) {
	x = 1
	return
}

func main() {
	var x Interface = Type(1)
	y, z := ex4(1)
	a := ex1(1) + ex2(1, 1) + x.ex3() + y + z + ex5() + ex6(1)
	fmt.Println(a)
}

`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", example, parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}

	/*
		for _, d := range f.Decls {
			ast.Print(fset, d)
			fmt.Println()
		}
	*/

	ast.Inspect(f, func(n ast.Node) bool {
		decl, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		funcName := decl.Name
		fmt.Printf("> function name\n%s\n", funcName)

		params := decl.Type.Params
		fmt.Println("> arguments")
		if params != nil && params.List != nil {
			for i := range params.List {
				for _, name := range params.List[i].Names {
					fmt.Printf("%s ", name)
				}
				fmt.Printf(": %s\n", params.List[i].Type)
			}
		} else {
			fmt.Println("None")
		}

		recv := decl.Recv
		if recv != nil && recv.List != nil {
			fmt.Println("> receiver")
			for i := range recv.List {
				for _, name := range recv.List[i].Names {
					fmt.Printf("%s ", name)
				}
				fmt.Printf(": %s\n", recv.List[i].Type)
			}
		}

		results := decl.Type.Results
		fmt.Println("> return values")
		if results != nil && results.List != nil {
			for i := range results.List {
				for _, name := range results.List[i].Names {
					fmt.Printf("%s ", name)
				}
				if len(results.List[i].Names) == 0 {
					fmt.Print("(Anonymus)")
				}
				fmt.Printf(": %s\n", results.List[i].Type)
			}
		} else {
			fmt.Println("None")
		}

		fmt.Println()

		return true
	})
}
