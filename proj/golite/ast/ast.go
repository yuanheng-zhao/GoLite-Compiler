package ast

import (
	"bytes"
	"fmt"
	"proj/golite/ir"
	st "proj/golite/symboltable"
	"proj/golite/token"
	"proj/golite/types"
)

// Node The base Node interface that all ast nodes have to access
type Node interface {
	TokenLiteral() string
	String() string
	TypeCheck([]string, *st.SymbolTable) []string
	TranslateToILoc([]ir.Instruction, *st.SymbolTable) []ir.Instruction
}

// Expr All expression nodes implement this interface
type Expr interface {
	Node
	GetType(*st.SymbolTable) types.Type
	GetTargetReg() int // for Factor.TranslateToILoc to retrieve
}

// Stmt All statement nodes implement this interface
type Stmt interface {
	Node
	PerformSABuild([]string, *st.SymbolTable) []string
}

// Func prog, funcs, func nodes need to implement this interface
// two types of TranslateToILoc() functions are defined, but only one is useful
type Func interface {
	Node
	TranslateToILocFunc([]*ir.FuncFrag, *st.SymbolTable) []*ir.FuncFrag
}

/******* Stmt : Statement *******/

type Program struct {
	Token *token.Token
	st    *st.SymbolTable

	Package      *Package
	Import       *Import
	Types        *Types
	Declarations *Declarations
	Functions    *Functions
}

func (p *Program) TokenLiteral() string {
	if p.Token != nil {
		return p.Token.Literal
	}
	panic("Could not determine token literal for program statement.")
}
func (p *Program) String() string {
	out := bytes.Buffer{}
	out.WriteString(p.Package.String())
	out.WriteString(p.Import.String())
	out.WriteString(p.Types.String())
	out.WriteString(p.Declarations.String())
	out.WriteString(p.Functions.String())
	return out.String()
}
func (p *Program) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	p.st = symTable
	errors = p.Package.PerformSABuild(errors, symTable)
	errors = p.Import.PerformSABuild(errors, symTable)
	errors = p.Types.PerformSABuild(errors, symTable)
	errors = p.Declarations.PerformSABuild(errors, symTable)

	// Add **new** global functions for creating an instance of a declared struct
	newScopeSt := st.New(symTable, "new")
	newScopeSt.ScopeParamNames = append(newScopeSt.ScopeParamNames, "structName")
	newScopeSt.ScopeParamTys = append(newScopeSt.ScopeParamTys, types.StructTySig)
	var structEntry st.Entry
	structEntry = st.NewStructEntry(nil)
	newScopeSt.Insert("structName", &structEntry)
	var newEntry st.Entry
	newEntry = st.NewFuncEntry(types.StructTySig, newScopeSt)
	symTable.Insert("new", &newEntry)

	// Add **delete** global functions for deleting/releasing an instance of a declared struct
	deleteScopeSt := st.New(symTable, "delete")
	deleteScopeSt.ScopeParamNames = append(deleteScopeSt.ScopeParamNames, "structName")
	deleteScopeSt.ScopeParamTys = append(deleteScopeSt.ScopeParamTys, types.StructTySig)
	structEntry = st.NewStructEntry(nil)
	deleteScopeSt.Insert("structName", &structEntry)
	var deleteEntry st.Entry
	deleteEntry = st.NewFuncEntry(types.StructTySig, deleteScopeSt)
	symTable.Insert("delete", &deleteEntry)

	errors = p.Functions.PerformSABuild(errors, symTable)
	return errors
}
func (p *Program) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = p.Package.TypeCheck(errors, symTable)
	errors = p.Import.TypeCheck(errors, symTable)
	errors = p.Types.TypeCheck(errors, symTable)
	errors = p.Declarations.TypeCheck(errors, symTable)
	errors = p.Functions.TypeCheck(errors, symTable)
	return errors
}
func (p *Program) TranslateToILocFunc(funcFrag []*ir.FuncFrag, symTable *st.SymbolTable) []*ir.FuncFrag {
	funcFrag = p.Declarations.TranslateToILocFunc(funcFrag, symTable)
	funcFrag = p.Functions.TranslateToILocFunc(funcFrag, symTable)
	return funcFrag
}
func (p *Program) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instructions
}

type Package struct {
	Token *token.Token
	//st    *st.SymbolTable
	Ident IdentLiteral
}

func (pkg *Package) TokenLiteral() string {
	if pkg.Token != nil {
		return pkg.Token.Literal
	}
	panic("Could not determine token literal for package statement")
}
func (pkg *Package) String() string {
	out := bytes.Buffer{}
	out.WriteString("package")
	out.WriteString(" ")
	out.WriteString(pkg.Ident.String())
	out.WriteString(";")
	out.WriteString("\n")
	return out.String()
}
func (pkg *Package) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	return errors
}
func (pkg *Package) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	if pkg.Ident.TokenLiteral() != "main" {
		errors = append(errors, fmt.Sprintf("Only package main is allowed"))
	}
	return errors
}
func (pkg *Package) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}

type Import struct {
	Token *token.Token
	Ident IdentLiteral
}

func (imp *Import) TokenLiteral() string {
	if imp.Token != nil {
		return imp.Token.Literal
	}
	panic("Could not determine token literal for import statement")
}
func (imp *Import) String() string {
	out := bytes.Buffer{}
	out.WriteString("import")
	out.WriteString(" ")
	out.WriteString("\"")
	out.WriteString("fmt") // imp.Ident.String() equivalent
	out.WriteString("\"")
	out.WriteString(";")
	out.WriteString("\n")
	return out.String()
}
func (imp *Import) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	return errors
}
func (imp *Import) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	return errors
}
func (imp *Import) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}

type Types struct {
	Token *token.Token
	//st    *st.SymbolTable
	TypeDeclarations []TypeDeclaration
}

func (tys *Types) TokenLiteral() string {
	if tys.Token != nil {
		return tys.Token.Literal
	}
	panic("Could not determine token literals for the types declarations")
}
func (tys *Types) String() string {
	out := bytes.Buffer{}
	for _, typedec := range tys.TypeDeclarations {
		out.WriteString(typedec.String())
		out.WriteString("\n")
	}
	return out.String()
}
func (tys *Types) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	for _, typedec := range tys.TypeDeclarations {
		errors = typedec.PerformSABuild(errors, symTable)
	}
	return errors
}
func (tys *Types) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	for _, typedec := range tys.TypeDeclarations {
		errors = typedec.TypeCheck(errors, symTable)
	}
	return errors
}
func (tys *Types) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}

type TypeDeclaration struct {
	Token  *token.Token
	st     *st.SymbolTable
	Ident  IdentLiteral
	Fields *Fields
}

func (td *TypeDeclaration) TokenLiteral() string {
	if td.Token != nil {
		return td.Token.Literal
	}
	panic("Could not determine token literals for the type declaration")
}
func (td *TypeDeclaration) String() string {
	out := bytes.Buffer{}
	out.WriteString("type")
	out.WriteString(" ")
	out.WriteString(td.Ident.String())
	out.WriteString(" ")
	out.WriteString("struct")
	out.WriteString("{\n")
	out.WriteString(td.Fields.String())
	out.WriteString("\n}")
	out.WriteString(";")
	out.WriteString("\n")
	return out.String()
}
func (td *TypeDeclaration) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: find duplicate structures
	structName := td.Ident.TokenLiteral()
	scopeSymTable := st.New(symTable, structName)
	td.st = scopeSymTable

	if entry := symTable.Contains(structName); entry != nil {
		errors = append(errors, fmt.Sprintf("[%v]: struct %v already declared", td.Token.LineNum, structName))
	} else {
		var entry st.Entry
		entry = st.NewStructEntry(td.st)
		symTable.Insert(structName, &entry)
		errors = td.Fields.PerformSABuild(errors, td.st)
	}
	return errors
}
func (td *TypeDeclaration) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	//errors2 := td.Fields.TypeCheck(errors, td.st)
	scopeSymTable := symTable.Contains(td.Ident.TokenLiteral()).GetScopeST()
	errors2 := td.Fields.TypeCheck(errors, scopeSymTable)

	errors = append(errors, errors2...)
	return errors
}
func (td *TypeDeclaration) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}

type Fields struct {
	Token *token.Token
	Decls []Decl
}

func (fields *Fields) TokenLiteral() string {
	if fields.Token != nil {
		return fields.Token.Literal
	}
	panic("Could not determine token literals for fields")
}
func (fields *Fields) String() string {
	out := bytes.Buffer{}
	out.WriteString(fields.Decls[0].String())
	out.WriteString(";\n")
	remaining := fields.Decls[1:]
	for _, decl := range remaining {
		out.WriteString(decl.String())
		out.WriteString(";\n")
	}
	return out.String()
}
func (fields *Fields) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, decl := range fields.Decls {
		errors = decl.PerformSABuild(errors, symTable)
	}
	return errors
}
func (fields *Fields) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, decl := range fields.Decls {
		errors = decl.TypeCheck(errors, symTable)
	}
	return errors
}
func (fields *Fields) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}

type Decl struct {
	Token *token.Token
	Ident IdentLiteral
	Ty    *Type
}

func (decl *Decl) TokenLiteral() string {
	if decl.Token != nil {
		return decl.Token.Literal
	}
	panic("Could not determine token literals for decl")
}
func (decl *Decl) String() string {
	out := bytes.Buffer{}
	out.WriteString(decl.Ident.String())
	out.WriteString(" ")
	out.WriteString(decl.Ty.String())
	return out.String()
}
func (decl *Decl) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: find duplicate declarations in functions / structures
	varName := decl.Ident.TokenLiteral()
	if entry := symTable.Contains(varName); entry != nil {
		errors = append(errors, fmt.Sprintf("[%v]: variable %v already declared", decl.Token.LineNum, varName))
	} else {
		var entry st.Entry
		entry = st.NewVarEntry()
		symTable.Insert(varName, &entry)
	}
	return errors
}
func (decl *Decl) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: set type for the variable id

	// Decl = 'id' Type
	// get the type from Type
	varType := decl.Ty.GetType(symTable)
	// update / set type of 'id' in the symbol table
	entry := symTable.Contains(decl.Ident.TokenLiteral())
	// entry must be valid when processing declaration-type statements, because it must have been added to symboltable by PerformSABuild
	entry.SetType(varType)
	// include id and its type as a parameter of the function
	// can be function or struct, but only useful in function
	symTable.ScopeParamTys = append(symTable.ScopeParamTys, varType)
	symTable.ScopeParamNames = append(symTable.ScopeParamNames, decl.Ident.TokenLiteral())
	return errors
}
func (decl *Decl) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}

type Declarations struct {
	Token        *token.Token
	Declarations []Declaration
}

