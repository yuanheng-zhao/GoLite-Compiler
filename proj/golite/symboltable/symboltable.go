package symboltable

import "proj/golite/types"

type SymbolTable struct {
	Parent *SymbolTable
	htable map[string] *Entry
	ScopeName string
	ScopeParamTy []types.Type
}

func New (parent *SymbolTable, scopeName string, scopeParamTy []types.Type) *SymbolTable {
	return &SymbolTable{parent, make(map[string] *Entry), scopeName, scopeParamTy}
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
	GetEntryType() types.Type
	GetScopeST() *SymbolTable
}

type VarEntry struct {
	Entry
	T types.Type
}

func NewVarEntry() *VarEntry {
	return &VarEntry{}
}

type FuncEntry struct {
	Entry
	T types.Type
}

func NewFuncEntry() *FuncEntry {
	return &FuncEntry{}
}

type StructEntry struct {
	Entry
	T types.Type
}

func NewStructEntry() *StructEntry {
	return &StructEntry{}
}

