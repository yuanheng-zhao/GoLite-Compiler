package ir

import (
	"bytes"
	"fmt"
)

type New struct {
	target   int
	dataType string
}

func NewNew(target int, dataType string) *New {
	return &New{target, dataType}
}

func (instr *New) GetTargets() []int {
	target := []int{}
	target = append(target, instr.target)
	return target
}

func (instr *New) GetSources() []int { return []int{} }

func (instr *New) GetImmediate() *int { return nil }

func (instr *New) GetSourceString() string {
	return instr.dataType
}

func (instr *New) GetLabel() string { return "" }

func (instr *New) SetLabel(newLabel string) {}

func (instr *New) String() string {
	var out bytes.Buffer
	targetReg := fmt.Sprintf("r%v",instr.target)
	out.WriteString(fmt.Sprintf("    new %s,%s",targetReg,instr.dataType))
	return out.String()
}

func (instr *New) TranslateToAssembly(funcVarDict map[int]int) []string {
	inst := []string{}
	return inst
}