func (ds *Declarations) TokenLiteral() string {
	if ds.Token != nil {
		return ds.Token.Literal
	}
	panic("Could not determine token literals for the declarations")
}
func (ds *Declarations) String() string {
	out := bytes.Buffer{}
	for _, dec := range ds.Declarations {
		out.WriteString(dec.String())
		out.WriteString("\n")
	}
	return out.String()
}
func (ds *Declarations) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, dec := range ds.Declarations {
		errors = dec.PerformSABuild(errors, symTable)
	}
	return errors
}
func (ds *Declarations) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, dec := range ds.Declarations {
		errors = dec.TypeCheck(errors, symTable)
	}
	return errors
}
func (ds *Declarations) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}
func (ds *Declarations) TranslateToILocFunc(funcFrag []*ir.FuncFrag, symTable *st.SymbolTable) []*ir.FuncFrag {
	var frag ir.FuncFrag
	frag.Label = ir.NewLabelWithPre("Global Variable")
	funcLabelInstruct := ir.NewLabelStmt(frag.Label)
	frag.Body = append(frag.Body, funcLabelInstruct)

	for _, dec := range ds.Declarations {
		frag.Body = dec.TranslateToILoc(frag.Body, symTable)
	}

	funcFrag = append(funcFrag, &frag)
	return funcFrag
}

type Declaration struct {
	Token *token.Token
	Ids   *Ids
	Ty    *Type
}

func (d *Declaration) TokenLiteral() string {
	if d.Token != nil {
		return d.Token.Literal
	}
	panic("Could not determine token literals for declaration")
}
func (d *Declaration) String() string {
	out := bytes.Buffer{}
	out.WriteString("var")
	out.WriteString(" ")
	out.WriteString(d.Ids.String())
	out.WriteString(" ")
	out.WriteString(d.Ty.String())
	out.WriteString(";")
	return out.String()
}
func (d *Declaration) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	//if d.Ty.GetType(symTable) == types.StructTySig {
	//	// we want to construct scope table for it/them at this point
	//	structName := d.Ty.TypeLiteral[1:]  // remove the * at position 0; e.g. *foo -> foo
	//	//fmt.Println(structName)
	//
	//	// get the symbol table of the prototype of the struct
	//	if protoStructEntry := symTable.PowerContains(structName); protoStructEntry != nil {
	//		protoScopeSt := protoStructEntry.GetScopeST()
	//		for _, id := range d.Ids.Idents {
	//			duplicateScopeSt := protoScopeSt.GetCopy(id.String(), symTable)
	//			var duplicateEntry st.Entry
	//			duplicateEntry = st.NewStructEntry(duplicateScopeSt)
	//			symTable.Insert(id.String(), &duplicateEntry)
	//		}
	//	} else {
	//		errors = append(errors, fmt.Sprintf("[%v]: struct %v has not been defined", d.Token.LineNum, structName))
	//	}
	//} else {
	//	// objective: none, duplicate definitions are examined in d.Ids.PerformSABuild()
	//	errors = d.Ids.PerformSABuild(errors, symTable)
	//}

	// objective: none, duplicate definitions are examined in d.Ids.PerformSABuild()
	errors = d.Ids.PerformSABuild(errors, symTable)

	return errors
}
func (d *Declaration) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: set type for ids, symbol table only
	decType := d.Ty.GetType(symTable)
	for _, id := range d.Ids.Idents {
		entry := symTable.Contains(id.TokenLiteral())
		entry.SetType(decType)
		if entry.GetEntryType() == types.StructTySig {
			structName := d.Ty.TypeLiteral[1:]
			protoStructEntry := symTable.PowerContains(structName)
			protoScopeSt := protoStructEntry.GetScopeST()
			for _, id := range d.Ids.Idents {
				duplicateScopeSt := protoScopeSt.GetCopy(id.String(), symTable)
				var duplicateEntry st.Entry
				duplicateEntry = st.NewStructEntry(duplicateScopeSt)
				symTable.Insert(id.String(), &duplicateEntry)
			}
		}
	}
	return errors
}
func (d *Declaration) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instrcs = d.Ids.TranslateToILoc(instrcs, symTable)
	return instrcs
}

type Ids struct {
	Token  *token.Token
	Idents []IdentLiteral
}

func (ids *Ids) TokenLiteral() string {
	if ids.Token != nil {
		return ids.Token.Literal
	}
	panic("Could not determine token literals for ids")
}
func (ids *Ids) String() string {
	out := bytes.Buffer{}
	out.WriteString(ids.Idents[0].String())
	remaining := ids.Idents[1:]
	for _, id := range remaining {
		out.WriteString(",")
		out.WriteString(id.String())
	}
	return out.String()
}
func (ids *Ids) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// Objective: find duplicate declarations
	for _, id := range ids.Idents {
		varName := id.TokenLiteral()
		if entry := symTable.Contains(varName); entry != nil {
			errors = append(errors, fmt.Sprintf("[%v]: variable [%v] already declared", id.Token.LineNum, varName))
		} else {
			var entry st.Entry
			entry = st.NewVarEntry()
			symTable.Insert(varName, &entry)
		}
	}
	return errors
}
func (ids *Ids) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: none, accomplished in Declaration
	return errors
}
func (ids *Ids) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	for _, id := range ids.Idents {
		entry := symTable.Contains(id.TokenLiteral())
		regId := entry.GetRegId()
		movIns := ir.NewMov(regId, 0, ir.AL, ir.IMMEDIATE)
		strIns := ir.NewStr(regId, -1, -1, id.TokenLiteral(), ir.GLOBALVAR)
		instructions = append(instructions, movIns)
		instructions = append(instructions, strIns)
	}
	return instructions
}

type Functions struct {
	Token     *token.Token
	Functions []Function
}

func (fs *Functions) TokenLiteral() string {
	if fs.Token != nil {
		return fs.Token.Literal
	}
	panic("Could not determine token literals for the functions")
}
func (fs *Functions) String() string {
	out := bytes.Buffer{}
	for _, fun := range fs.Functions {
		out.WriteString(fun.String())
	}
	return out.String()
}
func (fs *Functions) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, fun := range fs.Functions {
		errors = fun.PerformSABuild(errors, symTable)
	}
	return errors
}
func (fs *Functions) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	for _, fun := range fs.Functions {
		errors = fun.TypeCheck(errors, symTable)
	}
	return errors
}
func (fs *Functions) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instructions
}
func (fs *Functions) TranslateToILocFunc(funcFrag []*ir.FuncFrag, symTable *st.SymbolTable) []*ir.FuncFrag {
	for _, fun := range fs.Functions {
		//funcFrag = fun.TranslateToILocFunc(funcFrag, fun.st)
		scopeSt := symTable.Contains(fun.Ident.TokenLiteral()).GetScopeST()
		funcFrag = fun.TranslateToILocFunc(funcFrag, scopeSt)
	}
	return funcFrag
}

type Function struct {
	Token        *token.Token
	st           *st.SymbolTable
	Ident        IdentLiteral
	Parameters   *Parameters
	ReturnType   *ReturnType
	Declarations *Declarations
	Statements   *Statements
}

func (f *Function) TokenLiteral() string {
	if f.Token != nil {
		return f.Token.Literal
	}
	panic("Could not determine token literals for functions")
}
func (f *Function) String() string {
	out := bytes.Buffer{}
	out.WriteString("func")
	out.WriteString(" ")
	out.WriteString(f.Ident.String())
	out.WriteString(" ")
	out.WriteString(f.Parameters.String())
	out.WriteString(" ")
	out.WriteString(f.ReturnType.String())
	out.WriteString("{")
	out.WriteString("\n")
	out.WriteString(f.Declarations.String())
	out.WriteString(" ")
	out.WriteString(f.Statements.String())
	out.WriteString("}")
	out.WriteString("\n")
	return out.String()
}
func (f *Function) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// Objective: find duplicate function definitions
	funcName := f.Ident.TokenLiteral()
	scopeSymTable := st.New(symTable, funcName)
	f.st = scopeSymTable

	// built-in delete
	// in each function, append the parent of delete to the current scope st
	// so that from the start to the end of the current function scope, we look for
	// struct names inside this current scope (delete still in global st)
	deleteEntry := symTable.PowerContains("delete")
	deleteEntry.GetScopeST().Parent = scopeSymTable

	if entry := symTable.Contains(funcName); entry != nil {
		errors = append(errors, fmt.Sprintf("[%v]: function [%v] has been declared", f.Token.LineNum, funcName))
	} else {
		var entry st.Entry
		entry = st.NewFuncEntry(f.ReturnType.GetType(symTable), f.st)
		symTable.Insert(funcName, &entry)
		errors = f.Parameters.PerformSABuild(errors, f.st)
		errors = f.Declarations.PerformSABuild(errors, f.st)
		errors = f.Statements.PerformSABuild(errors, f.st)
	}
	return errors
}
func (f *Function) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// Objective: add parameters, return type to function symbol table and entry in the outer symbol table
	// parameters are added to both inner symbol table and function signature in the outer symbol table by Decl invoked next line
	currScopeSt := symTable.Contains(f.Ident.TokenLiteral()).GetScopeST()
	f.st = currScopeSt
	errors = f.Parameters.TypeCheck(errors, f.st)
	errors = f.Declarations.TypeCheck(errors, f.st)
	errors = f.Statements.TypeCheck(errors, f.st)
	return errors
}
func (f *Function) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instructions
}
func (f *Function) TranslateToILocFunc(funcFrag []*ir.FuncFrag, symTable *st.SymbolTable) []*ir.FuncFrag {
	var frag ir.FuncFrag
	// function label
	frag.Label = f.Ident.TokenLiteral()
	funcLabelInstruct := ir.NewLabelStmt(frag.Label)
	frag.Body = append(frag.Body, funcLabelInstruct)
	// push values in registers associated with the registers to stack
	//pushReg := []int{}
	//params := symTable.ScopeParamNames
	//for _, param := range params {
	//	entry := symTable.Contains(param)
	//	pushReg = append(pushReg, entry.GetRegId())
	//}
	//if len(pushReg) != 0 {
	//	pushInst := ir.NewPush(pushReg)
	//	frag.Body = append(frag.Body, pushInst)
	//	// reversed for future pop
	//	for i, j := 0, len(pushReg)-1; i < j; i, j = i+1, j-1 {
	//		pushReg[i], pushReg[j] = pushReg[j], pushReg[i]
	//	}
	//}
	// move values from dedicated registers to parameters
	//for i := 1; i <= len(params); i++ {
	//	entry := symTable.Contains(params[i-1])
	//	movInst := ir.NewMov(entry.GetRegId(), i, ir.AL, ir.REGISTER)
	//	frag.Body = append(frag.Body, movInst)
	//}
	// translate function statements
	frag.Body = f.Statements.TranslateToILoc(frag.Body, symTable)
	// pop the previously pushed values in registers associated with parameters
	//if len(pushReg) != 0 {
	//	popInst := ir.NewPop(pushReg)
	//	frag.Body = append(frag.Body, popInst)
	//}

	funcFrag = append(funcFrag, &frag)
	return funcFrag
}

