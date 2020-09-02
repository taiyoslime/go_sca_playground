package main

import (
	"toomanyargs"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(toomanyargs.Analyzer) }

