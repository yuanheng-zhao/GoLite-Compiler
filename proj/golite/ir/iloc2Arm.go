package ir

import (
	"fmt"
	"strings"
)

var regList map[int]bool

func TranslateToAssembly(funcfrags []*FuncFrag) []string {

	armInstructions := []string{}
	regList = make(map[int]bool)
	for i := 0; i < 32; i++ {
		regList[i] = true
	}

	// program title
	armInstructions = append(armInstructions, "\t.arch armv8-a")
	// global variables
	if len(funcfrags) > 0 && strings.Contains(funcfrags[0].Label, "Global Variable") {
		for _, instruction := range funcfrags[0].Body {
			if instruction.GetSourceString() != "" {
				varName := instruction.GetSourceString()
				armInstructions = append(armInstructions, fmt.Sprintf("\t.comm %v,8,8", varName))
			}
		}
	}
	// code
	armInstructions = append(armInstructions, "\t.text")

	remainfuncFrags := funcfrags[1:]
	for _, funcfrag := range remainfuncFrags {
		//funcEntry := symTable.Contains(funcfrag.Label)
		//scopeSt := funcEntry.GetScopeST() // symbol table for the current scope
		//countVar := 0                     // count the variables inside this scope
		offset := 0
		funcVarDict := make(map[int]int) // variable name -> offset, e.g. a -> -8, b -> -16
		for _, instruction := range funcfrag.Body {
			if instruction.GetTargets() != nil && len(instruction.GetTargets()) > 0 {
				offset -= 8
				funcVarDict[instruction.GetTargets()[0]] = offset
			}
		}
		//for _, e := range scopeSt.HashTable() {
		//	entry := *e
		//	offset -= -8
		//	funcVarDict[entry.GetRegId()] = offset
		//	countVar += 1
		//	//fmt.Println(varName, entry.GetEntryType().GetName())
		//}

		armInstructions = append(armInstructions, "\t.type "+funcfrag.Label+",%function")
		armInstructions = append(armInstructions, "\t.global "+funcfrag.Label)
		armInstructions = append(armInstructions, "\t.p2align\t\t2")

		funcSize := offset
		if funcSize%16 != 0 {
			funcSize -= 8
		}
		armInstructions = append(armInstructions, fmt.Sprintf("%v:", funcfrag.Label))
		armInstructions = append(armInstructions, prologue(-funcSize)...)

		remainingInstruction := funcfrag.Body[1:]
		for _, instruction := range remainingInstruction {
			armInstructions = append(armInstructions, instruction.TranslateToAssembly(funcVarDict)...)
		}

		armInstructions = append(armInstructions, epilogue(-funcSize)...)
		armInstructions = append(armInstructions, "\t.size "+funcfrag.Label+", (. - "+funcfrag.Label+")")
	}

	return armInstructions
}

func prologue(size int) []string {
	proInst := []string{}
	proInst = append(proInst, "\tsub sp, sp, 16")
	proInst = append(proInst, "\tstp x29, x30, [sp]")
	proInst = append(proInst, "\tmov x29, sp")
	proInst = append(proInst, fmt.Sprintf("\tsub sp, sp, #%v", size))
	return proInst
}

func epilogue(size int) []string {
	epiInst := []string{}
	epiInst = append(epiInst, fmt.Sprintf( "\tadd sp, sp, #%v", size))
	epiInst = append(epiInst, "\tldp x29, x30, [sp]")
	epiInst = append(epiInst, "\tadd sp, sp, 16")
	epiInst = append(epiInst, "\tret")
	return epiInst
}

func NextAvailReg() int {
	//for id, val := range regList {
	//	if val {
	//		regList[id] = false
	//		return id
	//	}
	//}
	for i := 0; i < 32; i++ {
		if regList[i] {
			regList[i] = false
			return i
		}
	}
	return -1
}

func ReleaseReg(regId int) {
	regList[regId] = true
}
