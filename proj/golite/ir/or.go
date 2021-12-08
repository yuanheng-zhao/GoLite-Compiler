package ir

import (
	"bytes"
	"fmt"
)

type Or struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	operand   int       // The operand either register or constant
	opty      OperandTy // The type for the operand (REGISTER, IMMEDIATE)
}

func NewOr(target int, sourceReg int, operand int, opty OperandTy) *Or {
	return &Or{target, sourceReg, operand, opty}
}

func (instr *Or) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *Or) GetSources() []int {
	sources := []int{}
	if instr.opty != IMMEDIATE {
		sources = append(sources, instr.sourceReg, instr.operand)
	} else {
		sources = append(sources, instr.sourceReg)
	}
	return sources
}

func (instr *Or) GetImmediate() *int {

	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}

func (instr *Or) GetSourceString() string {
	return ""
}

func (instr *Or) GetLabel() string {
	return ""
}

func (instr *Or) SetLabel(newLabel string) {}

func (instr *Or) String() string {

	var out bytes.Buffer

	targetReg := fmt.Sprintf("r%v", instr.target)
	sourceReg := fmt.Sprintf("r%v", instr.sourceReg)

	var prefix string

	if instr.opty == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	operand2 := fmt.Sprintf("%v%v", prefix, instr.operand)

	out.WriteString(fmt.Sprintf("    or %s,%s,%s", targetReg, sourceReg, operand2))

	return out.String()

}

func (instr *Or) TranslateToAssembly(funcVarDict map[int]int) []string {
	inst := []string{}

	return inst
}
