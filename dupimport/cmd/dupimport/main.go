package main

import (
	"dupimport"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(dupimport.Analyzer) }

