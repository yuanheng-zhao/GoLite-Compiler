package arm

import (
	"fmt"
	"proj/golite/ir"
	st "proj/golite/symboltable"
)

func TranslateToArm(funcfrags []*ir.FuncFrag, symTable *st.SymbolTable) []string {
	armInstructions := []string{}
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
	}
	return armInstructions
}
