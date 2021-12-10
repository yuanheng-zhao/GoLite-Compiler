package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"proj/golite/arm"
	ct "proj/golite/context"
	"proj/golite/ir"
	ps "proj/golite/parser"
	"proj/golite/sa"
	sc "proj/golite/scanner"
	"strings"
)

// StartCompile starts the compilation process of the compiler
func StartCompile(ctx ct.CompilerContext) []string {
	scanner := sc.New(ctx)
	parser := ps.New(*scanner)
	ast := parser.Parse()
	//fmt.Println(ast)
	globalSymtabl := sa.PerformSA(ast)
	globalFuncFrag := ast.TranslateToILocFunc([]*ir.FuncFrag{}, globalSymtabl)
	//for _, funcFrag := range globalFuncFrag {
	//	instructions := funcFrag.Body
	//	for _, instruction := range instructions {
	//		fmt.Println(instruction.String())
	//	}
	//}

	armInstructString := arm.TranslateToAssembly(globalFuncFrag, globalSymtabl)
	//for _, instruction := range armInstructString {
	//	fmt.Println(instruction)
	//}
	return armInstructString
}

func main() {

	// Define all optional flags for the compiler
	lexOpt := flag.Bool("lex", false, "Send to standard-out the tokens from scanner.")
	astOpt := flag.Bool("ast", false, "Send to standard-out the tokens from parser.")
	ilocOpt := flag.Bool("iloc", false, "Send to standard-out the tokens from IR")
	armOpt := flag.Bool("S", false, "Send to standard-out the tokens of translating to Arm code")
	flag.Parse()
	// Define the usage statement for the compiler
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage of golite: [flags] program.golite  \n")
		flag.PrintDefaults()
		fmt.Fprintf(out, "Usage of golite: [flags] program.golite  \n")
	}

	// Create the compiler configuration struct
	ctx := ct.New(false, false, false, false, "")
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
		ctx.SetILoc(*ilocOpt)
		ctx.SetArm(*armOpt)
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
	} else if ctx.OutputILoc() {
		scanner := sc.New(*ctx)
		parser := ps.New(*scanner)
		ast := parser.Parse()
		//fmt.Println(ast)
		globalSymtabl := sa.PerformSA(ast)
		globalFuncFrag := ast.TranslateToILocFunc([]*ir.FuncFrag{}, globalSymtabl)
		for _, funcFrag := range globalFuncFrag {
			instructions := funcFrag.Body
			for _, instruction := range instructions {
				fmt.Println(instruction.String())
			}
		}
	} else if ctx.OutputArm() {

		baseName := filepath.Base(ctx.SourcePath())
		ext := filepath.Ext(ctx.SourcePath())
		fileName := strings.TrimSuffix(baseName, ext) // get the file name with the extension removed
		fileName = fileName + ".s"

		armCode := StartCompile(*ctx)
		armCodeStr := "" // dump arm code into a string
		for _, line := range armCode {
			armCodeStr = armCodeStr + line + "\n"
		}

		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err2 := f.WriteString(armCodeStr)
		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("Done!")
	}
}