type Parameters struct {
	Token *token.Token
	Decls []Decl
}

func (params *Parameters) TokenLiteral() string {
	if params.Token != nil {
		return params.Token.Literal
	}
	panic("Could not determine token literals for parameters")
}
func (params *Parameters) String() string {
	out := bytes.Buffer{}
	out.WriteString("(")
	var remaining []Decl
	if len(params.Decls) > 0 {
		out.WriteString(params.Decls[0].String())
		remaining = params.Decls[1:]
	}
	for _, decl := range remaining {
		out.WriteString(",")
		out.WriteString(decl.String())
	}
	out.WriteString(")")
	return out.String()
}
func (params *Parameters) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, decl := range params.Decls {
		errors = decl.PerformSABuild(errors, symTable)
	}
	return errors
}
func (params *Parameters) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, decl := range params.Decls {
		errors = decl.TypeCheck(errors, symTable)
	}
	return errors
}

type Statements struct {
	Token      *token.Token
	Statements []Statement
}

func (stmts *Statements) TokenLiteral() string {
	if stmts.Token != nil {
		return stmts.Token.Literal
	}
	panic("Could not determine token literals for the statements")
}
func (stmts *Statements) String() string {
	out := bytes.Buffer{}
	for _, stmt := range stmts.Statements {
		out.WriteString("\t")
		out.WriteString(stmt.String())
	}
	out.WriteString("\n")
	return out.String()
}
func (stmts *Statements) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, stmt := range stmts.Statements {
		errors = stmt.PerformSABuild(errors, symTable)
	}
	return errors
}
func (stmts *Statements) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	for _, stmt := range stmts.Statements {
		errors = stmt.TypeCheck(errors, symTable)
	}
	return errors
}
func (stmts *Statements) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	for _, stmt := range stmts.Statements {
		instructions = stmt.TranslateToILoc(instructions, symTable)
	}
	return instructions
}

type Statement struct {
	Token *token.Token
	Stmt  Stmt
}

func (s *Statement) TokenLiteral() string {
	if s.Token != nil {
		return s.Token.Literal
	}
	panic("Could not determine token literals for statement")
}
func (s *Statement) String() string {
	out := bytes.Buffer{}
	out.WriteString(s.Stmt.String())
	return out.String()
}
func (s *Statement) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	errors = s.Stmt.PerformSABuild(errors, symTable)
	return errors
}
func (s *Statement) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	errors = s.Stmt.TypeCheck(errors, symTable)
	return errors
}
func (s *Statement) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = s.Stmt.TranslateToILoc(instructions, symTable)
	return instructions
}

type Block struct {
	Token      *token.Token
	Statements *Statements
}

func (b *Block) TokenLiteral() string {
	if b.Token != nil {
		return b.Token.Literal
	}
	panic("Could not determine token literals for block")
}
func (b *Block) String() string {
	out := bytes.Buffer{}
	out.WriteString("{")
	out.WriteString("\n")
	out.WriteString(b.Statements.String())
	out.WriteString("}")
	return out.String()
}
func (b *Block) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	errors = b.Statements.PerformSABuild(errors, symTable)
	return errors
}
func (b *Block) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	errors = b.Statements.TypeCheck(errors, symTable)
	return errors
}
func (b *Block) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = b.Statements.TranslateToILoc(instructions, symTable)
	return instructions
}

type Assignment struct {
	Token  *token.Token
	Lvalue *LValue
	Expr   *Expression
}

func (a *Assignment) TokenLiteral() string {
	if a.Token != nil {
		return a.Token.Literal
	}
	panic("Could not determine token literals for assignment")
}
func (a *Assignment) String() string {
	out := bytes.Buffer{}
	out.WriteString(a.Lvalue.String())
	out.WriteString(" ")
	out.WriteString("=")
	out.WriteString(" ")
	out.WriteString(a.Expr.String())
	out.WriteString(";")
	out.WriteString("\n")
	return out.String()
}
func (a *Assignment) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	return errors
}
func (a *Assignment) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: matching of types on both sides of the assignment statement
	errors = a.Lvalue.TypeCheck(errors, symTable)
	errors = a.Expr.TypeCheck(errors, symTable)
	if len(errors) == 0 {
		leftType := a.Lvalue.GetType(symTable)
		rightType := a.Expr.GetType(symTable)
		if leftType != rightType {
			errors = append(errors, fmt.Sprintf("[%v]: type mismatch: Cannot assign %v (Type %v) to %v (Type %v)",
				a.Token.LineNum, a.Expr.String(), rightType.GetName(), a.Lvalue.String(), leftType.GetName()))
			return errors
		}
		if leftType != types.IntTySig && leftType != types.BoolTySig && leftType != types.StructTySig {
			errors = append(errors, fmt.Sprintf("[%v]: %v is not assignable", a.Token.LineNum, a.Lvalue.String()))
			return errors
		}
		if leftType == types.IntTySig && a.Lvalue.Token.Type == token.NUM {
			errors = append(errors, fmt.Sprintf("[%v]: %v is not assignable", a.Token.LineNum, a.Lvalue.String()))
			return errors
		}
		if leftType == types.BoolTySig && (a.Lvalue.Token.Type == token.TRUE || a.Lvalue.Token.Type == token.FALSE) {
			errors = append(errors, fmt.Sprintf("[%v]: %v is not assignable", a.Token.LineNum, a.Lvalue.String()))
			return errors
		}
		// Ignore the following case here:
		// new(foo) = new(foo2)
		//if leftType == types.StructTySig {
		//	// TO-DO :
		//	structName := a.Expr.String()
		//	fmt.Println(structName)
		//	ent := a.Lvalue.getStructEntry(symTable)
		//	fmt.Println(ent.GetScopeST())
		//
		//}

		entry := a.Lvalue.getStructEntry(symTable)
		if entry != nil {
			entry.SetValue(a.Expr.TokenLiteral())
		}
		// to-discuss: the value to be assigned to the variable in the symbol table
	}
	return errors
}

func (a *Assignment) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = a.Lvalue.TranslateToILoc(instructions, symTable)
	instructions = a.Expr.TranslateToILoc(instructions, symTable)
	var instruction ir.Instruction
	exprReg := a.Expr.GetTargetReg()
	if a.Lvalue.Idents == nil || len(a.Lvalue.Idents) == 0 {
		varName := a.Lvalue.Ident.String()
		if symTable.CheckGlobalVariable(varName) {
			//ldrInst := ir.NewLdr(a.Lvalue.Ident.targetReg, -1, -1, a.Lvalue.Ident.Id, ir.GLOBALVAR)
			//instructions = append(instructions, ldrInst)
			// global variable assignment
			instruction = ir.NewStr(exprReg, -1, -1, varName, ir.GLOBALVAR)
		} else {
			// base type assignment
			lvReg := a.Lvalue.GetTargetReg()
			instruction = ir.NewMov(lvReg, exprReg, ir.AL, ir.REGISTER)
		}
	} else {
		// struct assignment
		structAddr := a.Lvalue.GetTargetReg()
		field := a.Lvalue.Idents[len(a.Lvalue.Idents)-1].TokenLiteral()
		instruction = ir.NewStrRef(exprReg, structAddr, field)
	}
	instructions = append(instructions, instruction)
	return instructions
}

type Read struct {
	Token *token.Token
	Ident IdentLiteral
}

func (r *Read) TokenLiteral() string {
	if r.Token != nil {
		return r.Token.Literal
	}
	panic("Could not determine token literals for read")
}
func (r *Read) String() string {
	out := bytes.Buffer{}
	out.WriteString("fmt")
	out.WriteString(".")
	out.WriteString("Scan")
	out.WriteString("(")
	out.WriteString("&")
	out.WriteString(r.Ident.String())
	out.WriteString(")")
	out.WriteString(";")
	out.WriteString("\n")
	return out.String()
}
func (r *Read) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	return errors
}
func (r *Read) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: verify the variable is declared
	varName := r.Ident.TokenLiteral()
	entry := symTable.Contains(varName)
	if entry == nil {
		errors = append(errors, fmt.Sprintf("[%v]: variable %v has not been declared", r.Token.LineNum, varName))
	}
	return errors
}
func (r *Read) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	entry := symTable.Contains(r.Ident.TokenLiteral())
	instruction := ir.NewRead(entry.GetRegId(), r.Ident.String(), symTable.PowerContains(r.Ident).GetRegId())
	instructions = append(instructions, instruction)
	return instructions
}

type Print struct {
	Token       *token.Token
	printMethod string // "Print" | "Println"
	Ident       IdentLiteral
}

func (p *Print) TokenLiteral() string {
	if p.Token != nil {
		return p.Token.Literal
	}
	panic("Could not determine token literals for print")
}
func (p *Print) String() string {
	out := bytes.Buffer{}
	out.WriteString("fmt")
	out.WriteString(".")
	out.WriteString(p.printMethod)
	out.WriteString("(")
	out.WriteString(p.Ident.String())
	out.WriteString(")")
	out.WriteString(";")
	out.WriteString("\n")
	return out.String()
}
func (p *Print) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	return errors
}
func (p *Print) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: verify the variable is declared
	varName := p.Ident.TokenLiteral()
	entry := symTable.Contains(varName)
	if entry == nil {
		errors = append(errors, fmt.Sprintf("[%v]: variable %v has not been declared", p.Token.LineNum, varName))
	}
	return errors
}
func (p *Print) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	var instruction ir.Instruction
	entry := symTable.Contains(p.Ident.TokenLiteral())
	reg := entry.GetRegId()

	if p.printMethod == "Print" {
		instruction = ir.NewPrint(reg)
	} else {
		instruction = ir.NewPrintln(reg)
	}

	instructions = append(instructions, instruction)
	return instructions
}

type Conditional struct {
	Token     *token.Token
	Expr      *Expression
	Block     *Block
	ElseBlock *Block
}

