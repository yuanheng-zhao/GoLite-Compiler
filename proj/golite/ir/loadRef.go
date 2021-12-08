package ir

import (
	"bytes"
	"fmt"
)

// to access fields of a struct
type LoadRef struct {
	target int
	source int
	field  string
}

func NewLoadRef(target int, source int, field string) *LoadRef {
	return &LoadRef{target, source, field}
}

func (instr *LoadRef) GetTargets() []int {
	targets := []int{}
	targets = append(targets, instr.target)
	return targets
}

func (instr *LoadRef) GetSources() []int {
	sources := []int{}
	sources = append(sources, instr.source)
	return sources
}

func (instr *LoadRef) GetImmediate() *int { return nil }

func (instr *LoadRef) GetSourceString() string {
	return instr.field
}

func (instr *LoadRef) GetLabel() string { return "" }

func (instr *LoadRef) SetLabel(newLabel string) {}

func (instr *LoadRef) String() string {
	var out bytes.Buffer

	targetReg := fmt.Sprintf("r%v",instr.target)
	sourceReg := fmt.Sprintf("r%v",instr.source)
	strField := fmt.Sprintf("@%v",instr.field)

	out.WriteString(fmt.Sprintf("    loadRef %s,%s,%s",targetReg,sourceReg,strField))

	return out.String()
}