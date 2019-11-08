// +build go1.12

package main

import (
	"github.com/jingyugao/rowserrcheck/passes/rowserr"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

// Analyzers returns analyzers of rowserr.
func analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		rowserr.NewAnalyzer(),
	}
}

func main() {
	unitchecker.Main(analyzers()...)
}
