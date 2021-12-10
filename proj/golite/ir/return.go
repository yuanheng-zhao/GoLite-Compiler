package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
)

type Ret struct {
	operand int
	opty    OperandTy
}

func NewRet(operand int, opty OperandTy) *Ret {
	return &Ret{operand, opty}
}

func (instr *Ret) GetTargets() []int {
	targets := []int{}
	// Use LR (X30) if register is omitted, but can use other register (ARMv8_InstructionSetOverview.pdf)
	targets = append(targets, 30) // TO-DO : Targets: Return register or (stack)
	return targets
}

func (instr *Ret) GetSources() []int {
	sources := []int{}
	if instr.opty == REGISTER {
		sources = append(sources, instr.operand)
	}
	return sources
}

func (instr *Ret) GetImmediate() *int {

	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}

func (instr *Ret) GetSourceString() string {
	return ""
}
func (instr *Ret) GetLabel() string {
	return ""
}

func (instr *Ret) SetLabel(newLabel string) {}

func (instr *Ret) String() string {

	var out bytes.Buffer

	if instr.opty == VOID {
		out.WriteString(fmt.Sprintf("ret"))
	} else {

		var sourceReg string
		if instr.opty == IMMEDIATE {
			sourceReg = fmt.Sprintf("#%v", instr.operand)
		} else { // REGISTER
			sourceReg = fmt.Sprintf("r%v", instr.operand)
		}

		out.WriteString(fmt.Sprintf("    ret %v", sourceReg))
	}

	return out.String()

}

func (instr *Ret) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	var retRegId int
	var isParam bool

	if instr.opty == REGISTER {
		if retRegId, isParam = paramRegIds[instr.operand]; !isParam {
			operandOffset := funcVarDict[instr.operand]
			retRegId = utility.NextAvailReg()
			instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]",retRegId,operandOffset))
		}
		instruction = append(instruction, fmt.Sprintf("\tmov x0,x%v",retRegId))
	} else if instr.opty == IMMEDIATE {
		instruction = append(instruction, fmt.Sprintf("\tmov x0,%v",instr.operand))
	}

	if !isParam {
		utility.ReleaseReg(retRegId)
	}

	return instruction
}
