package arm

import (
	"proj/golite/ir"
	st "proj/golite/symboltable"
)


func TranslateToArm(funcfrags []*ir.FuncFrag, symTable *st.SymbolTable) []string {
	armInstructions := []string{}
	for _, funcfrag := range funcfrags {
		
	}
	return armInstructions
}