func (cond *Conditional) TokenLiteral() string {
	if cond.Token != nil {
		return cond.Token.Literal
	}
	panic("Could not determine token literals for conditional")
}
func (cond *Conditional) String() string {
	out := bytes.Buffer{}
	out.WriteString("if")
	out.WriteString(" ")
	out.WriteString("(")
	out.WriteString(cond.Expr.String())
	out.WriteString(")")
	out.WriteString(" ")
	out.WriteString(cond.Block.String())
	if cond.ElseBlock != nil {
		out.WriteString("else")
		out.WriteString(cond.ElseBlock.String())
	}
	return out.String()
}
func (cond *Conditional) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	errors = cond.Block.PerformSABuild(errors, symTable)
	errors = cond.ElseBlock.PerformSABuild(errors, symTable)
	return errors
}
func (cond *Conditional) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: boolean expression as the conditional expression surrounded by parenthesis
	condType := cond.Expr.GetType(symTable)
	errors = cond.Expr.TypeCheck(errors, symTable)
	errors = cond.Block.TypeCheck(errors, symTable)
	errors = cond.ElseBlock.TypeCheck(errors, symTable)
	if len(errors) == 0 {
		if condType != types.BoolTySig {
			errors = append(errors, fmt.Sprintf("[%v]: boolean expression is desired, received %v Type %v", cond.Token.LineNum, cond.Expr.String(), condType.GetName()))
		}
	}
	return errors
}
func (cond *Conditional) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	elseLabel := ir.NewLabelWithPre("else")
	doneLabel := ir.NewLabelWithPre("done")
	// conditional expression
	instructions = cond.Expr.TranslateToILoc(instructions, symTable)
	// jump to else if false
	cmpInstruct := ir.NewCmp(cond.Expr.targetReg, 1, ir.IMMEDIATE)
	var brFalseInst ir.Instruction
	if cond.ElseBlock != nil {
		brFalseInst = ir.NewBranch(ir.NE, elseLabel)
	} else {
		brFalseInst = ir.NewBranch(ir.NE, doneLabel)
	}
	instructions = append(instructions, cmpInstruct)
	instructions = append(instructions, brFalseInst)
	// if clause
	instructions = cond.Block.TranslateToILoc(instructions, symTable)
	// else clause
	if cond.ElseBlock != nil {
		brEndInst := ir.NewBranch(ir.AL, doneLabel)
		elsLabelInst := ir.NewLabelStmt(elseLabel)
		instructions = append(instructions, brEndInst)
		instructions = append(instructions, elsLabelInst)
		instructions = cond.ElseBlock.TranslateToILoc(instructions, symTable)
	}
	// end of if statement
	doneLabelInstr := ir.NewLabelStmt(doneLabel)
	instructions = append(instructions, doneLabelInstr)

	return instructions
}

type Loop struct {
	Token *token.Token
	Expr  *Expression
	Block *Block
}

func (lp *Loop) TokenLiteral() string {
	if lp.Token != nil {
		return lp.Token.Literal
	}
	panic("Could not determine token literals for loop")
}
func (lp *Loop) String() string {
	out := bytes.Buffer{}
	out.WriteString("for")
	out.WriteString(" ")
	out.WriteString("(")
	out.WriteString(lp.Expr.String())
	out.WriteString(")")
	out.WriteString(" ")
	out.WriteString(lp.Block.String())
	return out.String()
}
func (lp *Loop) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	errors = lp.Block.PerformSABuild(errors, symTable)
	return errors
}
func (lp *Loop) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: boolean expression as the conditional expression surrounded by parenthesis
	condType := lp.Expr.GetType(symTable)
	errors = lp.Expr.TypeCheck(errors, symTable)
	errors = lp.Block.TypeCheck(errors, symTable)
	if len(errors) == 0 {
		if condType != types.BoolTySig {
			errors = append(errors, fmt.Sprintf("[%v]: boolean expression is desired, received %v Type %v",
				lp.Token.LineNum, lp.Expr.String(), condType.GetName()))
		}
	}
	return errors
}
func (lp *Loop) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	condLabel := ir.NewLabelWithPre("condLabel")
	bodyLabel := ir.NewLabelWithPre("loopBody")
	// b condLabel1
	branchInstruct := ir.NewBranch(ir.AL, condLabel)
	instructions = append(instructions, branchInstruct)
	// loopBody1:
	bodyLabelInstruct := ir.NewLabelStmt(bodyLabel)
	instructions = append(instructions, bodyLabelInstruct)
	// loop body
	instructions = lp.Block.TranslateToILoc(instructions, symTable)
	// condLabel1:
	condLabelInstruct := ir.NewLabelStmt(condLabel)
	instructions = append(instructions, condLabelInstruct)
	// conditional expression
	instructions = lp.Expr.TranslateToILoc(instructions, symTable)
	cmpInstruct := ir.NewCmp(lp.Expr.targetReg, 1, ir.IMMEDIATE)
	lpCondCheckInstruct := ir.NewBranch(ir.EQ, bodyLabel)
	instructions = append(instructions, cmpInstruct)
	instructions = append(instructions, lpCondCheckInstruct)

	return instructions
}

type Return struct {
	Token *token.Token // "RETURN"
	Expr  *Expression  // the return type, nil if not exists
}

func (ret *Return) TokenLiteral() string {
	if ret.Token != nil {
		return ret.Token.Literal
	}
	panic("Could not determine token literals for return")
}
func (ret *Return) String() string {
	out := bytes.Buffer{}
	out.WriteString("return")
	if ret.Expr != nil {
		out.WriteString(" ")
		out.WriteString(ret.Expr.String())
	}
	out.WriteString(";")
	return out.String()
}
func (ret *Return) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	return errors
}
func (ret *Return) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// objective: match return type with signature
	errors = ret.Expr.TypeCheck(errors, symTable)
	// go to outer symbol table and retrieve the entry
	funcEntry := symTable.Parent.Contains(symTable.ScopeName) // must exist
	decRetType := funcEntry.GetReturnTy()                     // must exist
	actRetType := ret.Expr.GetType(symTable)
	if len(errors) == 0 {
		if actRetType != decRetType {
			errors = append(errors, fmt.Sprintf("[%v]: return type expected %v, found %v", ret.Token.LineNum, decRetType, actRetType))
		}
	}
	return errors
}
func (ret *Return) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	var retInst ir.Instruction

	if ret.Expr == nil {
		retInst = ir.NewRet(-1, ir.VOID)
	} else {
		instructions = ret.Expr.TranslateToILoc(instructions, symTable)
		retInst = ir.NewRet(ret.Expr.targetReg, ir.REGISTER)
	}
	instructions = append(instructions, retInst)
	return instructions
}

// Invocation Statement, compared with InvocExpr
type Invocation struct {
	Token *token.Token
	Ident IdentLiteral
	Args  *Arguments
}

func (invoc *Invocation) TokenLiteral() string {
	if invoc.Token != nil {
		return invoc.Token.Literal
	}
	panic("Could not determine token literals for invocation statement")
}
func (invoc *Invocation) String() string {
	out := bytes.Buffer{}
	out.WriteString(invoc.Ident.String())
	out.WriteString(invoc.Args.String())
	out.WriteString(";")
	out.WriteString("\n")
	return out.String()
}
func (invoc *Invocation) PerformSABuild(errors []string, symTable *st.SymbolTable) []string {
	// objective: none
	return errors
}
func (invoc *Invocation) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	if invoc.Ident.String() == "delete" { // built-in delete
		// TO-DO : check the struct name has been declared
		// OK if unmodified
	}

	// check whether function is declared
	funcName := invoc.Ident.TokenLiteral()
	entry := invoc.getFuncEntry(symTable)
	if entry == nil {
		errors = append(errors, fmt.Sprintf("[%v]: function %v has not been defined", invoc.Token.LineNum, funcName))
	} else {
		symTable = entry.GetScopeST()
		errors = invoc.Args.TypeCheck(errors, symTable)
	}
	return errors
}
func (invoc *Invocation) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	if invoc.Ident.TokenLiteral() == "delete" {
		entry := symTable.PowerContains(invoc.Args.Exprs[0].TokenLiteral())
		delInst := ir.NewDelete(entry.GetRegId())
		instructions = append(instructions, delInst)
		return instructions
	}
	// push register values to stack, make space for parameter passing
	arguments := invoc.Args.Exprs
	pushReg := []int{}
	for i := 0; i < len(arguments); i++ {
		pushReg = append(pushReg, i+1)
	}
	if len(pushReg) != 0 {
		pushInstruct := ir.NewPush(pushReg, invoc.Ident.String())
		instructions = append(instructions, pushInstruct)
		// reversed for future pop
		for i, j := 0, len(pushReg)-1; i < j; i, j = i+1, j-1 {
			pushReg[i], pushReg[j] = pushReg[j], pushReg[i]
		}
	}

	// move argument to dedicated registers
	for i := 0; i < len(arguments); i++ {
		movInstruct := ir.NewMov(i, arguments[i].targetReg, ir.MARG, ir.REGISTER)
		instructions = append(instructions, movInstruct)
	}

	// branch to function
	branchInstruct := ir.NewBl(invoc.Ident.TokenLiteral())
	instructions = append(instructions, branchInstruct)

	// after returning from the function
	// pop from stack to restore the previously pushed values
	popInstruct := ir.NewPop(pushReg, invoc.Ident.String())
	instructions = append(instructions, popInstruct)
	return instructions

	//// arguments to be passed in the invocation
	//arguments := invoc.Args.Exprs
	//// formal parameters in the function definition
	//entry := invoc.getFuncEntry(symTable)
	//scopeSymTable := entry.GetScopeST()
	//paramNames := scopeSymTable.ScopeParamNames
	//// move the values of each argument to the registers associated with parameters
	//if arguments != nil && len(arguments) != 0 {
	//	for idx, arg := range arguments {
	//		entry := scopeSymTable.Contains(paramNames[idx])
	//		targetReg := entry.GetRegId()
	//		// targetReg : register with parameter
	//		// sourceReg : register with arguments
	//		passParamInstruct := ir.NewMov(targetReg, arg.targetReg, ir.AL, ir.REGISTER)
	//		instructions = append(instructions, passParamInstruct)
	//	}
	//}
	//branchInstruct := ir.NewBl(invoc.Ident.TokenLiteral())
	//instructions = append(instructions, branchInstruct)
	//return instructions
}
func (invoc *Invocation) getFuncEntry(symTable *st.SymbolTable) st.Entry {
	var entry st.Entry
	varName := invoc.Ident.TokenLiteral()
	for {
		if entry = symTable.Contains(varName); entry == nil {
			if symTable.Parent == nil {
				return nil
			} else {
				symTable = symTable.Parent
			}
		} else {
			break
		}
	}
	return entry
}

