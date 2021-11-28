package symboltable

import (
	"proj/golite/ir"
	"proj/golite/types"
)

type SymbolTable struct {
	Parent          *SymbolTable
	htable          map[string]*Entry
	ScopeName       string
	ScopeParamTys   []types.Type
	ScopeParamNames []string
}

func New(parent *SymbolTable, scopeName string) *SymbolTable {
	return &SymbolTable{parent, make(map[string]*Entry), scopeName, []types.Type{}, []string{}}
}

func (st *SymbolTable) Contains(tokLiteral string) Entry {
	if entry, exists := st.htable[tokLiteral]; exists {
		// token literal exists in the symbol table
		return *entry
	}
	return nil
}

func (st *SymbolTable) Insert(tokLiteral string, entry *Entry) {
	st.htable[tokLiteral] = entry
}

type Entry interface {
	SetType(t types.Type)
	SetValue(s string)
	GetEntryType() types.Type
	GetScopeST() *SymbolTable
	GetReturnTy() types.Type // Only implement for funcEntry
	GetRegId() int
}

type VarEntry struct {
	ty    types.Type
	value string
	regId int
}

func NewVarEntry() *VarEntry {
	return &VarEntry{types.UnknownTySig, "", ir.NewRegister()}
}
func (ve *VarEntry) GetEntryType() types.Type {
	return ve.ty
}
func (ve *VarEntry) GetScopeST() *SymbolTable {
	return nil // dummy one, for consistency of Entry interface
}
func (ve *VarEntry) SetType(t types.Type) {
	ve.ty = t
}
func (ve *VarEntry) SetValue(s string) {
	ve.value = s
}
func (ve *VarEntry) GetReturnTy() types.Type {
	// dummy one, never use
	return types.UnknownTySig
}
func (ve *VarEntry) GetRegId() int {
	return ve.regId
}

type FuncEntry struct {
	ty         types.Type
	returnType types.Type // expected return type
	scopeSt    *SymbolTable
}

func NewFuncEntry(retTy types.Type, symTable *SymbolTable) *FuncEntry {
	return &FuncEntry{types.FuncTySig, retTy, symTable}
}
func (fe *FuncEntry) GetEntryType() types.Type {
	return fe.ty // types.FuncTySig
}
func (fe *FuncEntry) GetScopeST() *SymbolTable {
	return fe.scopeSt
}
func (fe *FuncEntry) SetType(t types.Type) {}
func (fe *FuncEntry) SetValue(s string)    {}
func (fe *FuncEntry) GetReturnTy() types.Type {
	return fe.returnType
}
func (fe *FuncEntry) GetRegId() int { return -1 }

type StructEntry struct {
	ty      types.Type
	scopeSt *SymbolTable
}

func NewStructEntry(symTable *SymbolTable) *StructEntry {
	return &StructEntry{types.StructTySig, symTable}
}
func (se *StructEntry) GetEntryType() types.Type {
	return se.ty // types.StructTySig
}
func (se *StructEntry) GetScopeST() *SymbolTable {
	return se.scopeSt
}
func (se *StructEntry) SetType(t types.Type) {}
func (se *StructEntry) SetValue(s string)    {}
func (se *StructEntry) GetReturnTy() types.Type {
	// dummy one. never use
	return types.UnknownTySig
}
func (se *StructEntry) GetRegId() int { return -1 }
