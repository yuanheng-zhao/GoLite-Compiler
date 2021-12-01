package ir

import (
	"bytes"
	"fmt"
)

type Label struct {
	label string
}

func NewLabelStmt(label string) *Label {
	return &Label{label}
}

func (instr *Label) GetTargets() []int { return []int{} }

func (instr *Label) GetSources() []int { return []int{} }

func (instr *Label) GetImmediate() *int { return nil }

func (instr *Label) GetSourceString() string { return "" }

func (instr *Label) GetLabel() string { return instr.label }

func (instr *Label) SetLabel(newLabel string) {
	instr.label = newLabel
}

func (instr *Label) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("%s: ",instr.label))

	return out.String()
}