package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
)

type Str struct {
	target    int
	sourceReg int
	operand   int
	globalVar string
	opty      OperandTy
}

func NewStr(target int, sourceReg int, operand int, globalVar string, opty OperandTy) *Str {
	return &Str{target, sourceReg, operand, globalVar, opty}
}

func (instr *Str) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *Str) GetSources() []int {
	sources := []int{}
	if instr.opty == REGISTER {
		sources = append(sources, instr.sourceReg, instr.operand)
	} else if instr.opty == IMMEDIATE || instr.opty == ONEOPERAND {
		sources = append(sources, instr.sourceReg)
	}
	return sources
}

func (instr *Str) GetImmediate() *int {
	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}

func (instr *Str) GetSourceString() string {
	if instr.opty == GLOBALVAR {
		return instr.globalVar
	}
	return ""
}

func (instr *Str) GetLabel() string {
	return ""
}

func (instr *Str) SetLabel(newLabel string) {}

func (instr *Str) String() string {
	var out bytes.Buffer

	targetReg := fmt.Sprintf("r%v", instr.target)

	if instr.opty == REGISTER {
		sourceReg1 := fmt.Sprintf("r%v", instr.sourceReg)
		sourceReg2 := fmt.Sprintf("r%v", instr.operand)
		out.WriteString(fmt.Sprintf("    str %s,%s,%s", targetReg, sourceReg1, sourceReg2))
	} else if instr.opty == IMMEDIATE {
		sourceReg := fmt.Sprintf("r%v", instr.sourceReg)
		operand2 := fmt.Sprintf("#%v", instr.operand)
		out.WriteString(fmt.Sprintf("    str %s,%s,%s", targetReg, sourceReg, operand2))
	} else if instr.opty == ONEOPERAND {
		sourceReg := fmt.Sprintf("r%v", instr.sourceReg)
		out.WriteString(fmt.Sprintf("    str %s,%s", targetReg, sourceReg))
	} else if instr.opty == GLOBALVAR {
		globVarName := fmt.Sprintf("@%v", instr.globalVar)
		out.WriteString(fmt.Sprintf("    str %s,%s", targetReg, globVarName))
	}

	return out.String()
}

func (instr *Str) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	if instr.opty == GLOBALVAR {
		addrRegId := utility.NextAvailReg()
		sourceRegId := utility.NextAvailReg()
		sourceOffset := funcVarDict[instr.target]
		instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]",sourceRegId,sourceOffset))
		instruction = append(instruction, fmt.Sprintf("\tadrp x%v,%v",addrRegId,instr.globalVar))
		instruction = append(instruction, fmt.Sprintf("\tadd x%v,x%v, :lo12:%v",addrRegId,addrRegId,instr.globalVar))
		instruction = append(instruction, fmt.Sprintf("\tstr x%v,[x%v]",sourceRegId,addrRegId))
		utility.ReleaseReg(sourceRegId)
		utility.ReleaseReg(addrRegId)
	}

	//ldr x2,[x29,#-16]
	//adrp x1, p1
	//add x1, x1, :lo12:p1
	//str x2, [x1]

	return instruction
}
