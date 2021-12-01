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
	//ProtoName 		string  // name of the prototype struct
}

func New(parent *SymbolTable, scopeName string) *SymbolTable {
	//return &SymbolTable{parent, make(map[string]*Entry), scopeName, []types.Type{}, []string{}, ""}
	return &SymbolTable{parent, make(map[string]*Entry), scopeName, []types.Type{}, []string{}}
}

// TO-DO : revise Contains
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

// PowerContains returns the entry of a variable searching all symbol tables at the current and above levels
func (st *SymbolTable) PowerContains(varName string) Entry {
	var entry Entry
	currSymtable := st
	for {
		if entry = currSymtable.Contains(varName); entry == nil {
			if currSymtable.Parent == nil {
				return nil
			} else {
				currSymtable = currSymtable.Parent
			}
		} else {
			break
		}
	}
	return entry
}

// CheckGlobalVariable returns true if the given variable name is a global variable name, otherwise false
func (st *SymbolTable) CheckGlobalVariable(varName string) bool {
	currSymtable := st
	for {
		if currSymtable.ScopeName == "global" {
			// in the global symbol table
			if varEntry := currSymtable.Contains(varName); varEntry != nil {
				// global variable varName
				return true
			} else {
				return false
			}
		} else {
			if currSymtable.Parent != nil {
				currSymtable = currSymtable.Parent
			} else {
				return false
			}
		}
	}
}

//func (st *SymbolTable) SetProtoName(name string) {
//	st.ProtoName = name
//}
//
//func (st *SymbolTable) GetProtoName() string {
//	return st.ProtoName
//}

//func (st *SymbolTable) GetCopy(scopeName string, parentSt *SymbolTable) *SymbolTable {
//	instanceSt := New(parentSt, scopeName)
//	for key, entry := range st.htable {
//		instanceSt.htable[key] = entry.
//	}
//	instanceSt.SetProtoName(st.ScopeName)
//	return instanceSt
//}

type Entry interface {
	SetType(t types.Type)
	SetValue(s string)
	GetEntryType() types.Type
	GetScopeST() *SymbolTable
	GetReturnTy() types.Type // Only implement for funcEntry
	GetRegId() int
	//GetCopy(parentSt *SymbolTable) *Entry
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

//func (ve *VarEntry) GetCopy(parentSt *SymbolTable) *VarEntry {
//	return &VarEntry{ve.ty, ve.value, ir.NewRegister()}
//}

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

//func (fe *FuncEntry) GetCopy(parentSt *SymbolTable) *Entry {  // cannot copy a function when initialize a struct
//	return nil
//}

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

//func (se *StructEntry) GetCopy(parentSt *SymbolTable) *StructEntry {
//	return &StructEntry{se.ty, se.scopeSt.GetCopy(se.scopeSt.ProtoName, parentSt)}
//}
