package ir

import (
	"bytes"
	"fmt"
)

// Add represents a ADD instruction in ILOC
type Add struct{
	target    int        // The target register for the instruction
	sourceReg int        // The first source register of the instruction
	operand   int        // The operand either register or constant
	opty   OperandTy     // The type for the operand (REGISTER, IMMEDIATE)
}

//NewAdd is a constructor and initialization function for a new Add instruction
func NewAdd(target int,sourceReg int, operand int, opty OperandTy ) *Add {
	return &Add{target,sourceReg,operand,opty}
}

func (instr *Add) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *Add) GetSources() []int {
	sources := []int{}
	if instr.opty == REGISTER {
		sources = append(sources, instr.sourceReg, instr.operand)
	} else {
		sources = append(sources, instr.sourceReg)
	}
	return sources
}

func (instr *Add) GetImmediate() *int {

	//Add instruction has two forms for the second operand: register, and immediate (constant)
	//make sure to check for that.
	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	//Return nil if this instruction does not have an immediate
	return nil
}

func (instr *Add) GetSourceString() string {
	return ""
}
func (instr *Add) GetLabel() string {
	// Add does not work with labels so we can just return a default value
	return ""
}
func (instr *Add) SetLabel(newLabel string){
	// Add does not work with labels can we can skip implementing this method.
}

func (instr *Add) String() string {

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

	out.WriteString(fmt.Sprintf("add %s,%s,%s",targetReg,sourceReg,operand2))

	return out.String()

}