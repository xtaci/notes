package main

import (
	"github.com/davecgh/go-spew/spew"
	//"go/ast"
	"go/parser"
)

func main() {
	cmd := `func(){a,b :=1}`
	expr, err := parser.ParseExpr(cmd)
	spew.Dump(expr, err)
}
