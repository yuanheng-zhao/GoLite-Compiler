package ir

import (
	"bytes"
	"fmt"
)

// to access fields of a struct
type StrRef struct {
	target int
	source int
	field  string
}

func NewStrRef(target int, source int, field string) *StrRef {
	return &StrRef{target, source, field}
}

func (instr *StrRef) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *StrRef) GetSources() []int {
	sources := []int{}
	sources = append(sources, instr.source)
	return sources
}

func (instr *StrRef) GetImmediate() *int { return nil }

func (instr *StrRef) GetSourceString() string {
	return instr.field
}

func (instr *StrRef) GetLabel() string { return "" }

func (instr *StrRef) SetLabel(newLabel string) {}

func (instr *StrRef) String() string {
	var out bytes.Buffer

	targetReg := fmt.Sprintf("r%v",instr.target)
	sourceReg := fmt.Sprintf("r%v",instr.source)
	strField := fmt.Sprintf("@%v",instr.field)

	out.WriteString(fmt.Sprintf("strRef %s,%s,%s",targetReg,sourceReg,strField))

	return out.String()
}