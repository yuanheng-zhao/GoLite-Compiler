package arm

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
	ctx := ct.New(false, false, false, "test2_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	mainEnt := globalSymTable.Contains("main")
	fmt.Println(mainEnt)
	if globalSymTable == nil {
		t.Errorf("\nExpected: returned symbol table; Got nil\n")
	}

	globalFuncFrag := ast.TranslateToILocFunc([]*ir.FuncFrag{}, globalSymTable)
	if globalFuncFrag == nil {
		t.Errorf("\nExpected: returned FuncFrag; Got nil\n")
	}
	//for _, funcFrag := range globalFuncFrag {
	//	instructions := funcFrag.Body
	//	for _, instruction := range instructions {
	//		fmt.Println(instruction.String())
	//	}
	//}
	resStr := TranslateToAssembly(globalFuncFrag, globalSymTable)
	fmt.Println(resStr)
}
