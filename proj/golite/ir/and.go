package ir

import (
	"bytes"
	"fmt"
)

type And struct {
	target    int       // The target register for the instruction
	sourceReg int       // The first source register of the instruction
	operand   int       // The operand either register or constant
	opty      OperandTy // The type for the operand (REGISTER, IMMEDIATE)
}

func NewAnd(target int, sourceReg int, operand int, opty OperandTy) *And {
	return &And{target, sourceReg, operand, opty}
}

func (instr *And) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *And) GetSources() []int {
	sources := []int{}
	if instr.opty == REGISTER {
		sources = append(sources, instr.sourceReg, instr.operand)
	} else if instr.opty == IMMEDIATE {
		sources = append(sources, instr.sourceReg)
	}
	return sources
}

func (instr *And) GetImmediate() *int {

	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}

func (instr *And) GetSourceString() string {
	return ""
}
func (instr *And) GetLabel() string {
	return ""
}

func (instr *And) SetLabel(newLabel string) {}

func (instr *And) String() string {

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

	out.WriteString(fmt.Sprintf("    and %s,%s,%s", targetReg, sourceReg, operand2))

	return out.String()

}

func (instr *And) TranslateToAssembly(funcVarDict map[int]int) []string {
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

	targetRegId := NextAvailReg()
	instruction = append(instruction, fmt.Sprintf("and x%v, x%v, x%v", targetRegId, sourceRegId, source2RegId))

	// store result
	targetOffset := funcVarDict[instr.target]
	instruction = append(instruction, fmt.Sprintf("str x%v, [x29, #%v]", targetRegId, targetOffset))

	// release
	ReleaseReg(sourceRegId)
	ReleaseReg(source2RegId)
	ReleaseReg(targetRegId)

	return instruction
}
