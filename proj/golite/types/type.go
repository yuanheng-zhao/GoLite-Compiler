package types

type Type interface {
	GetName() string
}

type IntTy struct {}

func (intTy *IntTy) GetName() string {
	return "int"
}

type BoolTy struct {}

func (boolTy *BoolTy) GetName() string {
	return "bool"
}

type FuncTy struct {}

func (funcTy *FuncTy) GetName() string {
	return "function"
}

type StructTy struct {}

func (structTy *StructTy) GetName() string {
	return "struct"
}

type UnknownTy struct{}

func (unknownTy *UnknownTy) GetName() string {
	return "unknownType"
}

type VoidTy struct {}

func (voidTy *VoidTy) GetName() string {
	return "void"
}


var IntTySig *IntTy
var BoolTySig *BoolTy
var FuncTySig *FuncTy
var StructTySig *StructTy
var UnknownTySig *UnknownTy
var VoidTySig *VoidTy

func init() {
	IntTySig = &IntTy{}
	BoolTySig = &BoolTy{}
	FuncTySig = &FuncTy{}
	StructTySig = &StructTy{}
	UnknownTySig = &UnknownTy{}
	VoidTySig = &VoidTy{}
}