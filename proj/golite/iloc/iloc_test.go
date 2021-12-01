package iloc

import (
	"fmt"
	ct "proj/golite/context"
	"proj/golite/ir"
	"proj/golite/parser"
	"proj/golite/sa"
	"proj/golite/scanner"
	"testing"
)

func Test1(t *testing.T) {
	ctx := ct.New(false, false, false, "test1_iloc.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	if globalSymTable == nil {
		t.Errorf("\nExpected: returned symbol table; Got nil\n")
	}

	globalFuncFrag := ast.TranslateToILocFunc([]*ir.FuncFrag{}, globalSymTable)
	if globalFuncFrag == nil {
		t.Errorf("\nExpected: returned FuncFrag; Got nil\n")
	}
	for _, funcFrag := range globalFuncFrag {
		instructions := funcFrag.Body
		for _, instruction := range instructions {
			fmt.Println(instruction.String())
		}
	}
}

func Test2(t *testing.T) {
	ctx := ct.New(false, false, false, "test2_iloc.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	if globalSymTable == nil {
		t.Errorf("\nExpected: returned symbol table; Got nil\n")
	}

	globalFuncFrag := ast.TranslateToILocFunc([]*ir.FuncFrag{}, globalSymTable)
	if globalFuncFrag == nil {
		t.Errorf("\nExpected: returned FuncFrag; Got nil\n")
	}
	for _, funcFrag := range globalFuncFrag {
		instructions := funcFrag.Body
		for _, instruction := range instructions {
			fmt.Println(instruction.String())
		}
	}
}

func Test3(t *testing.T) {
	ctx := ct.New(false, false, false, "test3_iloc.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	if globalSymTable == nil {
		t.Errorf("\nExpected: returned symbol table; Got nil\n")
	}

	globalFuncFrag := ast.TranslateToILocFunc([]*ir.FuncFrag{}, globalSymTable)
	if globalFuncFrag == nil {
		t.Errorf("\nExpected: returned FuncFrag; Got nil\n")
	}
	for _, funcFrag := range globalFuncFrag {
		instructions := funcFrag.Body
		for _, instruction := range instructions {
			fmt.Println(instruction.String())
		}
	}
}
