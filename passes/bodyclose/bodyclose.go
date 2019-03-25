package bodyclose

import (
	"fmt"
	"go/ast"
	"go/types"
	"log"
	"strconv"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

var Analyzer = &analysis.Analyzer{
	Name: "bodyclose",
	Doc:  Doc,
	Run:  new(runner).run,
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
	resObj    types.Object
	resTyp    *types.Pointer
	bodyObj   types.Object
	closeMthd *types.Func
	skipFile  map[*ast.File]bool
}

func (r *runner) run(pass *analysis.Pass) (interface{}, error) {
	r.pass = pass
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs

	r.resObj = analysisutil.LookupFromImports(pass.Pkg.Imports(), nethttpPath, "Response")
	if r.resObj == nil {
		// skip checking
		return nil, nil
	}

	resNamed, ok := r.resObj.Type().(*types.Named)
	if !ok {
		return nil, fmt.Errorf("cannot find http.Response")
	}
	r.resTyp = types.NewPointer(resNamed)

	resStruct, ok := r.resObj.Type().Underlying().(*types.Struct)
	if !ok {
		return nil, fmt.Errorf("cannot find http.Response")
	}
	for i := 0; i < resStruct.NumFields(); i++ {
		field := resStruct.Field(i)
		switch field.Id() {
		case "Body":
			r.bodyObj = field
		}
	}
	if r.bodyObj == nil {
		return nil, fmt.Errorf("cannot find the object http.Response.Body")
	}
	bodyNamed := r.bodyObj.Type().(*types.Named)
	bodyItrf := bodyNamed.Underlying().(*types.Interface)
	for i := 0; i < bodyItrf.NumMethods(); i++ {
		bmthd := bodyItrf.Method(i)
		switch bmthd.Id() {
		case "Close":
			r.closeMthd = bmthd
		}
	}

	r.skipFile = map[*ast.File]bool{}
	for _, f := range funcs {
		if r.noImportedNetHTTP(f) {
			// skip this
			continue
		}

		for _, b := range f.Blocks {
			if !r.containsRes(b) {
				continue
			}
			for i := range b.Instrs {
				pos := b.Instrs[i].Pos()
				if r.isopen(b, i) {
					for _, instr := range b.Instrs {
						log.Printf("%+v, parent: %+v", instr, instr.Parent())
					}

					pass.Reportf(pos, "response body must be closed")
				}
			}
		}
	}

	return nil, nil
}

func (r *runner) containsRes(b *ssa.BasicBlock) bool {
	for _, instr := range b.Instrs {
		switch instr := instr.(type) {
		case ssa.Value:
			if instr.Type().String() == r.resTyp.String() {
				return true
			}
		}
	}
	return false
}

func (r *runner) isopen(b *ssa.BasicBlock, i int) bool {
	for _, instr := range b.Instrs {
		switch instr := instr.(type) {
		case ssa.Value:
			if instr.Type().String() != r.resTyp.String() {
				continue
			}
			if instr.Referrers() == nil {
				continue
			}
			resRefs := *instr.Referrers()
			for _, resRef := range resRefs {
				b := resRef.(*ssa.FieldAddr)
				if b.Referrers() == nil {
					continue
				}

				bRefs := *b.Referrers()
				for _, bRef := range bRefs {
					bOp := bRef.(*ssa.UnOp)
					if bOp.Type() != r.bodyObj.Type() {
						continue
					}

					if bOp.Referrers() == nil {
						continue
					}
					ccalls := *bOp.Referrers()
					for _, ccall := range ccalls {
						switch ccall := ccall.(type) {
						case *ssa.Defer:
							if ccall.Call.Method.Name() == r.closeMthd.Name() {
								return false
							}
						case *ssa.Call:
							if ccall.Call.Method.Name() == r.closeMthd.Name() {
								return false
							}
						}
					}
				}
			}
		}
	}
	return true
}

func (r *runner) noImportedNetHTTP(f *ssa.Function) (ret bool) {
	obj := f.Object()
	if obj == nil {
		return false
	}

	file := analysisutil.File(r.pass, obj.Pos())
	if file == nil {
		return false
	}

	if skip, has := r.skipFile[file]; has {
		return skip
	}
	defer func() {
		r.skipFile[file] = ret
	}()

	for _, impt := range file.Imports {
		path, err := strconv.Unquote(impt.Path.Value)
		if err != nil {
			continue
		}
		path = analysisutil.RemoveVendor(path)
		if path == nethttpPath {
			return false
		}
	}

	return true
}
