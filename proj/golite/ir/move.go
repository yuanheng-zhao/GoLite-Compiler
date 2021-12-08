package ir

import (
	"bytes"
	"fmt"
)

type Mov struct {
	flag    ApsrFlag
	target  int
	operand int
	opty    OperandTy
}

func NewMov(target int, operand int, flag ApsrFlag, opty OperandTy) *Mov {
	return &Mov{flag, target, operand, opty}
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

func (instr *Mov) String () string {
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
	operator := fmt.Sprintf("mov%v",flag)
	targetReg  := fmt.Sprintf("r%v",instr.target)
	var prefix string
	if instr.opty == IMMEDIATE {
		prefix = "#"
	} else {
		prefix = "r"
	}
	operand2 := fmt.Sprintf("%v%v",prefix,instr.operand)
	out.WriteString(fmt.Sprintf("    %s %s,%s",operator,targetReg,operand2))

	return out.String()
}

func (instr *Mov) TranslateToAssembly(funcVarDict map[int]int) []string {
	inst := []string{}
	return inst
}