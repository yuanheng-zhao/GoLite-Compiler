package ast

import (
	"bytes"
	"proj/golite/token"
)

// Node The base Node interface that all ast nodes have to access
type Node interface {
	TokenLiteral() string
	String() string
	//TypeCheck() []string  // TO-DO
}

// Expr All expression nodes implement this interface
type Expr interface {
	Node
	GetType() // TO-DO
}

// Stmt All statement nodes implement this interface
type Stmt interface {
	Node
	PerformSABuild() // TO-DO
}

type Arguments struct {
	Token *token.Token
	Exprs []Expression // MARKING
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
func (args *Arguments) GetType() {}

type LValue struct {
	Token  *token.Token
	Ident  IdentLiteral
	Idents []IdentLiteral
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
func (lv *LValue) GetType() {}

type Expression struct {
	Token  *token.Token
	Left   *BoolTerm
	Rights []BoolTerm
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
func (exp *Expression) GetType() {}

type BoolTerm struct {
	Token  *token.Token
	Left   *EqualTerm
	Rights []EqualTerm
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
func (bt *BoolTerm) GetType() {}

type EqualTerm struct {
	Token         *token.Token
	Left          *RelationTerm
	EqualOperator []string // '=='|'!='
	Rights        []RelationTerm
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
func (et *EqualTerm) GetType() {}

type RelationTerm struct {
	Token             *token.Token
	Left              *SimpleTerm
	RelationOperators []string // '>'| '<' | '<=' | '>='
	Rights            []SimpleTerm
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
func (rt *RelationTerm) GetType() {}

type SimpleTerm struct {
	Token               *token.Token
	Left                *Term
	SimpleTermOperators []string // '+' | '-'
	Rights              []Term
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
func (st *SimpleTerm) GetType() {}

type Term struct {
	Token         *token.Token
	Left          *UnaryTerm
	TermOperators []string // '*' | '/'
	Rights        []UnaryTerm
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
func (t *Term) GetType() {}

type UnaryTerm struct {
	Token         *token.Token
	UnaryOperator string // '!' | '-' | '' <- default
	SelectorTerm  *SelectorTerm
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
func (ut *UnaryTerm) GetType() {}

type SelectorTerm struct {
	Token  *token.Token
	Fact   *Factor
	Idents []IdentLiteral
}

func (st *SelectorTerm) TokenLiteral() string {
	if st.Token != nil {
		return st.Token.Literal
	}
	panic("Could not determine token literals for selectorTerm")
}
func (st *SelectorTerm) String() string {
	out := bytes.Buffer{}
	out.WriteString(st.Fact.String())
	for _, id := range st.Idents {
		out.WriteString(".")
		out.WriteString(id.String())
	}
	return out.String()
}
func (st *SelectorTerm) GetType() {}

type Factor struct {
	Token *token.Token
	Expr  Node
}

func (f *Factor) TokenLiteral() string {
	if f.Token != nil {
		return f.Token.Literal
	}
	panic("Could not determine token literal for factor")
}
func (f *Factor) String() string {
	out := bytes.Buffer{}
	out.WriteString(f.Expr.String()) // TO-DO
	return out.String()
}
func (f *Factor) GetType() {}

func NewArgs(exprs []Expression) *Arguments                       { return &Arguments{nil, exprs} }
func NewLvalue(ident IdentLiteral, idents []IdentLiteral) *LValue { return &LValue{nil, ident, idents} }
func NewExpression(l *BoolTerm, rs []BoolTerm) *Expression {
	return &Expression{nil, l, rs}
}
func NewBoolTerm(l *EqualTerm, rs []EqualTerm) *BoolTerm { return &BoolTerm{nil, l, rs} }
func NewEqualTerm(l *RelationTerm, operators []string, rs []RelationTerm) *EqualTerm {
	return &EqualTerm{nil, l, operators, rs}
}
func NewRelationTerm(l *SimpleTerm, operators []string, rs []SimpleTerm) *RelationTerm {
	return &RelationTerm{nil, l, operators, rs}
}
func NewSimpleTerm(l *Term, operators []string, rs []Term) *SimpleTerm {
	return &SimpleTerm{nil, l, operators, rs}
}
func NewTerm(l *UnaryTerm, operators []string, rs []UnaryTerm) *Term {
	return &Term{nil, l, operators, rs}
}
func NewUnaryTerm(operator string, selectorTerm *SelectorTerm) *UnaryTerm {
	return &UnaryTerm{nil, operator, selectorTerm}
}
func NewSelectorTerm(factor *Factor, idents []IdentLiteral) *SelectorTerm {
	return &SelectorTerm{nil, factor, idents}
}
func NewFactor(expr Node) *Factor { return &Factor{nil, expr} }

/******* Single Literals Expr *******/

// InvocExpr : invocation in Factor ('id' [Arguments])
type InvocExpr struct {
	Token     *token.Token
	Ident     IdentLiteral
	InnerArgs *Arguments
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
func (ie *InvocExpr) GetType() {}

// PriorityExpression : '(' Expression ')' (inside Factor)
type PriorityExpression struct {
	Token           *token.Token
	InnerExpression *Expression
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
func (pe *PriorityExpression) GetType() {}

// NilNode : nil (keyword "nil")
type NilNode struct {
	Token *token.Token
}

func (n *NilNode) TokenLiteral() string { return n.Token.Literal }
func (n *NilNode) String() string       { return n.Token.Literal }
func (n *NilNode) GetType()             {}

// BoolLiteral : True/False
type BoolLiteral struct {
	Token *token.Token
	Value bool
}

func (bl *BoolLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BoolLiteral) String() string       { return bl.Token.Literal }
func (bl *BoolLiteral) GetType()             {}

// IntLiteral : number (integer)
type IntLiteral struct {
	Token *token.Token
	Value int64
}

func (il *IntLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntLiteral) String() string       { return il.Token.Literal }
func (il *IntLiteral) GetType()             {}

// IdentLiteral : identifier
type IdentLiteral struct {
	Token *token.Token
	Id    string
}

func (idl *IdentLiteral) TokenLiteral() string { return idl.Token.Literal }
func (idl *IdentLiteral) String() string       { return idl.Token.Literal }
func (idl *IdentLiteral) GetType()             {}
