package ir

import (
	"bytes"
	"fmt"
)

type Sub struct{
	target    int        // The target register for the instruction
	sourceReg int        // The first source register of the instruction
	operand   int        // The operand either register or constant
	opty   OperandTy     // The type for the operand (REGISTER, IMMEDIATE)
}

func NewSub(target int,sourceReg int, operand int, opty OperandTy ) *Sub {
	return &Sub{target,sourceReg,operand,opty}
}

func (instr *Sub) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *Sub) GetSources() []int {
	sources := []int{}
	if instr.opty == REGISTER {
		sources = append(sources, instr.sourceReg, instr.operand)
	} else {
		sources = append(sources, instr.sourceReg)
	}
	return sources
}

func (instr *Sub) GetImmediate() *int {

	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}

func (instr *Sub) GetSourceString() string {
	return ""
}

func (instr *Sub) GetLabel() string {
	return ""
}

func (instr *Sub) SetLabel(newLabel string){}

func (instr *Sub) String() string {

	var out bytes.Buffer

	targetReg  := fmt.Sprintf("r%v",instr.target)
	sourceReg  := fmt.Sprintf("r%v",instr.sourceReg)

	var prefix string

	if instr.opty == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	operand2   := fmt.Sprintf("%v%v",prefix, instr.operand)

	out.WriteString(fmt.Sprintf("    sub %s,%s,%s",targetReg,sourceReg,operand2))

	return out.String()

}