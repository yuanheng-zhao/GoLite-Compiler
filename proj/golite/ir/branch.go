package ir

import (
	"bytes"
	"fmt"
)

type Branch struct {
	flagVal  ApsrFlag
	label    string
}

func NewBranch(flagVal ApsrFlag, label string) *Branch {
	return &Branch{flagVal, label}
}

func (instr *Branch) GetTargets() []int { return []int{} }

func (instr *Branch) GetSources() []int { return []int{} }

func (instr *Branch) GetImmediate() *int { return nil }

func (instr *Branch) GetSourceString() string { return "" }

func (instr *Branch) GetLabel() string { return instr.label }

func (instr *Branch) SetLabel(newLabel string) {
	instr.label = newLabel
}

func (instr *Branch) String() string {
	var out bytes.Buffer
	var flag string

	switch instr.flagVal {
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

	operand := fmt.Sprintf("b%v", flag)
	out.WriteString(fmt.Sprintf("    %s %s",operand,instr.label))

	return out.String()
}