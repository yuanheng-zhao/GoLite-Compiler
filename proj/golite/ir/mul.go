package ir

import (
	"bytes"
	"fmt"
)

type Mul struct{
	target    int        // The target register for the instruction
	sourceReg1 int       // The first source register of the instruction
	sourceReg2 int       // The second source register of the instruction
}

func NewMul(target int,sourceReg1 int, sourceReg2 int) *Mul {
	return &Mul{target,sourceReg1,sourceReg2}
}

func (instr *Mul) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}
func (instr *Mul) GetSources() []int {
	sources := []int{}
	sources = append(sources, instr.sourceReg1, instr.sourceReg2)
	return sources
}
func (instr *Mul) GetImmediate() *int {

	//Return nil if this instruction does not have an immediate
	return nil
}
func (instr *Mul) GetSourceString() string {
	return ""
}
func (instr *Mul) GetLabel() string {
	return ""
}
func (instr *Mul) SetLabel(newLabel string){}

func (instr *Mul) String() string {

	var out bytes.Buffer

	targetReg  := fmt.Sprintf("r%v",instr.target)
	sourceReg1 := fmt.Sprintf("r%v",instr.sourceReg1)
	sourceReg2 := fmt.Sprintf("r%v",instr.sourceReg2)

	out.WriteString(fmt.Sprintf("mul %s,%s,%s",targetReg,sourceReg1,sourceReg2))

	return out.String()

}