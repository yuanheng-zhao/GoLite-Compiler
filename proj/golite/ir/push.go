package ir

import (
	"bytes"
	"fmt"
	"strconv"
)

type Push struct {
	sourceReg []int
}

func NewPush(sourceReg []int) *Push {
	return &Push{sourceReg}
}

func (instr *Push) GetTargets() []int { return []int{} }

func (instr *Push) GetSources() []int {
	sources := []int{}
	for _, src := range instr.sourceReg {
		sources = append(sources, src)
	}
	return sources
}

func (instr *Push) GetImmediate() *int { return nil}

func (instr *Push) GetSourceString() string { return "" }

func (instr *Push) GetLabel() string { return "" }

func (instr *Push) SetLabel(newLabel string){}

func (instr *Push) String() string {
	var out bytes.Buffer
	var strSource string

	for id, src := range instr.sourceReg {
		if id != 0 {
			strSource = strSource + ","
		}
		strSource = strSource + "r" + strconv.Itoa(src)
	}

	out.WriteString(fmt.Sprintf("    push {%s}",strSource))

	return out.String()
}