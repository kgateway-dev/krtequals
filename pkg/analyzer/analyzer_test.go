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

func TestChainedExpressions(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer(&Config{})
	analysistest.Run(t, testdata, a, "chained")
}

func TestDelegatedEquals(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer(&Config{})
	analysistest.Run(t, testdata, a, "delegated")
}

func TestMethodArgs(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer(&Config{})
	analysistest.Run(t, testdata, a, "methodargs")
}

func TestUnexportedCheckEnabled(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer(&Config{CheckUnexported: true})
	analysistest.Run(t, testdata, a, "unexportedon")
}

func TestUnexportedCheckDisabled(t *testing.T) {
	testdata := analysistest.TestData()
	a := NewAnalyzer(&Config{CheckUnexported: false})
	analysistest.Run(t, testdata, a, "unexportedoff")
}
