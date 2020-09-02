package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"intexpr"
)

func main() { unitchecker.Main(intexpr.Analyzer) }
