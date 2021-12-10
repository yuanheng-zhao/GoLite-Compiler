package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
)

type Read struct {
	targetReg int
	variable  string // e.g. "read r5 @b" in benchmarks/simple/simple1/simple1.iloc
	varReg    int
}

func NewRead(targetReg int, varName string, varReg int) *Read {
	return &Read{targetReg, varName, varReg}
}

func (instr *Read) GetTargets() []int {
	target := []int{}
	target = append(target, instr.targetReg)
	return target
}

func (instr *Read) GetSources() []int { return []int{} }

func (instr *Read) GetImmediate() *int { return nil }

func (instr *Read) GetSourceString() string { return "" }

func (instr *Read) GetLabel() string { return "" }

func (instr *Read) SetLabel(newLabel string) {}

func (instr *Read) String() string {
	var out bytes.Buffer

	targetRegister := fmt.Sprintf("r%v", instr.targetReg)
	out.WriteString(fmt.Sprintf("    read %s @%v", targetRegister, instr.variable))
	return out.String()
}

func ReadArmFormat() []string {
	readInsts := []string{}
	readInsts = append(readInsts, ".READ")
	readInsts = append(readInsts, "\t.asciz\t\"%ld\"")
	readInsts = append(readInsts, "\t.size\t.READ, 4")
	return readInsts
}

func (instr *Read) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	utility.SetScan()
	instruction := []string{}

	varTargetRegId := utility.NextAvailReg()
	varTargetOffset := funcVarDict[instr.varReg]
	instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]",varTargetRegId, varTargetOffset))

	sourceReg := utility.NextAvailReg()
	instruction = append(instruction, fmt.Sprintf("\tadrp x%v, .READ",sourceReg))
	instruction = append(instruction, fmt.Sprintf("\tadd x%v,x%v,:lo12:.READ",sourceReg,sourceReg))

	instruction = append(instruction, fmt.Sprintf("\tadd x%v,x29,#%v",varTargetRegId,varTargetOffset))
	instruction = append(instruction, fmt.Sprintf("\tmov x0,x%v", sourceReg))
	instruction = append(instruction, "\tbl scanf")

	return instruction
}
