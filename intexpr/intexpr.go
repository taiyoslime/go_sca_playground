package intexpr

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "intexpr is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "intexpr",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case ast.Expr:
			typ := pass.TypesInfo.TypeOf(n)
			if types.Identical(typ, types.Typ[types.Int]) || types.Identical(typ, types.Typ[types.UntypedInt]){
				fmt.Println(pass.Fset.Position(n.Pos()))
			}
		}
	})

	return nil, nil
}

