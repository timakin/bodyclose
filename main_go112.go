// +build go1.12

package main

import (
	"github.com/timakin/bodyclose/passes/bodyclose"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

// Analyzers returns analyzers of bodyclose.
func analyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		bodyclose.Analyzer,
	}
}

func main() {
	unitchecker.Main(analyzers()...)
}
