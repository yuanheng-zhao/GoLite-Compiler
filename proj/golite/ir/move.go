package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
)

type Mov struct {
	flag    ApsrFlag
	target  int
	operand int
	opty    OperandTy
	retFlag bool
}

func NewMov(target int, operand int, flag ApsrFlag, opty OperandTy) *Mov {
	return &Mov{flag, target, operand, opty, false}
}

func (instr *Mov) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *Mov) GetSources() []int {
	sources := []int{}
	if instr.opty == REGISTER {
		sources = append(sources, instr.operand)
	}
	return sources
}

func (instr *Mov) GetImmediate() *int {
	if instr.opty == IMMEDIATE {
		return &instr.operand
	}
	return nil
}

func (instr *Mov) GetSourceString() string { return "" }

func (instr *Mov) GetLabel() string { return "" }

func (instr *Mov) SetLabel(newLabel string) {}

func (instr *Mov) String() string {
	var out bytes.Buffer
	var flag string

	switch instr.flag {
	case GT:
		flag = "gt"
		break
	case LT:
		flag = "lt"
		break
	case GE:
		flag = "ge"
		break
	case LE:
		flag = "le"
		break
	case EQ:
		flag = "eq"
		break
	case NE:
		flag = "ne"
		break
	case AL:
		flag = ""
		break
	}
	operator := fmt.Sprintf("mov%v", flag)
	targetReg := fmt.Sprintf("r%v", instr.target)
	var prefix string
	if instr.opty == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	operand2 := fmt.Sprintf("%v%v", prefix, instr.operand)
	out.WriteString(fmt.Sprintf("    %s %s,%s", operator, targetReg, operand2))

	if instr.retFlag {
		out.WriteString(fmt.Sprintf(" @Return"))
	}

	return out.String()
}

func (instr *Mov) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	if instr.retFlag {
		tempRegId := utility.NextAvailReg()
		tempOffset := funcVarDict[instr.target]
		instruction = append(instruction, fmt.Sprintf("\tmov x%v,x0",tempRegId))
		instruction = append(instruction, fmt.Sprintf("\tstr x%v,[x29,#%v]",tempRegId,tempOffset))
		utility.ReleaseReg(tempRegId)
		return instruction
	}
	var sourceRegId, targetRegId int
	var isSourceParam, isTargetParam bool

	if targetRegId, isTargetParam = paramRegIds[instr.target]; !isTargetParam {
		targetRegId = utility.NextAvailReg()
	}


	if instr.opty == REGISTER {
		if sourceRegId, isSourceParam = paramRegIds[instr.operand]; !isSourceParam {
			sourceOffset := funcVarDict[instr.operand]
			sourceRegId = utility.NextAvailReg()
			instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]", sourceRegId, sourceOffset))
		}
	}

	if instr.opty == REGISTER {
		instruction = append(instruction, fmt.Sprintf("\tmov x%v,x%v",targetRegId,sourceRegId))
	} else {
		instruction = append(instruction, fmt.Sprintf("\tmov x%v,#%v",targetRegId,instr.operand))
	}

	if !isTargetParam {
		targetOffset := funcVarDict[instr.target]
		instruction = append(instruction, fmt.Sprintf("\tstr x%v,[x29,#%v]", targetRegId,targetOffset))
	}

	if instr.opty == REGISTER && !isSourceParam {
		utility.ReleaseReg(sourceRegId)
	}
	if !isTargetParam {
		utility.ReleaseReg(targetRegId)
	}

	//targetRegId := NextAvailReg()
	//var sourceRegId int
	//if instr.opty == IMMEDIATE {
	//	instruction = append(instruction, fmt.Sprintf("mov x%v, #%v", targetRegId, instr.operand))
	//} else {
	//	sourceRegId = NextAvailReg()
	//	sourceOffset := funcVarDict[instr.operand]
	//	instruction = append(instruction, fmt.Sprintf("ldr x%v, [x29, #%v]", sourceRegId, sourceOffset))
	//	instruction = append(instruction, fmt.Sprintf("mov x%v, x%v", targetRegId, sourceRegId))
	//}
	//
	//targetOffset := funcVarDict[instr.target]
	//instruction = append(instruction, fmt.Sprintf("str x%v, [x29, #%v]", targetRegId, targetOffset))
	//
	//ReleaseReg(targetRegId)
	//if instr.opty == REGISTER {
	//	ReleaseReg(sourceRegId)
	//}

	return instruction
}

func (instr *Mov) SetRetFlag() {
	instr.retFlag = true
}
