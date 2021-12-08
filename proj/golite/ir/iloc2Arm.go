package ir

import "fmt"

var regList map[int]bool

func TranslateToAssembly(funcfrags []*FuncFrag) []string {

	armInstructions := []string{}
	regList = make(map[int]bool)
	for i := 0; i < 32; i++ {
		regList[i] = true
	}

	// program title
	armInstructions = append(armInstructions, ".arch armv8-a")
	// global variables

	// code
	armInstructions = append(armInstructions, ".text")

	for _, funcfrag := range funcfrags {
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
		armInstructions = append(armInstructions, prologue(-funcSize)...)

		for _, instruction := range funcfrag.Body {
			armInstructions = append(armInstructions, instruction.TranslateToAssembly(funcVarDict)...)
		}

		armInstructions = append(armInstructions, epilogue(-funcSize)...)
		armInstructions = append(armInstructions, ".size "+funcfrag.Label+", (. - "+funcfrag.Label+")")
	}

	return armInstructions
}

func prologue(size int) []string {
	fmt.Println("size inside prolo: ", size)
	proInst := []string{}
	proInst = append(proInst, "sub sp, sp, 16")
	proInst = append(proInst, "stp x29, x30, [sp]")
	proInst = append(proInst, "mov x29, sp")
	proInst = append(proInst, fmt.Sprintf("sub sp, sp, #%v", size))
	return proInst
}

func epilogue(size int) []string {
	epiInst := []string{}
	epiInst = append(epiInst, fmt.Sprintf("add sp, sp, #%v", size))
	epiInst = append(epiInst, "ldp x29, x30, [sp]")
	epiInst = append(epiInst, "add sp, sp, 16")
	epiInst = append(epiInst, "ret")
	return epiInst
}

func NextAvailReg() int {
	for id, val := range regList {
		if val {
			regList[id] = false
			return id
		}
	}
	return -1
}

func ReleaseReg(regId int) {
	regList[regId] = true
}
