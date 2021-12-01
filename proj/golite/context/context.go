package context

import (
	"fmt"
)

type CompilerContext struct {
	lexOut     bool   // Determines whether the scanner prints it's output
	astOut     bool   // Determines whether the parser prints it's output (ast)
	iLocOut    bool   // Determines whether to print out ILOC representation
	sourcePath string // The path of the input source file for a golite program
}

func New(lexOut bool, astOut bool, iLocOut bool, sourcePath string) *CompilerContext {
	return &CompilerContext{lexOut, astOut, iLocOut, sourcePath}
}

func (ctx *CompilerContext) SetSourcePath(path string) { ctx.sourcePath = path }
func (ctx *CompilerContext) SetLex(b bool)             { ctx.lexOut = b }
func (ctx *CompilerContext) SetAst(b bool)             { ctx.astOut = b }
func (ctx *CompilerContext) SetILoc(b bool)            { ctx.iLocOut = b }

// OutputLex returns true if the scanner should print-out its output to the user
func (ctx *CompilerContext) OutputLex() bool { return ctx.lexOut }

// OutputAst returns true if the parser should print-out its output to the user
func (ctx *CompilerContext) OutputAst() bool { return ctx.astOut }

// OutputILoc returns true if we want to print out iloc to the user
func (ctx *CompilerContext) OutputILoc() bool { return ctx.iLocOut }

// SourcePath returns the source path for a golite program
func (ctx *CompilerContext) SourcePath() string { return ctx.sourcePath }

// RuntimeError the compiler has put itself in a state that it cannot recover from sa it must exit with an error
func (ctx *CompilerContext) RuntimeError(msg string, e error) {
	if e != nil { // TO-DO
		fmt.Println(msg)
		//os.Exit(2)
	}
}
