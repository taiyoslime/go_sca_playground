package main

import (
	"intexpr"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(intexpr.Analyzer) }

