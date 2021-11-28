package ir

import (
	"bytes"
	"fmt"
	"strconv"
)

type Pop struct {
	sourceReg []int
}

func NewPop(sourceReg []int) *Pop {
	return &Pop{sourceReg}
}

func (instr *Pop) GetTargets() []int { return []int{} }

func (instr *Pop) GetSources() []int {
	sources := []int{}
	for _, src := range instr.sourceReg {
		sources = append(sources, src)
	}
	return sources
}

func (instr *Pop) GetImmediate() *int { return nil}

func (instr *Pop) GetSourceString() string { return "" }

func (instr *Pop) GetLabel() string { return "" }

func (instr *Pop) SetLabel(newLabel string){}

func (instr *Pop) String() string {
	var out bytes.Buffer
	var strSource string

	for id, src := range instr.sourceReg {
		if id != 0 {
			strSource = strSource + ","
		}
		strSource = strSource + "r" + strconv.Itoa(src)
	}

	out.WriteString(fmt.Sprintf("pop {%s}",strSource))

	return out.String()
}