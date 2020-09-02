package toomanyargs

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "toomanyargs is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "toomanyargs",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		reportIf := func(t *ast.FuncType) {
			if t.Params.NumFields() >= 5 {
				pass.Reportf(t.Pos(), "too many arguments")
			}
		}
		switch n := n.(type) {
		case *ast.FuncDecl:
			reportIf(n.Type)
		case *ast.FuncLit:
			reportIf(n.Type)
		}
	})

	return nil, nil
}

