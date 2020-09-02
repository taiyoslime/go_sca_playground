package intexpr_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"intexpr"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, intexpr.Analyzer, "a")
}
