package bodyclose_test

import (
	"testing"

	"github.com/timakin/bodyclose/passes/bodyclose"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, bodyclose.Analyzer, "a")
}
