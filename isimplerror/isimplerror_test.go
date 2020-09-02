package isimplerror_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"isimplerror"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, isimplerror.Analyzer, "a")
}
