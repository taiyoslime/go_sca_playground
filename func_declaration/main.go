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
		params := decl.Type.Params
		recv := decl.Recv
		results := decl.Type.Results

		if recv != nil {
			fmt.Printf("> method name\n%s\n", funcName)
		} else {
			fmt.Printf("> function name\n%s\n", funcName)
		}

		printFieldList := func(list *ast.FieldList) {
			if list != nil && list.List != nil {
				for i := range list.List {
					for _, name := range list.List[i].Names {
						fmt.Printf("%s ", name)
					}
					if len(list.List[i].Names) == 0 {
						fmt.Print("(Anonymus)")
					}
					fmt.Printf(": %s\n", list.List[i].Type)
				}
			} else {
				fmt.Println("None")
			}
		}

		fmt.Println("> arguments")
		printFieldList(params)

		if recv != nil {
			fmt.Println("> receivers")
			printFieldList(recv)
		}

		fmt.Println("> return values")
		printFieldList(results)

		fmt.Println()

		return true
	})
}