func NewProgram(pac *Package, imp *Import, typ *Types, decs *Declarations, funs *Functions) *Program {
	return &Program{nil, nil, pac, imp, typ, decs, funs}
}
func NewPackage(ident IdentLiteral) *Package {
	return &Package{nil, ident}
}
func NewImport(ident IdentLiteral) *Import      { return &Import{nil, ident} }
func NewTypes(typdecs []TypeDeclaration) *Types { return &Types{nil, typdecs} }
func NewTypeDeclaration(ident IdentLiteral, fields *Fields) *TypeDeclaration {
	return &TypeDeclaration{nil, nil, ident, fields}
}
func NewFields(decls []Decl) *Fields                   { return &Fields{nil, decls} }
func NewDecl(ident IdentLiteral, ty *Type) *Decl       { return &Decl{nil, ident, ty} }
func NewDeclarations(decs []Declaration) *Declarations { return &Declarations{nil, decs} }
func NewDeclaration(ids *Ids, Type *Type) *Declaration { return &Declaration{nil, ids, Type} }
func NewIds(idents []IdentLiteral) *Ids                { return &Ids{nil, idents} }
func NewFunctions(funs []Function) *Functions          { return &Functions{nil, funs} }
func NewFunction(ident IdentLiteral, params *Parameters, returnType *ReturnType,
	declarations *Declarations, statements *Statements) *Function {
	return &Function{nil, nil, ident, params, returnType, declarations, statements}
}
func NewParameters(decls []Decl) *Parameters      { return &Parameters{nil, decls} }
func NewReturnType(str string) *ReturnType        { return &ReturnType{nil, NewType(str)} }
func NewStatements(stmts []Statement) *Statements { return &Statements{nil, stmts} }
func NewStatement(stmt Stmt) *Statement           { return &Statement{nil, stmt} }
func NewBlock(statement *Statements) *Block       { return &Block{nil, statement} }
func NewAssignment(lvalue *LValue, expr *Expression) *Assignment {
	return &Assignment{nil, lvalue, expr}
}
func NewRead(ident IdentLiteral) *Read { return &Read{nil, ident} }
func NewPrint(printMethod string, ident IdentLiteral) *Print {
	return &Print{nil, printMethod, ident}
}
func NewConditional(expr *Expression, block *Block, elseBlock *Block) *Conditional {
	return &Conditional{nil, expr, block, elseBlock}
}
func NewLoop(expr *Expression, block *Block) *Loop { return &Loop{nil, expr, block} }
func NewReturn(expr *Expression) *Return           { return &Return{nil, expr} }
func NewInvocation(ident IdentLiteral, args *Arguments) *Invocation {
	return &Invocation{nil, ident, args}
}

/***************** Expr : Expression *******************/

type Type struct {
	Token *token.Token
	// either "int"/"bool"/"*id", where id will actually be the literal for the struct name being defined.
	TypeLiteral string
}

func (t *Type) TokenLiteral() string {
	if t.Token != nil {
		return t.Token.Literal
	}
	panic("Could not determine token literals for type")
}
func (t *Type) String() string {
	return t.TypeLiteral
}
func (t *Type) GetType(symTable *st.SymbolTable) types.Type {
	if t.TypeLiteral == "int" {
		return types.IntTySig
	} else if t.TypeLiteral == "bool" {
		return types.BoolTySig
	} else {
		return types.StructTySig
	}
}
func (t *Type) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	return errors
}
func (t *Type) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}
func (t *Type) GetTargetReg() int {
	return -1 // dummy one, no usage
}

type ReturnType struct {
	Token *token.Token
	Ty    *Type
}

func (rt *ReturnType) TokenLiteral() string {
	if rt.Token != nil {
		return rt.Token.Literal
	}
	panic("Could not determine token literals for returnType")
}
func (rt *ReturnType) String() string {
	out := bytes.Buffer{}
	out.WriteString(rt.Ty.String())
	return out.String()
}
func (rt *ReturnType) GetType(symTable *st.SymbolTable) types.Type {
	if rt.Ty.TypeLiteral == "" {
		return types.VoidTySig
	} else {
		return rt.Ty.GetType(symTable)
	}
}
func (rt *ReturnType) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = rt.Ty.TypeCheck(errors, symTable)
	return errors
}
func (rt *ReturnType) TranslateToILoc(instrcs []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	return instrcs
}
func (rt *ReturnType) GetTargetReg() int {
	return -1 // dummy one, no usage
}

type Arguments struct {
	Token     *token.Token
	Exprs     []Expression // MARKING
	targetReg int
}

func (args *Arguments) TokenLiteral() string {
	if args.Token != nil {
		return args.Token.Literal
	}
	panic("Could not determine token literals for arguments")
}
func (args *Arguments) String() string {
	out := bytes.Buffer{}
	out.WriteString("(")
	if len(args.Exprs) > 0 {
		out.WriteString(args.Exprs[0].String())
		remaining := args.Exprs[1:]
		for _, exp := range remaining {
			out.WriteString(",")
			out.WriteString(exp.String())
		}
	}
	out.WriteString(")")
	return out.String()
}
func (args *Arguments) GetType(symTable *st.SymbolTable) types.Type {
	return types.VoidTySig
}
func (args *Arguments) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	if symTable.ScopeName == "delete" { // built-in delete as a special case

	}
	// used as parameters for calling a function
	expectedTys := symTable.ScopeParamTys
	paramNames := symTable.ScopeParamNames
	if len(expectedTys) != len(args.Exprs) {
		errors = append(errors, fmt.Sprintf("[%v]: Prompted %v parameters, found %v given parameters",
			args.Token.LineNum, len(expectedTys), len(args.Exprs)))
		return errors
	}
	for idx, expr := range args.Exprs {
		errors = expr.TypeCheck(errors, symTable)
		givenParamTy := expr.GetType(symTable)
		if givenParamTy != expectedTys[idx] {
			errors = append(errors, fmt.Sprintf("[%v]: Expected parameter %v type %v; given parameter %v type %v",
				args.Token.LineNum, paramNames[idx], expectedTys[idx].GetName(), expr.String(), givenParamTy.GetName()))
		}
	}
	return errors
}
func (args *Arguments) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	for _, exp := range args.Exprs {
		instructions = exp.TranslateToILoc(instructions, symTable)
	}
	return instructions
}
func (args *Arguments) GetTargetReg() int {
	return args.targetReg
}

type LValue struct {
	Token     *token.Token
	Ident     IdentLiteral
	Idents    []IdentLiteral
	targetReg int
}

func (lv *LValue) TokenLiteral() string {
	if lv.Token != nil {
		return lv.Token.Literal
	}
	panic("Could not determine token literals for lvalue")
}
func (lv *LValue) String() string {
	out := bytes.Buffer{}
	out.WriteString(lv.Ident.String())
	for _, id := range lv.Idents {
		out.WriteString(".")
		out.WriteString(id.String())
	}
	return out.String()
}
func (lv *LValue) GetType(symTable *st.SymbolTable) types.Type {
	entry := lv.getStructEntry(symTable)
	if entry != nil {
		return entry.GetEntryType()
	} else {
		return types.UnknownTySig
	}
}
func (lv *LValue) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	if lv.GetType(symTable) == types.UnknownTySig {
		errors = append(errors, fmt.Sprintf("[%v]: (LValue) inner field has not been defined", lv.Token.LineNum))
	}
	return errors
}
func (lv *LValue) getStructEntry(symTable *st.SymbolTable) st.Entry {
	var entry st.Entry
	varName := lv.Ident.TokenLiteral()
	for {
		if entry = symTable.Contains(varName); entry == nil {
			if symTable.Parent == nil {
				return nil
			} else {
				symTable = symTable.Parent
			}
		} else {
			break
		}
	}
	if lv.Idents == nil {
		return entry
	}
	// here entry is the entry of the first id in Idents
	symTable = entry.GetScopeST() // TO-DO : DEBUG the following
	//remaining := lv.Idents[1:]
	remaining := lv.Idents

	for idx, id := range remaining {
		if entry = symTable.Contains(id.String()); entry == nil {
			return nil
		} else {
			if idx == len(lv.Idents)-1 {
				break
			}
			symTable = entry.GetScopeST()
		}
	}
	return entry
}
func (lv *LValue) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = lv.Ident.TranslateToILoc(instructions, symTable)

	if symTable.CheckGlobalVariable(lv.Ident.String()) { // ldr global variable
		ldrInst := ir.NewLdr(lv.Ident.targetReg, -1, -1, lv.Ident.Id, ir.GLOBALVAR)
		instructions = append(instructions, ldrInst)
	}

	lv.targetReg = lv.Ident.targetReg
	if lv.Idents == nil || len(lv.Idents) == 0 {
		return instructions
	}
	// id.id / id.id.id / ...
	remainingBeforeLast := lv.Idents[:len(lv.Idents)-1]
	source := lv.Ident.targetReg
	for _, ident := range remainingBeforeLast {
		instructions = ident.TranslateToILoc(instructions, symTable)
		target := ir.NewRegister()
		instruction := ir.NewLoadRef(target, source, ident.Id)
		instructions = append(instructions, instruction)
		source = target
		lv.targetReg = target
	}
	return instructions
}
func (lv *LValue) GetTargetReg() int {
	return lv.targetReg
}

type Expression struct {
	Token     *token.Token
	Left      *BoolTerm
	Rights    []BoolTerm
	targetReg int // bind the result of the current Expression to a target register id
}

func (exp *Expression) TokenLiteral() string {
	if exp.Token != nil {
		return exp.Token.Literal
	}
	panic("Could not determine token literals for expression")
}
func (exp *Expression) String() string {
	out := bytes.Buffer{}
	out.WriteString(exp.Left.String())
	for _, boolTerm := range exp.Rights {
		out.WriteString("||")
		out.WriteString(boolTerm.String())
	}
	return out.String()
}
func (exp *Expression) GetType(symTable *st.SymbolTable) types.Type {
	leftType := exp.Left.GetType(symTable)

	for _, rTerm := range exp.Rights {
		rightType := rTerm.GetType(symTable)
		if rightType != types.BoolTySig {
			return types.UnknownTySig
		}
	}
	if exp.Rights != nil && len(exp.Rights) > 0 {
		if leftType != types.BoolTySig {
			return types.UnknownTySig
		}
		return types.BoolTySig
	}
	return leftType
}
func (exp *Expression) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = exp.Left.TypeCheck(errors, symTable)
	leftMostTy := exp.Left.GetType(symTable)
	if len(exp.Rights) != 0 {
		// OR operation, needs bool types on both sides
		if leftMostTy != types.BoolTySig {
			errors = append(errors, fmt.Sprintf("[%v]: (Expression) expected bool type, found %v (%v)",
				exp.Token.LineNum, leftMostTy.GetName(), exp.Left.String()))
			return errors
		}
	}
	// check every expression is in the same type (types.BoolTySig)
	for _, curr := range exp.Rights {
		errors = curr.TypeCheck(errors, symTable)
		currTy := curr.GetType(symTable)
		if currTy != leftMostTy {
			errors = append(errors, fmt.Sprintf("[%v]: (Expression) expected %v type, found %v (%v)",
				exp.Token.LineNum, leftMostTy.GetName(), currTy.GetName(), curr.String()))
			break
		}
	}
	return errors
}
func (exp *Expression) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = exp.Left.TranslateToILoc(instructions, symTable)
	leftSource := exp.Left.targetReg
	if exp.Rights == nil || len(exp.Rights) == 0 {
		exp.targetReg = leftSource
		return instructions
	}

	for _, rTerm := range exp.Rights {
		instructions = rTerm.TranslateToILoc(instructions, symTable)
		target := ir.NewRegister()
		// in this way, OperandTy is always REGISTER
		instruction := ir.NewOr(target, leftSource, rTerm.targetReg, ir.REGISTER)
		instructions = append(instructions, instruction)
		leftSource = target
	}
	exp.targetReg = leftSource
	return instructions
}
func (exp *Expression) GetTargetReg() int {
	return exp.targetReg
}

