package ir

import (
	"bytes"
	"fmt"
)

type Read struct {
	targetReg int
	variable  string // e.g. "read r5 @b" in benchmarks/simple/simple1/simple1.iloc
}

func NewRead(targetReg int, varName string) *Read {
	return &Read{targetReg, varName}
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
	inst := []string{}

	return inst
}
