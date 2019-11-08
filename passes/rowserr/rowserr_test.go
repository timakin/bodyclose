package rowserr_test

import (
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/jingyugao/rowserrcheck/passes/rowserr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	_ = sqlx.DB{}
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, rowserr.NewAnalyzer(), "a")
}
