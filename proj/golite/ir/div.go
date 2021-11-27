package ir

import (
	"bytes"
	"fmt"
)

type Div struct{
	target    int        // The target register for the instruction
	sourceReg1 int       // The first source register of the instruction
	sourceReg2 int       // The second source register of the instruction
}

func NewDiv(target int,sourceReg1 int, sourceReg2 int) *Div	 {
	return &Div{target,sourceReg1,sourceReg2}
}

func (instr *Div) GetTargets() []int {
	targets := make([]int, 1)
	targets = append(targets, instr.target)
	return targets
}
func (instr *Div) GetSources() []int {
	sources := make([]int, 2)
	sources = append(sources, instr.sourceReg1, instr.sourceReg2)
	return sources
}
func (instr *Div) GetImmediate() *int {

	//Return nil if this instruction does not have an immediate
	return nil
}
func (instr *Div) GetSourceString() string {
	return ""
}
func (instr *Div) GetLabel() string {
	return ""
}
func (instr *Div) SetLabel(newLabel string){}

func (instr *Div) String() string {

	var out bytes.Buffer

	targetReg  := fmt.Sprintf("r%v",instr.target)
	sourceReg1 := fmt.Sprintf("r%v",instr.sourceReg1)
	sourceReg2 := fmt.Sprintf("r%v",instr.sourceReg2)

	out.WriteString(fmt.Sprintf("div %s,%s,%s",targetReg,sourceReg1,sourceReg2))

	return out.String()

}