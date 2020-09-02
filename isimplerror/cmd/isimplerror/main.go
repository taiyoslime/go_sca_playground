package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"
	"isimplerror"
)

func main() { unitchecker.Main(isimplerror.Analyzer) }