type BoolTerm struct {
	Token     *token.Token
	Left      *EqualTerm
	Rights    []EqualTerm
	targetReg int
}

func (bt *BoolTerm) TokenLiteral() string {
	if bt.Token != nil {
		return bt.Token.Literal
	}
	panic("Could not determine token literals for boolTerm")
}
func (bt *BoolTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString(bt.Left.String())
	for _, equalTerm := range bt.Rights {
		out.WriteString("&&")
		out.WriteString(equalTerm.String())
	}
	return out.String()
}
func (bt *BoolTerm) GetType(symTable *st.SymbolTable) types.Type {
	leftType := bt.Left.GetType(symTable)

	for _, rTerm := range bt.Rights {
		rightType := rTerm.GetType(symTable)
		if rightType != types.BoolTySig {
			return types.UnknownTySig
		}
	}
	if bt.Rights != nil && len(bt.Rights) > 0 {
		if leftType != types.BoolTySig {
			return types.UnknownTySig
		}
		return types.BoolTySig
	}
	return leftType
}
func (bt *BoolTerm) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = bt.Left.TypeCheck(errors, symTable)
	leftMostTy := bt.Left.GetType(symTable)
	if len(bt.Rights) != 0 {
		// OR operation, needs bool types on both sides
		if leftMostTy != types.BoolTySig {
			errors = append(errors, fmt.Sprintf("[%v]: (BoolTerm) expected bool type, found %v (%v)",
				bt.Token.LineNum, leftMostTy.GetName(), bt.Left.String()))
			return errors
		}
	}
	// check every expression is in the same type (types.BoolTySig)
	for _, curr := range bt.Rights {
		errors = curr.TypeCheck(errors, symTable)
		currTy := curr.GetType(symTable)
		if currTy != leftMostTy {
			errors = append(errors, fmt.Sprintf("[%v]: (BoolTerm) expected %v type, found %v (%v)",
				bt.Token.LineNum, leftMostTy.GetName(), currTy.GetName(), curr.String()))
			break
		}
	}
	return errors
}
func (bt *BoolTerm) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = bt.Left.TranslateToILoc(instructions, symTable)
	if bt.Rights == nil || len(bt.Rights) == 0 {
		bt.targetReg = bt.Left.targetReg
		return instructions
	}

	leftSource := bt.Left.targetReg
	for _, rTerm := range bt.Rights {
		instructions = rTerm.TranslateToILoc(instructions, symTable)
		target := ir.NewRegister()
		// OperandTy is always REGISTER
		instruction := ir.NewAnd(target, leftSource, rTerm.targetReg, ir.REGISTER)
		instructions = append(instructions, instruction)
		leftSource = target
	}
	bt.targetReg = leftSource
	return instructions
}
func (bt *BoolTerm) GetTargetReg() int {
	return bt.targetReg
}

type EqualTerm struct {
	Token         *token.Token
	Left          *RelationTerm
	EqualOperator []string // '=='|'!='
	Rights        []RelationTerm
	targetReg     int
}

func (et *EqualTerm) TokenLiteral() string {
	if et.Token != nil {
		return et.Token.Literal
	}
	panic("Could not determine token literals for equalTerm")
}
func (et *EqualTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString(et.Left.String())
	for i, operator := range et.EqualOperator {
		relationTerm := et.Rights[i]
		out.WriteString(operator)
		out.WriteString(relationTerm.String())
	}
	return out.String()
}
func (et *EqualTerm) GetType(symTable *st.SymbolTable) types.Type {
	leftType := et.Left.GetType(symTable)

	for _, rTerm := range et.Rights {
		rightType := rTerm.GetType(symTable)
		if leftType != rightType {
			return types.UnknownTySig
		}
	}
	if et.Rights != nil && len(et.Rights) > 0 {
		return types.BoolTySig
	} else {
		return leftType
	}
}
func (et *EqualTerm) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = et.Left.TypeCheck(errors, symTable)
	for _, rTerm := range et.Rights {
		errors = rTerm.TypeCheck(errors, symTable)
	}
	return errors
}
func (et *EqualTerm) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = et.Left.TranslateToILoc(instructions, symTable)
	if et.Rights == nil || len(et.Rights) == 0 {
		et.targetReg = et.Left.targetReg
		return instructions
	}

	leftSource := et.Left.targetReg
	for idx, rTerm := range et.Rights {
		instructions = rTerm.TranslateToILoc(instructions, symTable)
		// Put into a new register the "false" value ("false" = 0) before the cmp
		target := ir.NewRegister()
		instruction1 := ir.NewMov(target, 0, ir.AL, ir.IMMEDIATE)
		instruction2 := ir.NewCmp(leftSource, rTerm.targetReg, ir.REGISTER)
		var instruction3 ir.Instruction
		if et.EqualOperator[idx] == "==" {
			instruction3 = ir.NewMov(target, 1, ir.EQ, ir.IMMEDIATE)
		} else { // "!="
			instruction3 = ir.NewMov(target, 1, ir.NE, ir.IMMEDIATE)
		}

		instructions = append(instructions, instruction1, instruction2, instruction3)
		leftSource = target
	}
	et.targetReg = leftSource
	return instructions
}
func (et *EqualTerm) GetTargetReg() int {
	return et.targetReg
}

type RelationTerm struct {
	Token             *token.Token
	Left              *SimpleTerm
	RelationOperators []string // '>'| '<' | '<=' | '>='
	Rights            []SimpleTerm
	targetReg         int
}

func (rt *RelationTerm) TokenLiteral() string {
	if rt.Token != nil {
		return rt.Token.Literal
	}
	panic("Could not determine token literals for relationTerm")
}
func (rt *RelationTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString(rt.Left.String())
	for i, operator := range rt.RelationOperators {
		simpleTerm := rt.Rights[i]
		out.WriteString(operator)
		out.WriteString(simpleTerm.String())
	}
	return out.String()
}
func (rt *RelationTerm) GetType(symTable *st.SymbolTable) types.Type {
	leftType := rt.Left.GetType(symTable)

	for _, rTerm := range rt.Rights {
		rightType := rTerm.GetType(symTable)
		if rightType != types.IntTySig {
			return types.UnknownTySig
		}
	}
	if rt.Rights != nil && len(rt.Rights) > 0 {
		if leftType == types.IntTySig {
			return types.BoolTySig
		}
		return types.UnknownTySig
	}
	return leftType
}
func (rt *RelationTerm) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = rt.Left.TypeCheck(errors, symTable)
	if len(rt.Rights) == 0 {
		return errors
	}
	// with + or - operations, every Term should be int type
	leftMostTy := rt.Left.GetType(symTable)
	if leftMostTy != types.IntTySig {
		errors = append(errors, fmt.Sprintf("[%v]: (RelationTerm) expected int, found %v (%v)",
			rt.Token.LineNum, leftMostTy.GetName(), rt.Left.String()))
		return errors
	}
	for _, rTerm := range rt.Rights {
		errors = rTerm.TypeCheck(errors, symTable)
		currTy := rTerm.GetType(symTable)
		if currTy != types.IntTySig {
			errors = append(errors, fmt.Sprintf("[%v]: (RelationTerm) expected int, found %v (%v)",
				rTerm.Token.LineNum, currTy.GetName(), rTerm.String()))
			return errors
		}
	}
	return errors
}
func (rt *RelationTerm) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = rt.Left.TranslateToILoc(instructions, symTable)
	if rt.Rights == nil || len(rt.Rights) == 0 {
		rt.targetReg = rt.Left.targetReg
		return instructions
	}

	leftSource := rt.Left.targetReg
	for idx, rTerm := range rt.Rights {
		instructions = rTerm.TranslateToILoc(instructions, symTable)
		relationOperator := rt.RelationOperators[idx]
		// Put into a new register the "false" value ("false" = 0) before the cmp
		target := ir.NewRegister()
		instruction1 := ir.NewMov(target, 0, ir.AL, ir.IMMEDIATE)
		instruction2 := ir.NewCmp(leftSource, rTerm.targetReg, ir.REGISTER)
		var instruction3 ir.Instruction
		if relationOperator == ">" {
			instruction3 = ir.NewMov(target, 1, ir.GT, ir.IMMEDIATE)
		} else if relationOperator == "<" {
			instruction3 = ir.NewMov(target, 1, ir.LT, ir.IMMEDIATE)
		} else if relationOperator == "<=" {
			instruction3 = ir.NewMov(target, 1, ir.LE, ir.IMMEDIATE)
		} else { // >=
			instruction3 = ir.NewMov(target, 1, ir.GE, ir.IMMEDIATE)
		}

		instructions = append(instructions, instruction1, instruction2, instruction3)
		leftSource = target
	}
	rt.targetReg = leftSource
	return instructions
}
func (rt *RelationTerm) GetTargetReg() int {
	return rt.targetReg
}

type SimpleTerm struct {
	Token               *token.Token
	Left                *Term
	SimpleTermOperators []string // '+' | '-'
	Rights              []Term
	targetReg           int
}

