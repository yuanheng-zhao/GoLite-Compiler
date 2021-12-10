package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
)

type Not struct {
	target  int       // The target register for the instruction
	operand int       // The operand either register or constant
	opty    OperandTy // The type for the operand (REGISTER, IMMEDIATE)
}

func NewNot(target int, operand int, opty OperandTy) *Not {
	return &Not{target, operand, opty}
}

func (instr *Not) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *Not) GetSources() []int {

	sources := []int{}
	if instr.opty != IMMEDIATE {
		sources = append(sources, instr.operand)
		return sources
	}
	return sources
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

func (instr *Not) SetLabel(newLabel string) {}

func (instr *Not) String() string {

	var out bytes.Buffer

	targetReg := fmt.Sprintf("r%v", instr.target)

	var prefix string
	var operand2 string

	if instr.opty == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	operand2 = fmt.Sprintf("%v%v", prefix, instr.operand)

	out.WriteString(fmt.Sprintf("    not %s,%s", targetReg, operand2))

	return out.String()

}

func (instr *Not) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	// load operand
	sourceRegId := utility.NextAvailReg()
	if instr.opty == REGISTER {
		source2Offset := funcVarDict[instr.operand]
		instruction = append(instruction, fmt.Sprintf("ldr x%v, [x29, #%v]", sourceRegId, source2Offset))
	} else {
		instruction = append(instruction, fmt.Sprintf("mov x%v, #%v", sourceRegId, instr.operand))
	}

	targetRegId := utility.NextAvailReg()
	instruction = append(instruction, fmt.Sprintf("neg x%v, x%v", targetRegId, sourceRegId))

	// store result
	targetOffset := funcVarDict[instr.target]
	instruction = append(instruction, fmt.Sprintf("str x%v, [x29, #%v]", targetRegId, targetOffset))

	utility.ReleaseReg(sourceRegId)
	utility.ReleaseReg(targetRegId)

	return instruction
}
