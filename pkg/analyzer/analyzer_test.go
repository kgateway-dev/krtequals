package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestDeepEqualCheckEnabled(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer(&Config{DeepEqual: true})
	analysistest.Run(t, testdata, a, "deepequalon")
}

func TestDeepEqualCheckDisabled(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer(&Config{DeepEqual: false})
	analysistest.Run(t, testdata, a, "deepequaloff")
}

func TestMarkers(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer(&Config{})
	analysistest.Run(t, testdata, a, "markers")
}
