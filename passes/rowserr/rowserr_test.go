package rowserr_test

import (
	"log"
	"testing"

	"github.com/jingyugao/rowserrcheck/passes/rowserr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, rowserr.NewAnalyzer(), "a")
}
