package bodyclose

import (
	"go/types"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

var Analyzer = &analysis.Analyzer{
	Name: "bodyclose",
	Doc:  Doc,
	//Run:  new(runner).run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

const (
	Doc = "bodyclose checks whether HTTP response body is closed successfully"

	nethttpPath = "net/http"
)


type runner struct {
	pass      *analysis.Pass
	iterObj   types.Object
	iterNamed *types.Named
	iterTyp   *types.Pointer
	closeMthd  *types.Func
	skipFile  map[*ast.File]bool
}
