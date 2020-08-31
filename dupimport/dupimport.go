package dupimport

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "dupimport is to detect dupulicated imports in a file"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "dupimport",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		hmap := map[string]struct{}{}
		for _, isp := range file.Imports {
			path, err := strconv.Unquote(isp.Path.Value)
			if err != nil {
				return nil, err
			}
			_, ok := hmap[path]
			if ok {
				pass.Reportf(isp.Pos(), "duplicated import: %s", path)
			} else {
				hmap[path] = struct{}{}
			}
		}
	}
	return nil, nil
}
