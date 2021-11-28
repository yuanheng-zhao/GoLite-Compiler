package main

import (
	"flag"
	"fmt"
	"os"
	ct "proj/golite/context"
	ps "proj/golite/parser"
	"proj/golite/sa"
	sc "proj/golite/scanner"
)

// StartCompile starts the compilation process of the compiler
func StartCompile(ctx ct.CompilerContext) {
	scanner := sc.New(ctx)
	parser := ps.New(*scanner)
	ast := parser.Parse()
	sa.PerformSA(ast)
	fmt.Println(ast)
}

func main() {

	// Define all optional flags for the compiler
	lexOpt := flag.Bool("lex", false, "Send to standard-out the tokens from scanner.")
	astOpt := flag.Bool("ast", false, "Send to standard-out the tokens from parser.")
	flag.Parse()
	// Define the usage statement for the compiler
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage of golite: [flags] program.golite  \n")
		flag.PrintDefaults()
		fmt.Fprintf(out, "Usage of golite: [flags] program.golite  \n")
	}

	// Create the compiler configuration struct
	ctx := ct.New(false, false, "")
	//fmt.Println("os.Args    :", os.Args)
	//fmt.Println("flag.Args():", flag.Args())

	// Verify that the user provided the input source file
	if flag.NArg() < 1 {
		flag.Usage()
		return
	} else {
		// The sourcePath is always the first argument from the remaining arguments on the command line
		ctx.SetSourcePath(flag.Arg(0))
		ctx.SetLex(*lexOpt)
		ctx.SetAst(*astOpt)
	}

	// Check if the source file path exists
	if _, err := os.Stat(ctx.SourcePath()); err != nil {
		panic(err)
	}

	// TO-DO : Uncomment if ready to compile a file
	//StartCompile(ctx)
	if ctx.OutputLex() {
		scanner := sc.New(*ctx)
		scanner.Tokens()
	} else if ctx.OutputAst() {
		scanner := sc.New(*ctx)
		parser := ps.New(*scanner)
		ast := parser.Parse()
		fmt.Println(ast.String())
	}
}
