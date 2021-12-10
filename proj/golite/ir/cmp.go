package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
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

func (instr *Cmp) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	var operand1Reg, operand2Reg int
	var isOperand1Param, isOperand2Param bool

	// get operand 1
	if operand1Reg, isOperand1Param = paramRegIds[instr.sourceReg]; !isOperand1Param {
		operand1Reg = utility.NextAvailReg()
		operand1Offset := funcVarDict[instr.sourceReg]
		instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]",operand1Reg,operand1Offset))
	}

	// get operand 2
	if operand2Reg, isOperand2Param = paramRegIds[instr.operand]; !isOperand2Param {
		operand2Reg = utility.NextAvailReg()
		if instr.opty == REGISTER {
			operand2Offset := funcVarDict[instr.operand]
			instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]",operand2Reg,operand2Offset))
		} else {
			instruction = append(instruction, fmt.Sprintf("\tmov x%v,#%v",operand2Reg,instr.operand))
		}
	}

	// compare
	instruction = append(instruction, fmt.Sprintf("\tcmp x%v,x%v",operand1Reg,operand2Reg))

	// release registers
	if !isOperand1Param {
		utility.ReleaseReg(operand1Reg)
	}
	if !isOperand2Param {
		utility.ReleaseReg(operand2Reg)
	}

	return instruction
}