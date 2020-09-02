package main

import (
	"snpkgval"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(snpkgval.Analyzer) }

