package ir

import (
	"bytes"
	"fmt"
)

type Read struct {
	targetReg int
}

func NewRead(targetReg int) *Read {
	return &Read{targetReg}
}

func (instr *Read) GetTargets() []int {
	target := []int{}
	target = append(target, instr.targetReg)
	return target
}

func (instr *Read) GetSources() []int { return []int{} }

func (instr *Read) GetImmediate() *int { return nil }

func (instr *Read) GetSourceString() string { return ""}

func (instr *Read) GetLabel() string { return "" }

func (instr *Read) SetLabel(newLabel string) {}

func (instr *Read) String() string {
	var out bytes.Buffer

	targetRegister := fmt.Sprintf("r%v",instr.targetReg)
	out.WriteString(fmt.Sprintf("    read %s",targetRegister))
	return out.String()
}

func ReadArmFormat() []string {
	readInsts := []string{}
	readInsts = append(readInsts, ".READ")
	readInsts = append(readInsts, "\t.asciz\t\"%ld\"")
	readInsts = append(readInsts, "\t.size\t.READ, 4")
	return readInsts
}