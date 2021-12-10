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
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}
func (instr *Div) GetSources() []int {
	sources := []int{}
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

	out.WriteString(fmt.Sprintf("    div %s,%s,%s",targetReg,sourceReg1,sourceReg2))

	return out.String()

}

func (instr *Div) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	//// load operand 1
	//source1Offset := funcVarDict[instr.sourceReg1]
	//source1RegId := NextAvailReg()
	//instruction = append(instruction, fmt.Sprintf("ldr x%v, [x29, #%v]", source1RegId, source1Offset))
	//
	//// load operand 2
	//source2RegId := NextAvailReg()
	//source2Offset := funcVarDict[instr.sourceReg2]
	//instruction = append(instruction, fmt.Sprintf("ldr x%v, [x29, #%v]", source2RegId, source2Offset))
	//
	//// divide
	//targetRegId := NextAvailReg()
	//instruction = append(instruction, fmt.Sprintf("sdiv x%v, x%v, x%v", targetRegId, source1RegId, source2RegId))
	//
	//// store result
	//targetOffset := funcVarDict[instr.target]
	//instruction = append(instruction, fmt.Sprintf("str x%v, [x29, #%v]", targetRegId, targetOffset))
	//
	//ReleaseReg(source1RegId)
	//ReleaseReg(source2RegId)
	//ReleaseReg(targetRegId)

	return instruction
}