package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
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

func (instr *Print) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	utility.SetPrint()
	instruction := []string{}

	var targetRegId int
	var isTargetParam bool
	if targetRegId, isTargetParam = paramRegIds[instr.sourceReg]; !isTargetParam {
		targetRegId = utility.NextAvailReg()
		targetOffset := funcVarDict[instr.sourceReg]
		instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]",targetRegId,targetOffset))
	}

	sourceRegId := utility.NextAvailReg()
	instruction = append(instruction, fmt.Sprintf("\tadrp x%v, .PRINT", sourceRegId))
	instruction = append(instruction, fmt.Sprintf("\tadd x%v,x%v, :lo12:.PRINT",sourceRegId, sourceRegId))
	instruction = append(instruction, fmt.Sprintf("\tmov x1,x%v",targetRegId))
	instruction = append(instruction, fmt.Sprintf("\tmov x0,x%v",sourceRegId))
	instruction = append(instruction, "\tbl printf")

	if !isTargetParam {
		utility.ReleaseReg(targetRegId)
	}
	utility.ReleaseReg(sourceRegId)
	return instruction
}
