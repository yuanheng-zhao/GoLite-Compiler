package sa

import (
	"flag"
	"fmt"
	"proj/golite/ast"
	st "proj/golite/symboltable"
)

// return true if there exists any error
//func reportErrors(errors []string) bool {
//	return !(len(errors) == 0)
//}

func reportErrors(errors []string) bool {
	// return true if there exists any error
	if len(errors) > 0 {
		for _, err := range errors {
			out := flag.CommandLine.Output()
			fmt.Fprintf(out, "semantic error: %s}\n", err)
		}
		return true
	}
	return false
}

func PerformSA(program *ast.Program) *st.SymbolTable {
	// Define a new global table
	globalST := st.New(nil, "global")
	errors := make([]string, 0)

	// First Build the Symbol Table(s) for all declarations
	errors = program.PerformSABuild(errors, globalST)

	// Report errors
	if !reportErrors(errors) {
		// second perform type checking
		errors := make([]string, 0)
		errors = program.TypeCheck(errors, globalST)
		if !reportErrors(errors) { // finally, no error
			return globalST
		}
	}
	return nil
}
