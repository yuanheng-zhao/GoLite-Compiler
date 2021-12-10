package ir

import (
	"bytes"
	"fmt"
	"proj/golite/utility"
	"strconv"
)

type Pop struct {
	sourceReg []int
	funcName  string // "pop {r4, r5} @Add" in benchmarks/simple/simple1/simple1.iloc
}

func NewPop(sourceReg []int, funcName string) *Pop {
	return &Pop{sourceReg, funcName}
}

func (instr *Pop) GetTargets() []int { return []int{} }

func (instr *Pop) GetSources() []int {
	sources := []int{}
	for _, src := range instr.sourceReg {
		sources = append(sources, src)
	}
	return sources
}

func (instr *Pop) GetImmediate() *int { return nil }

func (instr *Pop) GetSourceString() string { return "" }

func (instr *Pop) GetLabel() string { return "" }

func (instr *Pop) SetLabel(newLabel string) {}

func (instr *Pop) String() string {
	var out bytes.Buffer
	var strSource string

	for id, src := range instr.sourceReg {
		if id != 0 {
			strSource = strSource + ","
		}
		strSource = strSource + "r" + strconv.Itoa(src)
	}

	out.WriteString(fmt.Sprintf("    pop {%s} @%v", strSource, instr.funcName))

	return out.String()
}

func (instr *Pop) TranslateToAssembly(funcVarDict map[int]int, paramRegIds map[int]int) []string {
	instruction := []string{}

	offset := len(instr.sourceReg) * 8
	if offset % 16 != 0 {
		offset += 8
	}

	instruction = append(instruction, fmt.Sprintf("\tadd sp,sp,#%v",offset))

	iteration := 8
	if len(instr.sourceReg) <= 8 {
		iteration = len(instr.sourceReg)
	}

	for i := 1; i < iteration; i++ {
		utility.ReleaseReg(i)
	}
	return instruction
}
