package isimplerror

import (
	"go/types"
	"golang.org/x/tools/go/analysis"
)

const doc = "isimplerror judges whether the type implements error interface."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "isimplerror",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	errorTyp := types.Universe.Lookup("error").Type().Underlying().(*types.Interface)
	for _, name := range pass.Pkg.Scope().Names() {
		obj := pass.Pkg.Scope().Lookup(name)
		if types.Implements(obj.Type(), errorTyp) || types.Implements(types.NewPointer(obj.Type()), errorTyp) {
			pass.Reportf(obj.Pos(), "OK")
		}
	}
	return nil, nil
}
