package sa

import (
	"fmt"
	ct "proj/golite/context"
	"proj/golite/parser"
	"proj/golite/scanner"
	"testing"
)

func Test1(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test1_sa.golite")
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
	ctx := ct.New(false, false, false, false, "test2_sa.golite")
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

func Test3(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test3_sa.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()

	fmt.Println("AST Printout:")
	fmt.Println(ast.String())

	symTable := PerformSA(ast)
	if symTable == nil {
		t.Errorf("\nExpected: returned symbol table; Got nil\n")
	}

	//mainSt := symTable.Contains("main")
	//aEnt := mainSt.GetScopeST().Contains("b")
	//xEnt := aEnt.GetScopeST().Contains("x")
	//fmt.Println(xEnt) // Check the values by debugging
}

func Test4(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test4_sa.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()

	fmt.Println("AST Printout:")
	fmt.Println(ast.String())

	symTable := PerformSA(ast)
	if symTable == nil {
		t.Errorf("\nExpected: returned symbol table; Got nil\n")
	}

	//mainSt := symTable.Contains("main")
	//xEnt := mainSt.GetScopeST().Contains("x")
	//aEnt := xEnt.GetScopeST().Contains("a")
	//fmt.Println(aEnt) // Check the values by debugging
}
