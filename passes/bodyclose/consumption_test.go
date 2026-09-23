package bodyclose_test

import (
	"testing"

	"github.com/timakin/bodyclose/passes/bodyclose"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestConsumption(t *testing.T) {
	// Create analyzer with consumption flag enabled
	analyzer := *bodyclose.Analyzer // Copy the analyzer
	if err := analyzer.Flags.Set("check-consumption", "true"); err != nil {
		t.Fatal(err)
	}

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, &analyzer, "consumption")
}
