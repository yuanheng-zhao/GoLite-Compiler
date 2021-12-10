package ir

import (
	"bytes"
	"fmt"
)

type Bl struct {
	label  string
}

func NewBl(label string) *Bl {
	return &Bl{label}
}

func (instr *Bl) GetTargets() []int { return []int{} }

func (instr *Bl) GetSources() []int { return []int{} }

func (instr *Bl) GetImmediate() *int { return nil}

func (instr *Bl) GetSourceString() string { return "" }

func (instr *Bl) GetLabel() string { return instr.label }

func (instr *Bl) SetLabel(newLabel string) {}

func (instr *Bl) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("    bl %s", instr.label))

	return out.String()
}

func (instr *Bl) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	instruction = append(instruction, fmt.Sprintf("\tbl %v",instr.label))

	return instruction
}