package arm

import (
	"fmt"
	"proj/golite/ir"
	st "proj/golite/symboltable"
)

func TranslateToAssembly(funcfrags []*ir.FuncFrag, symTable *st.SymbolTable) []string {

	armInstructions := []string{}

	// program title
	armInstructions = append(armInstructions, ".arch armv8-a")
	// global variables

	// code
	armInstructions = append(armInstructions, ".text")

	for _, funcfrag := range funcfrags {
		funcEntry := symTable.Contains(funcfrag.Label)
		scopeSt := funcEntry.GetScopeST() // symbol table for the current scope
		countVar := 0                     // count the variables inside this scope
		for key, e := range scopeSt.HashTable() {
			entry := *e
			countVar += 1
			fmt.Println(key, entry.GetEntryType().GetName())
		}
		fmt.Println(countVar)

		armInstructions = append(armInstructions, "\t.type " + funcfrag.Label + ",%function")
		armInstructions = append(armInstructions, "\t.global " + funcfrag.Label)
		armInstructions = append(armInstructions, "\t.p2align\t\t2")

		for _, instruction := range funcfrag.Body {

		}


	}

	return armInstructions
}
