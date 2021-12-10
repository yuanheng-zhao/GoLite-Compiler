package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
)

// to access fields of a struct
type LoadRef struct {
	target   int
	source   int
	field    string
	fieldIdx int
}

func NewLoadRef(target int, source int, field string, fieldIdx int) *LoadRef {
	return &LoadRef{target, source, field, fieldIdx}
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

	targetReg := fmt.Sprintf("r%v", instr.target)
	sourceReg := fmt.Sprintf("r%v", instr.source)
	strField := fmt.Sprintf("@%v", instr.field)

	out.WriteString(fmt.Sprintf("    loadRef %s,%s,%s", targetReg, sourceReg, strField))

	return out.String()
}

func (instr *LoadRef) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	fieldRegId := utility.NextAvailReg()
	structRegId := utility.NextAvailReg()
	structOffset := funcVarDict[instr.source]
	targetOffset := funcVarDict[instr.target]

	instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]", structRegId, structOffset))
	instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x%v,#%v", fieldRegId))
	instruction = append(instruction, fmt.Sprintf("\tstr x%v,[x29,#%v]", fieldRegId, targetOffset))

	return instruction
}
