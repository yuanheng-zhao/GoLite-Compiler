package ir

import (
	"bytes"
	"fmt"
)

type Not struct{
	target    int        // The target register for the instruction
	sourceReg int        // The first source register of the instruction
	operand   int        // The operand either register or constant
	opty   OperandTy     // The type for the operand (REGISTER, IMMEDIATE)
}

func NewNot(target int,sourceReg int, operand int, opty OperandTy ) *Not {
	return &Not{target,sourceReg,operand,opty}
}

func (instr *Not) GetTargets() []int {
	targets := make([]int, 1)
	targets = append(targets, instr.target)
	return targets
}

func (instr *Not) GetSources() []int {

	if instr.opty != IMMEDIATE {
		sources := make([]int, 1)
		sources = append(sources, instr.sourceReg)
		return sources
	}
	return nil
}

func (instr *Not) GetImmediate() *int {

	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}

func (instr *Not) GetSourceString() string {
	return ""
}

func (instr *Not) GetLabel() string {
	return ""
}

func (instr *Not) SetLabel(newLabel string){}

func (instr *Not) String() string {

	var out bytes.Buffer

	targetReg  := fmt.Sprintf("r%v",instr.target)

	var prefix string
	var operand2 string

	if instr.opty == IMMEDIATE {
		prefix = "#"
		operand2   = fmt.Sprintf("%v%v",prefix, instr.operand)
	} else {
		prefix = "r"
		operand2   = fmt.Sprintf("%v%v",prefix, instr.sourceReg)
	}

	out.WriteString(fmt.Sprintf("sub %s,%s,%s",targetReg,operand2))

	return out.String()

}