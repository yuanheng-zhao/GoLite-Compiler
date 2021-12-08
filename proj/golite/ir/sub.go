package ir

import (
	"bytes"
	"fmt"
)

type Sub struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	operand   int       // The operand either register or constant
	opty      OperandTy // The type for the operand (REGISTER, IMMEDIATE)
}

func NewSub(target int, sourceReg int, operand int, opty OperandTy) *Sub {
	return &Sub{target, sourceReg, operand, opty}
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

func (instr *Sub) SetLabel(newLabel string) {}

func (instr *Sub) String() string {

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

	out.WriteString(fmt.Sprintf("    sub %s,%s,%s", targetReg, sourceReg, operand2))

	return out.String()

}

func (instr *Sub) TranslateToAssembly(funcVarDict map[int]int) []string {
	instruction := []string{}

	// load operand 1
	sourceOffset := funcVarDict[instr.sourceReg]
	sourceRegId := NextAvailReg()
	instruction = append(instruction, fmt.Sprintf("ldr x%v, [x29, #%v]", sourceRegId, sourceOffset))

	// load operand 2
	source2RegId := NextAvailReg()
	if instr.opty == REGISTER {
		source2Offset := funcVarDict[instr.operand]
		instruction = append(instruction, fmt.Sprintf("ldr x%v, [x29, #%v]", source2RegId, source2Offset))
	} else {
		instruction = append(instruction, fmt.Sprintf("mov x%v, #%v", source2RegId, instr.operand))
	}

	// subtract
	targetRegId := NextAvailReg()
	instruction = append(instruction, fmt.Sprintf("add x%v, x%v, x%v", targetRegId, sourceRegId, source2RegId))

	// store result
	targetOffset := funcVarDict[instr.target]
	instruction = append(instruction, fmt.Sprintf("str x%v, [x29, #%v]", targetRegId, targetOffset))

	ReleaseReg(sourceRegId)
	ReleaseReg(source2RegId)
	ReleaseReg(targetRegId)

	return instruction
}
