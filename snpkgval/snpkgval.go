package snpkgval

import (
	"go/types"
	"golang.org/x/tools/go/analysis"
)

const doc = "snpkgval is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "snpkgval",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, name := range pass.Pkg.Scope().Names() {
		obj := pass.Pkg.Scope().Lookup(name)
		if _, ok := obj.(*types.Var); ok && len(name) <= 1 {
			pass.Reportf(obj.Pos(), "NG")
		}
	}
	return nil, nil
}

