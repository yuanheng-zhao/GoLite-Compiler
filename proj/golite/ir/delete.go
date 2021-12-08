package ir

import (
	"bytes"
	"fmt"
)

type Delete struct {
	sourceReg int
}

func NewDelete(sourceReg int) *Delete {
	return &Delete{sourceReg}
}

func (instr *Delete) GetTargets() []int { return []int{} }

func (instr *Delete) GetSources() []int {
	source := []int{}
	source = append(source, instr.sourceReg)
	return source
}

func (instr *Delete) GetImmediate() *int { return nil }

func (instr *Delete) GetSourceString() string { return ""}

func (instr *Delete) GetLabel() string { return "" }

func (instr *Delete) SetLabel(newLabel string) {}

func (instr *Delete) String() string {
	var out bytes.Buffer
	sourceRegister := fmt.Sprintf("r%v",instr.sourceReg)
	out.WriteString(fmt.Sprintf("    delete %s",sourceRegister))
	return out.String()
}