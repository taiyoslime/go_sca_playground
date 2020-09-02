package snpkgval_test

import (
	"testing"

	"snpkgval"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, snpkgval.Analyzer, "a")
}

