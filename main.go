// +build !go1.12

package main

import (
	"github.com/jingyugao/rowserrcheck/passes/rowserr"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(rowserr.NewAnalyzer()) }
