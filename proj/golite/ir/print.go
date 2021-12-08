package ir

import (
	"bytes"
	"fmt"
)

type Print struct {
	sourceReg int
}

func NewPrint(sourceReg int) *Print {
	return &Print{sourceReg}
}

func (instr *Print) GetTargets() []int { return []int{} }

func (instr *Print) GetSources() []int {
	source := []int{}
	source = append(source, instr.sourceReg)
	return source
}

func (instr *Print) GetImmediate() *int { return nil }

func (instr *Print) GetSourceString() string { return "" }

func (instr *Print) GetLabel() string { return "" }

func (instr *Print) SetLabel(newLabel string) {}

func (instr *Print) String() string {
	var out bytes.Buffer
	sourceRegister := fmt.Sprintf("r%v", instr.sourceReg)
	out.WriteString(fmt.Sprintf("    print %s", sourceRegister))
	return out.String()
}

func PrintArmFormat() []string {
	printInst := []string{}
	printInst = append(printInst, ".PRINT:")
	printInst = append(printInst, "\t.asciz\t\"%ld\"")
	printInst = append(printInst, "\t.size\t.PRINT, 4")
	return printInst
}

func (instr *Print) TranslateToAssembly(funcVarDict map[int]int) []string {
	inst := []string{}

	return inst
}