func (st *SimpleTerm) TokenLiteral() string {
	if st.Token != nil {
		return st.Token.Literal
	}
	panic("Could not determine token literals for simpleTerm")
}
func (st *SimpleTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString(st.Left.String())
	for i, operator := range st.SimpleTermOperators {
		term := st.Rights[i]
		out.WriteString(operator)
		out.WriteString(term.String())
	}
	return out.String()
}
func (st *SimpleTerm) GetType(symTable *st.SymbolTable) types.Type {
	leftType := st.Left.GetType(symTable)

	for _, rTerm := range st.Rights {
		rightType := rTerm.GetType(symTable)
		if rightType != types.IntTySig {
			return types.UnknownTySig
		}
	}
	if st.Rights != nil && len(st.Rights) > 0 {
		if leftType == types.IntTySig {
			return types.IntTySig
		}
		return types.UnknownTySig
	}
	return leftType
}
func (st *SimpleTerm) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = st.Left.TypeCheck(errors, symTable)
	if len(st.Rights) == 0 {
		return errors
	}
	// with + or - operations, every Term should be int type
	leftMostTy := st.Left.GetType(symTable)
	if leftMostTy != types.IntTySig {
		errors = append(errors, fmt.Sprintf("[%v]: (SimpleTerm) expected int, found %v (%v)",
			st.Token.LineNum, leftMostTy.GetName(), st.Left.String()))
		return errors
	}
	for _, rTerm := range st.Rights {
		errors = rTerm.TypeCheck(errors, symTable)
		currTy := rTerm.GetType(symTable)
		if currTy != types.IntTySig {
			errors = append(errors, fmt.Sprintf("[%v]: (SimpleTerm) expected int, found %v (%v)",
				rTerm.Token.LineNum, currTy.GetName(), rTerm.String()))
			return errors
		}
	}
	return errors
}
func (st *SimpleTerm) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = st.Left.TranslateToILoc(instructions, symTable)
	leftSource := st.Left.targetReg

	st.targetReg = leftSource
	if st.Rights == nil || len(st.Rights) == 0 {
		return instructions
	}

	for idx, rTerm := range st.Rights {
		instructions = rTerm.TranslateToILoc(instructions, symTable)
		target := ir.NewRegister()
		var instruction ir.Instruction
		if st.SimpleTermOperators[idx] == "+" {
			instruction = ir.NewAdd(target, leftSource, rTerm.targetReg, ir.REGISTER)
		} else { // "-"
			instruction = ir.NewSub(target, leftSource, rTerm.targetReg, ir.REGISTER)
		}
		st.targetReg = target
		instructions = append(instructions, instruction)
		leftSource = target
	}
	st.targetReg = leftSource
	return instructions
}
func (st *SimpleTerm) GetTargetReg() int {
	return st.targetReg
}

type Term struct {
	Token         *token.Token
	Left          *UnaryTerm
	TermOperators []string // '*' | '/'
	Rights        []UnaryTerm
	targetReg     int
}

func (t *Term) TokenLiteral() string {
	if t.Token != nil {
		return t.Token.Literal
	}
	panic("Could not determine token literals for term")
}
func (t *Term) String() string {
	out := bytes.Buffer{}
	out.WriteString(t.Left.String())
	for i, operator := range t.TermOperators {
		unaryTerm := t.Rights[i]
		out.WriteString(operator)
		out.WriteString(unaryTerm.String())
	}
	return out.String()
}
func (t *Term) GetType(symTable *st.SymbolTable) types.Type {
	leftType := t.Left.GetType(symTable)

	for _, rTerm := range t.Rights {
		rightType := rTerm.GetType(symTable)
		if rightType != types.IntTySig {
			return types.UnknownTySig
		}
	}
	if t.Rights != nil && len(t.Rights) > 0 {
		if leftType == types.IntTySig {
			return types.IntTySig
		}
		return types.UnknownTySig
	}
	return leftType
}
func (t *Term) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = t.Left.TypeCheck(errors, symTable)
	if len(t.Rights) == 0 {
		return errors
	}
	// with * or / operations, every UnaryTerm should be int type
	leftMostTy := t.Left.GetType(symTable)
	if leftMostTy != types.IntTySig {
		errors = append(errors, fmt.Sprintf("[%v]: (Term) expected int, found %v (%v)",
			t.Token.LineNum, leftMostTy.GetName(), t.Left.String()))
		return errors
	}
	for _, rTerm := range t.Rights {
		errors = rTerm.TypeCheck(errors, symTable)
		currTy := rTerm.GetType(symTable)
		if currTy != types.IntTySig {
			errors = append(errors, fmt.Sprintf("[%v]: (Term) expected int, found %v (%v)",
				rTerm.Token.LineNum, currTy.GetName(), rTerm.String()))
			return errors
		}
	}
	return errors
}
func (t *Term) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = t.Left.TranslateToILoc(instructions, symTable)
	leftSource := t.Left.targetReg
	if t.Rights == nil || len(t.Rights) == 0 {
		t.targetReg = leftSource
		return instructions
	}

	for idx, rTerm := range t.Rights {
		instructions = rTerm.TranslateToILoc(instructions, symTable)
		target := ir.NewRegister()
		var instruction ir.Instruction
		if t.TermOperators[idx] == "*" {
			instruction = ir.NewMul(target, leftSource, rTerm.targetReg)
		} else { // "/"
			instruction = ir.NewDiv(target, leftSource, rTerm.targetReg)
		}
		instructions = append(instructions, instruction)
		leftSource = target
	}
	t.targetReg = leftSource
	return instructions
}
func (t *Term) GetTargetReg() int {
	return t.targetReg
}

type UnaryTerm struct {
	Token         *token.Token
	UnaryOperator string // '!' | '-' | '' <- default
	SelectorTerm  *SelectorTerm
	targetReg     int
}

func (ut *UnaryTerm) TokenLiteral() string {
	if ut.Token != nil {
		return ut.Token.Literal
	}
	panic("Could not determine token literals for unaryTerm")
}
func (ut *UnaryTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString(ut.UnaryOperator)
	out.WriteString(ut.SelectorTerm.String())
	return out.String()
}
func (ut *UnaryTerm) GetType(symTable *st.SymbolTable) types.Type {
	if ut.UnaryOperator == "!" {
		if ut.SelectorTerm.GetType(symTable) == types.BoolTySig {
			return types.BoolTySig
		} else {
			return types.UnknownTySig
		}
	} else if ut.UnaryOperator == "-" {
		if ut.SelectorTerm.GetType(symTable) == types.IntTySig {
			return types.IntTySig
		} else {
			return types.UnknownTySig
		}
	} else {
		return ut.SelectorTerm.GetType(symTable)
	}
}
func (ut *UnaryTerm) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = ut.SelectorTerm.TypeCheck(errors, symTable)
	if ut.GetType(symTable) == types.UnknownTySig {
		errors = append(errors, fmt.Sprintf("[%v]: (UnaryTerm) Unkown type", ut.Token.LineNum))
	}
	return errors
}
func (ut *UnaryTerm) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = ut.SelectorTerm.TranslateToILoc(instructions, symTable)
	if ut.UnaryOperator == "" {
		ut.targetReg = ut.SelectorTerm.targetReg
	} else if ut.UnaryOperator == "!" {
		target := ir.NewRegister()
		instruction := ir.NewNot(target, ut.SelectorTerm.targetReg, ir.REGISTER)
		instructions = append(instructions, instruction)
		ut.targetReg = target
	} else { // "-"
		target1 := ir.NewRegister()
		instruction1 := ir.NewMov(target1, 0, ir.AL, ir.IMMEDIATE) // mov r_x,#0
		target2 := ir.NewRegister()
		instruction2 := ir.NewSub(target2, target1, ut.SelectorTerm.targetReg, ir.REGISTER)
		instructions = append(instructions, instruction1, instruction2)
		ut.targetReg = target2
	}
	return instructions
}
func (ut *UnaryTerm) GetTargetReg() int {
	return ut.targetReg
}

type SelectorTerm struct {
	Token     *token.Token
	Fact      *Factor
	Idents    []IdentLiteral
	targetReg int
}

func (selt *SelectorTerm) TokenLiteral() string {
	if selt.Token != nil {
		return selt.Token.Literal
	}
	panic("Could not determine token literals for selectorTerm")
}
func (selt *SelectorTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString(selt.Fact.String())
	for _, id := range selt.Idents {
		out.WriteString(".")
		out.WriteString(id.String())
	}
	return out.String()
}
func (selt *SelectorTerm) GetType(symTable *st.SymbolTable) types.Type {
	facType := selt.Fact.GetType(symTable)
	if len(selt.Idents) == 0 {
		return facType
	} else if facType == types.StructTySig {
		var entry st.Entry
		varName := selt.Fact.String()
		for {
			if entry = symTable.Contains(varName); entry == nil {
				if symTable.Parent == nil {
					return types.UnknownTySig
				} else {
					symTable = symTable.Parent
				}
			} else {
				break
			}
		}
		if selt.Idents == nil {
			return entry.GetEntryType()
		}
		// here entry is the entry of the first id in Idents
		symTable = entry.GetScopeST()
		//remaining := selt.Idents[1:]  // comment here for passing sa test of the Twiddleedee benchmark
		remaining := selt.Idents
		for idx, id := range remaining {
			if entry = symTable.Contains(id.String()); entry == nil {
				return types.UnknownTySig
			} else {
				if idx == len(selt.Idents)-1 {
					break
				}
				symTable = entry.GetScopeST()
			}
		}
		return entry.GetEntryType()
	}
	return types.UnknownTySig
}
func (selt *SelectorTerm) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = selt.Fact.TypeCheck(errors, symTable)
	if selt.GetType(symTable) == types.UnknownTySig {
		errors = append(errors, fmt.Sprintf("[%v]: (SelectorTerm) Unknown type", selt.Fact.Token.LineNum))
		return errors
	}
	return errors
}
func (selt *SelectorTerm) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = selt.Fact.TranslateToILoc(instructions, symTable)
	if selt.Idents == nil || len(selt.Idents) == 0 {
		selt.targetReg = selt.Fact.targetReg
		return instructions
	}
	// Factor.id / Factor.id.id / ...
	source := selt.Fact.targetReg
	symTable = symTable.PowerContains(selt.Fact.String()).GetScopeST()
	for _, ident := range selt.Idents {
		instructions = ident.TranslateToILoc(instructions, symTable)
		target := ir.NewRegister()
		instruction := ir.NewLoadRef(target, source, ident.Id)
		instructions = append(instructions, instruction)
		source = target
		selt.targetReg = target
	}
	return instructions
}
func (selt *SelectorTerm) GetTargetReg() int {
	return selt.targetReg
}

type Factor struct {
	Token     *token.Token
	Expr      Expr
	targetReg int
}

func (f *Factor) TokenLiteral() string {
	if f.Token != nil {
		return f.Token.Literal
	}
	panic("Could not determine token literal for factor")
}
func (f *Factor) String() string {
	out := bytes.Buffer{}
	out.WriteString(f.Expr.String())
	return out.String()
}
func (f *Factor) GetType(symTable *st.SymbolTable) types.Type {
	return f.Expr.GetType(symTable)
}
func (f *Factor) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = f.Expr.TypeCheck(errors, symTable)
	return errors
}
func (f *Factor) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = f.Expr.TranslateToILoc(instructions, symTable)
	f.targetReg = f.Expr.GetTargetReg()
	return instructions
}
func (f *Factor) GetTargetReg() int {
	return f.targetReg
}

