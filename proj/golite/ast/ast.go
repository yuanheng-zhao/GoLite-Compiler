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

/******* Stmt : Statement *******/

type Program struct {
	Token *token.Token
	//st    *st.SymbolTable

	Package      *Package
	Import       *Import
	Types        *Types
	Declarations *Declarations
	Functions    *Functions
}

func NewProgram(pac *Package, imp *Import, typ *Types, decs *Declarations, funs *Functions) *Program {
	return &Program{nil, pac, imp, typ, decs, funs}
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
func (p *Program) PerformSABuild() {}

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
func (pkg *Package) PerformSABuild() {}

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
func (imp *Import) PerformSABuild() {}

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
func (tys *Types) PerformSABuild() {}

type TypeDeclaration struct {
	Token *token.Token
	//st     *st.SymbolTable
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
func (td *TypeDeclaration) PerformSABuild() {}

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
func (fields *Fields) PerformSABuild() {}

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
func (decl *Decl) PerformSABuild() {}

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
func (ds *Declarations) PerformSABuild() {}

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
func (d *Declaration) PerformSABuild() {}

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
func (ids *Ids) PerformSABuild() {}

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
func (fs *Functions) PerformSABuild() {}

type Function struct {
	Token *token.Token
	//st           *st.SymbolTable
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
	out.WriteString(f.Declarations.String())
	out.WriteString(" ")
	out.WriteString(f.Statements.String())
	out.WriteString("}")
	out.WriteString("\n")
	return out.String()
}
func (f *Function) PerformSABuild() {}

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
	out.WriteString(params.Decls[0].String())
	remaining := params.Decls[1:]
	for _, decl := range remaining {
		out.WriteString(",")
		out.WriteString(decl.String())
	}
	out.WriteString(")")
	return out.String()
}
func (params *Parameters) PerformSABuild() {}

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
func (rt *ReturnType) PerformSABuild() {}

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
	if len(stmts.Statements) > 0 {
		out.WriteString("\n")
	}
	for _, stmt := range stmts.Statements {
		out.WriteString("\t")
		out.WriteString(stmt.String())
	}
	out.WriteString("\n")
	return out.String()
}
func (stmts *Statements) PerformSABuild() {}

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
func (s *Statement) PerformSABuild() {}

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
	out.WriteString(b.Statements.String())
	out.WriteString("}")
	return out.String()
}
func (b *Block) PerformSABuild() {}

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
	return out.String()
}
func (a *Assignment) PerformSABuild() {}

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
	return out.String()
}
func (r *Read) PerformSABuild() {}

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
	return out.String()
}
func (p *Print) PerformSABuild() {}

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
func (cond *Conditional) PerformSABuild() {}

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
func (lp *Loop) PerformSABuild() {}

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
func (ret *Return) PerformSABuild() {}

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
	out.WriteString(" ")
	out.WriteString(invoc.Args.String())
	out.WriteString(";")
	return out.String()
}
func (invoc *Invocation) PerformSABuild() {}

func NewPackage(ident IdentLiteral) *Package {
	return &Package{nil, ident}
}
func NewImport(ident IdentLiteral) *Import      { return &Import{nil, ident} }
func NewTypes(typdecs []TypeDeclaration) *Types { return &Types{nil, typdecs} }
func NewTypeDeclaration(ident IdentLiteral, fields *Fields) *TypeDeclaration {
	return &TypeDeclaration{nil, ident, fields}
}
func NewFields(decls []Decl) *Fields                   { return &Fields{nil, decls} }
func NewDecl(ident IdentLiteral, ty *Type) *Decl       { return &Decl{nil, ident, ty} }
func NewDeclarations(decs []Declaration) *Declarations { return &Declarations{nil, decs} }
func NewDeclaration(ids *Ids, Type *Type) *Declaration { return &Declaration{nil, ids, Type} }
func NewIds(idents []IdentLiteral) *Ids                { return &Ids{nil, idents} }
func NewFunctions(funs []Function) *Functions          { return &Functions{nil, funs} }
func NewFunction(ident IdentLiteral, params *Parameters, returnType *ReturnType,
	declarations *Declarations, statements *Statements) *Function {
	return &Function{nil, ident, params, returnType, declarations, statements}
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
func (t *Type) GetType() {}

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
	Expr  Expr
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

func NewType(typeLit string) *Type                                { return &Type{nil, typeLit} }
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
func NewFactor(expr *Expr) *Factor { return &Factor{nil, *expr} }

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
