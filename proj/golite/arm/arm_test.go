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
	ctx := ct.New(false, false, false, false, "test1_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test6(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test6_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test7(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test7_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test8(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test8_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test9(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test9_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test10(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test10_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test11(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test11_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test12(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test12_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test13(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test13_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test14(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test14_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test15(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test15_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test16(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test16_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}

func Test17(t *testing.T) {
	ctx := ct.New(false, false, false, false, "test17_arm.golite")
	myScanner := scanner.New(*ctx)
	myParser := parser.New(*myScanner)
	ast := myParser.Parse()
	//fmt.Println("AST Printout:")
	//fmt.Println(ast.String())

	globalSymTable := sa.PerformSA(ast)
	//mainEnt := globalSymTable.Contains("main")
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
	for _, line := range resStr {
		fmt.Println(line)
	}
}