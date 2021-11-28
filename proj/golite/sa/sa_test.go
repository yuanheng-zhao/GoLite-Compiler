package sa

import (
	"fmt"
	ct "proj/golite/context"
	"proj/golite/parser"
	"proj/golite/scanner"
	"testing"
)

func Test1(t *testing.T) {
	ctx := ct.New(false, false, "test1_sa.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()

	fmt.Println("AST Printout:")
	fmt.Println(ast.String())

	symTable := PerformSA(ast)
	if symTable == nil {
		t.Errorf("\nExpected: returned symbol table; Got nil\n")
	}
}

func Test2(t *testing.T) {
	ctx := ct.New(false, false, "test2_sa.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()

	fmt.Println("AST Printout:")
	fmt.Println(ast.String())

	symTable := PerformSA(ast)
	if symTable != nil {
		t.Errorf("\nExpected: returned nil (errors reported); Got a symbol table\n")
	}
}