func NewType(typeLit string) *Type          { return &Type{nil, typeLit} }
func NewArgs(exprs []Expression) *Arguments { return &Arguments{nil, exprs, -1} }
func NewLvalue(ident IdentLiteral, idents []IdentLiteral) *LValue {
	return &LValue{nil, ident, idents, -1}
}
func NewExpression(l *BoolTerm, rs []BoolTerm) *Expression {
	return &Expression{nil, l, rs, -1}
}
func NewBoolTerm(l *EqualTerm, rs []EqualTerm) *BoolTerm { return &BoolTerm{nil, l, rs, -1} }
func NewEqualTerm(l *RelationTerm, operators []string, rs []RelationTerm) *EqualTerm {
	return &EqualTerm{nil, l, operators, rs, -1}
}
func NewRelationTerm(l *SimpleTerm, operators []string, rs []SimpleTerm) *RelationTerm {
	return &RelationTerm{nil, l, operators, rs, -1}
}
func NewSimpleTerm(l *Term, operators []string, rs []Term) *SimpleTerm {
	return &SimpleTerm{nil, l, operators, rs, -1}
}
func NewTerm(l *UnaryTerm, operators []string, rs []UnaryTerm) *Term {
	return &Term{nil, l, operators, rs, -1}
}
func NewUnaryTerm(operator string, selectorTerm *SelectorTerm) *UnaryTerm {
	return &UnaryTerm{nil, operator, selectorTerm, -1}
}
func NewSelectorTerm(factor *Factor, idents []IdentLiteral) *SelectorTerm {
	return &SelectorTerm{nil, factor, idents, -1}
}
func NewFactor(expr *Expr) *Factor { return &Factor{nil, *expr, -1} }

/********************************* Expr inside Factor ***************************************/

// InvocExpr invocation in Factor ('id' [Arguments])
type InvocExpr struct {
	Token     *token.Token
	Ident     IdentLiteral
	InnerArgs *Arguments
	targetReg int
}

func (ie *InvocExpr) TokenLiteral() string {
	if ie.Token != nil {
		return ie.Token.Literal
	}
	panic("Could not determine token literal for invocation expression inside Factor")
}
func (ie *InvocExpr) String() string {
	out := bytes.Buffer{}
	out.WriteString(ie.Ident.String())
	out.WriteString(ie.InnerArgs.String())
	return out.String()
}
func (ie *InvocExpr) GetType(symTable *st.SymbolTable) types.Type {
	if funcEntry := symTable.PowerContains(ie.Ident.Id); funcEntry != nil {
		return funcEntry.GetReturnTy()
	}
	return types.UnknownTySig
}
func (ie *InvocExpr) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	// refer from Invocation.TypeCheck
	funcName := ie.Ident.TokenLiteral()
	entry := symTable.PowerContains(funcName)
	if entry == nil {
		errors = append(errors, fmt.Sprintf("[%v]: function %v has not been defined", ie.Token.LineNum, funcName))
	} else {
		scopeSymTable := entry.GetScopeST()
		if symTable.ScopeName == "main" { // special case : when calling a function in main,
			// attach the symbol table of main as the parent of the symbol table of the function
			scopeSymTable.Parent = symTable
		}
		errors = ie.InnerArgs.TypeCheck(errors, scopeSymTable)
	}
	return errors
}
func (ie *InvocExpr) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	ie.targetReg = ir.NewRegister()

	if ie.Ident.String() == "new" {
		newInst := ir.NewNew(ie.GetTargetReg(), ie.InnerArgs.Exprs[0].TokenLiteral())
		instructions = append(instructions, newInst)
		return instructions
	}
	// push register values to stack, make space for parameter passing
	arguments := ie.InnerArgs.Exprs
	pushReg := []int{}
	for i := 0; i < len(arguments); i++ {
		// retrieve reg id of the arguments from the symbol table
		pushReg = append(pushReg, symTable.Contains(arguments[i].String()).GetRegId())
		//pushReg = append(pushReg, i+1)
	}
	if len(pushReg) != 0 {
		pushInstruct := ir.NewPush(pushReg, ie.Ident.String())
		instructions = append(instructions, pushInstruct)
		// reversed for future pop
		//for i, j := 0, len(pushReg)-1; i < j; i, j = i+1, j-1 {
		//	pushReg[i], pushReg[j] = pushReg[j], pushReg[i]
		//}
	}

	// move argument to dedicated registers
	//for i := 0; i < len(arguments); i++ {
	//	movInstruct := ir.NewMov(i, arguments[i].targetReg, ir.MARG, ir.REGISTER)
	//	instructions = append(instructions, movInstruct)
	//}

	// branch to function
	branchInstruct := ir.NewBl(ie.Ident.TokenLiteral())
	instructions = append(instructions, branchInstruct)

	// move the return value from the function to the target
	// by default the return value stored in r0
	movRetInstruct := ir.NewMov(ie.targetReg, 0, ir.AL, ir.REGISTER)
	movRetInstruct.SetRetFlag()
	instructions = append(instructions, movRetInstruct)

	// after returning from the function
	// pop from stack to restore the previously pushed values
	popInstruct := ir.NewPop(pushReg, ie.Ident.String())
	instructions = append(instructions, popInstruct)
	return instructions
}
func (ie *InvocExpr) GetTargetReg() int {
	return ie.targetReg
}

// PriorityExpression : '(' Expression ')' (inside Factor)
type PriorityExpression struct {
	Token           *token.Token
	InnerExpression *Expression
	targetReg       int
}

func (pe *PriorityExpression) TokenLiteral() string {
	if pe.Token != nil {
		return pe.Token.Literal
	}
	panic("Could not determine token literal for expression inside Factor")
}
func (pe *PriorityExpression) String() string {
	out := bytes.Buffer{}
	out.WriteString("(")
	out.WriteString(pe.InnerExpression.String())
	out.WriteString(")")
	return out.String()
}
func (pe *PriorityExpression) GetType(symTable *st.SymbolTable) types.Type {
	return pe.InnerExpression.GetType(symTable)
}
func (pe *PriorityExpression) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	errors = pe.InnerExpression.TypeCheck(errors, symTable)
	return errors
}
func (pe *PriorityExpression) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	instructions = pe.InnerExpression.TranslateToILoc(instructions, symTable)
	pe.targetReg = pe.InnerExpression.targetReg
	return instructions
}
func (pe *PriorityExpression) GetTargetReg() int {
	return pe.targetReg
}

// NilNode : nil (keyword "nil")
type NilNode struct {
	// TO-DO  : how to assign a targetReg to NilNode
	// Updated: currently we represent a NIL literal as just equal to 0 (discussed on Ed)
	Token     *token.Token
	targetReg int
}

func (n *NilNode) TokenLiteral() string                        { return n.Token.Literal }
func (n *NilNode) String() string                              { return n.Token.Literal }
func (n *NilNode) GetType(symTable *st.SymbolTable) types.Type { return types.VoidTySig }
func (n *NilNode) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	return errors
}
func (n *NilNode) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	n.targetReg = ir.NewRegister()
	return instructions
}
func (n *NilNode) GetTargetReg() int {
	return n.targetReg
}

// BoolLiteral : True/False
type BoolLiteral struct {
	Token     *token.Token
	Value     bool
	targetReg int
}

func (bl *BoolLiteral) TokenLiteral() string                        { return bl.Token.Literal }
func (bl *BoolLiteral) String() string                              { return bl.Token.Literal }
func (bl *BoolLiteral) GetType(symTable *st.SymbolTable) types.Type { return types.BoolTySig }
func (bl *BoolLiteral) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	return errors
}
func (bl *BoolLiteral) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	bl.targetReg = ir.NewRegister()
	operandValue := 0
	if bl.Value {
		operandValue = 1
	} // set operand #1 if the boolean value is true
	instruction := ir.NewMov(bl.targetReg, operandValue, ir.AL, ir.IMMEDIATE)
	instructions = append(instructions, instruction)
	return instructions
}
func (bl *BoolLiteral) GetTargetReg() int {
	return bl.targetReg
}

// IntLiteral : number (integer)
type IntLiteral struct {
	Token     *token.Token
	Value     int64
	targetReg int
}

func (il *IntLiteral) TokenLiteral() string                        { return il.Token.Literal }
func (il *IntLiteral) String() string                              { return il.Token.Literal }
func (il *IntLiteral) GetType(symTable *st.SymbolTable) types.Type { return types.IntTySig }
func (il *IntLiteral) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	return errors
}
func (il *IntLiteral) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {
	il.targetReg = ir.NewRegister()
	intValue := int(il.Value)
	instruction := ir.NewMov(il.targetReg, intValue, ir.AL, ir.IMMEDIATE)
	instructions = append(instructions, instruction)
	return instructions
}
func (il *IntLiteral) GetTargetReg() int {
	return il.targetReg
}

// IdentLiteral : identifier
type IdentLiteral struct {
	Token     *token.Token
	Id        string
	targetReg int
}

func (idl *IdentLiteral) TokenLiteral() string { return idl.Token.Literal }
func (idl *IdentLiteral) String() string       { return idl.Token.Literal }
func (idl *IdentLiteral) GetType(symTable *st.SymbolTable) types.Type {
	var entry st.Entry
	for {
		if entry = symTable.Contains(idl.TokenLiteral()); entry == nil {
			if symTable.Parent == nil {
				return types.UnknownTySig
			} else {
				symTable = symTable.Parent
			}
		} else {
			break
		}
	}
	return entry.GetEntryType()
}
func (idl *IdentLiteral) TypeCheck(errors []string, symTable *st.SymbolTable) []string {
	if idl.GetType(symTable) == types.UnknownTySig {
		errors = append(errors, fmt.Sprintf("[%v]: %v has not been defined.", idl.Token.LineNum, idl.Id))
	}
	return errors
}
func (idl *IdentLiteral) TranslateToILoc(instructions []ir.Instruction, symTable *st.SymbolTable) []ir.Instruction {

	//if symTable.CheckGlobalVariable(idl.Id) { // if the ident is a global variable
	//	idl.targetReg = ir.NewRegister()
	//	instruction := ir.NewLdr(idl.targetReg, -1, -1, idl.Id, ir.GLOBALVAR)
	//	instructions = append(instructions, instruction)
	//} else {
	//	// use Powercontain or contains here?
	//	sourceReg := symTable.PowerContains(idl.Id).GetRegId()
	//	idl.targetReg = sourceReg
	//}

	// use Powercontain or contains here?
	sourceReg := symTable.PowerContains(idl.Id).GetRegId()
	idl.targetReg = sourceReg

	return instructions
}
func (idl *IdentLiteral) GetTargetReg() int {
	return idl.targetReg
}

func NewIdentLiteral(tok *token.Token, id string) IdentLiteral { return IdentLiteral{tok, id, -1} }
