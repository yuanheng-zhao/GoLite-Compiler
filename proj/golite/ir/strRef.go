package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
)

// to access fields of a struct
type StrRef struct {
	target   int
	source   int
	field    string
	fieldIdx int
}

func NewStrRef(target int, source int, field string, fieldIdx int) *StrRef {
	return &StrRef{target, source, field, fieldIdx}
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

	targetReg := fmt.Sprintf("r%v", instr.target)
	sourceReg := fmt.Sprintf("r%v", instr.source)
	strField := fmt.Sprintf("@%v", instr.field)

	out.WriteString(fmt.Sprintf("    strRef %s,%s,%s", targetReg, sourceReg, strField))

	return out.String()
}

func (instr *StrRef) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	targetRegId := utility.NextAvailReg()
	targetOffSet := funcVarDict[instr.target]
	instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]",targetRegId,targetOffSet))

	sourceRegId := utility.NextAvailReg()
	sourceOffSet := funcVarDict[instr.source]
	instruction = append(instruction, fmt.Sprintf("\tldr x%v,[x29,#%v]",sourceRegId,sourceOffSet))

	fieldOffset := instr.fieldIdx * 8
	instruction = append(instruction, fmt.Sprintf("\tstr x%v,[x%v,#]",sourceRegId,fieldOffset))
	utility.ReleaseReg(targetRegId)
	utility.ReleaseReg(sourceRegId)
	return instruction

	//ILOC:     strRef r15,r14,@x
	//ldr x1,[x29,#-32]
	//ldr x2,[x29,#-24]
	//str x2, [x1,#0]
}
