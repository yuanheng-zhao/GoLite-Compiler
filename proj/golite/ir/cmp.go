package ir

import (
	"bytes"
	"fmt"
)

type Cmp struct {
	sourceReg int
	operand   int
	opty      OperandTy
}

func NewCmp(sourceReg int, operand int, opty OperandTy) *Cmp {
	return &Cmp{sourceReg, operand, opty}
}

func (instr *Cmp) GetTargets() []int {
	targets := []int{}
	return targets
}

func (instr *Cmp) GetSources() []int {
	sources := []int{}
	if instr.opty == REGISTER {
		sources = append(sources, instr.sourceReg, instr.operand)
	} else if instr.opty == IMMEDIATE {
		sources = append(sources, instr.sourceReg)
	}
	return sources
}

func (instr *Cmp) GetImmediate() *int {
	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}

func (instr *Cmp) GetSourceString() string { return ""}

func (instr *Cmp) GetLabel() string { return "" }

func (instr *Cmp) SetLabel(newLabel string) {}

func (instr *Cmp) String() string {
	var out bytes.Buffer

	sourceReg := fmt.Sprintf("r%v",instr.sourceReg)

	var prefix string

	if instr.opty == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}

	sourceOp2 := fmt.Sprintf("%v%v", prefix, instr.operand)

	out.WriteString(fmt.Sprintf("    cmp %s,%s", sourceReg,sourceOp2))

	return out.String()
}