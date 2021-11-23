package parser

import (
	"flag"
	"fmt"
	"proj/golite/ast"
	"proj/golite/scanner"
	ct "proj/golite/token"
	"strconv"
)

//Parser includes all fields necessary to perform recursive decent parsing
type Parser struct {
	scanner   scanner.Scanner
	currToken ct.Token
	errFound  bool
}

//New creates and initializes a new parser
func New(scanner scanner.Scanner) *Parser {
	parser := &Parser{}
	parser.scanner = scanner
	parser.currToken = scanner.NextToken()
	return parser
}

func (p *Parser) parseError(msg string) {
	out := flag.CommandLine.Output()
	fmt.Fprintf(out, "syntax error: #{msg}\n")
	p.errFound = true
}

func (p *Parser) match(token ct.TokenType) (ct.Token, bool) {
	lineNum := p.currToken.LineNum
	if token == p.currToken.Type {
		token := p.currToken
		p.currToken = p.scanner.NextToken()
		return token, true
	}
	return ct.Token{ct.ILLEGAL, "", lineNum}, false
}

func (p *Parser) expect(token ct.TokenType) bool {
	if _, match := p.match(token); match {
		return true
	}
	p.parseError("unexpected symbol error. Found: #{p.currToken.Type}, Expected: #{token.Type}")
	return false
}

func (p *Parser) Parse() *ast.Program {
	return program(p) // QI TOU starting
}

func program(p *Parser) *ast.Program {
	pac := packageStmt(p)
	if pac == nil {
		return nil
	}
	imp := importStmt(p)
	if imp == nil {
		return nil
	}
	typ := types(p)
	if typ == nil {
		return nil
	}
	decs := declarations(p)
	if decs == nil {
		return nil
	}
	funs := functions(p)
	if funs == nil {
		return nil
	}
	if p.currToken.Type != ct.EOF {
		p.parseError(fmt.Sprintf("Expected end of file but found :#{p.currToken.Literal}"))
	}
	if p.errFound == false {
		return ast.NewProgram(pac, imp, typ, decs, funs)
	}
	return nil
}

func packageStmt(p *Parser) *ast.Package {
	var pacTok, idTok ct.Token
	var pacMatch, idMatch bool

	if pacTok, pacMatch = p.match(ct.PACK); !pacMatch {
		return nil
	}
	if idTok, idMatch = p.match(ct.ID); !idMatch {
		return nil
	}
	if _, scMatch := p.match(ct.SEMICOLON); !scMatch {
		return nil
	}

	node := ast.NewPackage(ast.IdentLiteral{&idTok, idTok.Literal})
	node.Token = &pacTok
	return node
}

func importStmt(p *Parser) *ast.Import {
	var impTok, fmtTok ct.Token
	var impMatch, fmtMatch bool

	if impTok, impMatch = p.match(ct.IMPORT); !impMatch {
		return nil
	}
	if _, lQtdMatch := p.match(ct.QTDMARK); !lQtdMatch {
		return nil
	}
	if fmtTok, fmtMatch = p.match(ct.FMT); !fmtMatch {
		return nil
	}
	if _, rQtdMatch := p.match(ct.QTDMARK); !rQtdMatch {
		return nil
	}
	if _, scMatch := p.match(ct.SEMICOLON); !scMatch {
		return nil
	}

	node := ast.NewImport(ast.IdentLiteral{&fmtTok, fmtTok.Literal})
	node.Token = &impTok
	return node
}

func types(p *Parser) *ast.Types {
	var typedecs []ast.TypeDeclaration

	for {
		typedec := typeDeclaration(p)
		if typedec != nil {
			typedecs = append(typedecs, *typedec)
		} else {
			break
		}
	}

	node := ast.NewTypes(typedecs)
	return node
}

func typeDeclaration(p *Parser) *ast.TypeDeclaration {
	var typTok, idTok ct.Token
	var typMac, idMac, structMac, lbrMac, rbrMac, scMac bool
	if typTok, typMac = p.match(ct.TYPE); !typMac {
		return nil
	}
	if idTok, idMac = p.match(ct.ID); !idMac {
		return nil
	}
	if _, structMac = p.match(ct.STRUCT); !structMac {
		return nil
	}
	if _, lbrMac = p.match(ct.LBRACE); !lbrMac {
		return nil
	}
	astFields := fields(p)
	if astFields == nil {
		return nil
	}
	if _, rbrMac = p.match(ct.RBRACE); !rbrMac {
		return nil
	}
	if _, scMac = p.match(ct.SEMICOLON); !scMac {
		return nil
	}

	node := ast.NewTypeDeclaration(ast.IdentLiteral{&idTok, idTok.Literal}, astFields)
	node.Token = &typTok
	return node
}

func fields(p *Parser) *ast.Fields {
	declFirst := decl(p)
	var decls []ast.Decl

	if declFirst == nil {
		return nil
	}
	if _, match := p.match(ct.SEMICOLON); !match {
		return nil
	}

	decls = append(decls, *declFirst)
	for {
		dec := decl(p)
		if dec != nil {
			decls = append(decls, *dec)
		} else {
			break
		}
		if _, match := p.match(ct.SEMICOLON); !match {
			return nil
		}
	}

	node := ast.NewFields(decls)
	return node
}

func decl(p *Parser) *ast.Decl {
	var idTok ct.Token
	var match bool
	if idTok, match = p.match(ct.ID); !match {
		return nil
	}
	astType := typeExpression(p)
	if astType == nil {
		return nil
	}
	node := ast.NewDecl(ast.IdentLiteral{&idTok, idTok.Literal}, astType)
	node.Token = &idTok
	return node
}

func typeExpression(p *Parser) *ast.Type {
	var node *ast.Type
	var typeTok ct.Token
	var match bool

	if typeTok, match = p.match(ct.INT); match {
		node = ast.NewType(typeTok.Literal)
		node.Token = &typeTok
	} else if typeTok, match := p.match(ct.BOOL); match {
		node = ast.NewType(typeTok.Literal)
		node.Token = &typeTok
	} else if typeTok, match := p.match(ct.MULTIPLY); match {
		if idTok, idMatch := p.match(ct.ID); idMatch {
			node = ast.NewType(typeTok.Literal + idTok.Literal)
			node.Token = &typeTok
		}
	}

	if node != nil && node.Token != nil {
		return node
	}
	return nil
}

func declarations(p *Parser) *ast.Declarations {
	var decs []ast.Declaration

	for {
		dec := declaration(p)
		if dec != nil {
			decs = append(decs, *dec)
		} else {
			break
		}
	}

	node := ast.NewDeclarations(decs)
	return node
}

func declaration(p *Parser) *ast.Declaration {
	var varmatch, scmatch bool
	var varTok ct.Token

	if varTok, varmatch = p.match(ct.VAR); !varmatch {
		return nil
	}
	idTok := ids(p)
	if idTok == nil {
		return nil
	}
	typeTok := typeExpression(p)
	if typeTok == nil {
		return nil
	}
	if _, scmatch = p.match(ct.SEMICOLON); !scmatch {
		return nil
	}

	node := ast.NewDeclaration(idTok, typeTok)
	node.Token = &varTok
	return node
}

func ids(p *Parser) *ast.Ids {
	var idTokFirst ct.Token
	var idMatchFirst bool
	var ids []ast.IdentLiteral

	if idTokFirst, idMatchFirst = p.match(ct.ID); !idMatchFirst {
		return nil
	}

	ids = append(ids, ast.IdentLiteral{&idTokFirst, idTokFirst.Literal})
	for {
		if _, match := p.match(ct.COMMA); !match {
			break
		}
		if idTok, match := p.match(ct.ID); match {
			ids = append(ids, ast.IdentLiteral{&idTok, idTok.Literal})
		} else {
			return nil
		}
	}

	node := ast.NewIds(ids)
	node.Token = &idTokFirst
	return node
}

func functions(p *Parser) *ast.Functions {
	var funs []ast.Function

	for {
		if fun := function(p); fun != nil {
			funs = append(funs, *fun)
		} else {
			break
		}
	}

	node := ast.NewFunctions(funs)
	return node
}

func function(p *Parser) *ast.Function {
	var funcTok, idTok ct.Token
	var funcMatch, idMatch bool
	if funcTok, funcMatch = p.match(ct.FUNC); !funcMatch {
		return nil
	}
	if idTok, idMatch = p.match(ct.ID); !idMatch {
		return nil
	}
	paras := parameters(p)
	if paras == nil {
		return nil
	}
	retTyp := returnType(p)
	if retTyp == nil {
		return nil
	}
	if _, lbraceMatch := p.match(ct.LBRACE); !lbraceMatch {
		return nil
	}
	decls := declarations(p)
	if decls == nil {
		return nil
	}
	stmts := statements(p)
	if stmts == nil {
		return nil
	}
	if _, rbraceMatch := p.match(ct.RBRACE); !rbraceMatch {
		return nil
	}

	node := ast.NewFunction(ast.IdentLiteral{&idTok, idTok.Literal}, paras, retTyp, decls, stmts)
	node.Token = &funcTok
	return node
}

func parameters(p *Parser) *ast.Parameters {
	var lParenTok ct.Token
	var lParenMatch bool
	if lParenTok, lParenMatch = p.match(ct.LPAREN); !lParenMatch {
		return nil
	}

	var decls []ast.Decl

	declFirst := decl(p)
	if declFirst != nil {
		decls = append(decls, *declFirst)
		for {
			if _, commaMatch := p.match(ct.COMMA); !commaMatch {
				break
			}
			if decl := decl(p); decl != nil {
				decls = append(decls, *decl)
			} else {
				return nil
			}
		}
	}
	if _, match := p.match(ct.RPAREN); !match {
		return nil
	}

	node := ast.NewParameters(decls)
	node.Token = &lParenTok
	return node
}

func returnType(p *Parser) *ast.ReturnType {
	var node *ast.ReturnType
	typTok := typeExpression(p)

	if typTok != nil {
		node = ast.NewReturnType(typTok.TokenLiteral())
	} else {
		node = ast.NewReturnType("")
	}
	return node
}

func statements(p *Parser) *ast.Statements {
	var stmts []ast.Statement
	for {
		if stmt := statement(p); stmt != nil {
			stmts = append(stmts, *stmt)
		} else {
			break
		}
	}
	node := ast.NewStatements(stmts)
	return node
}

func statement(p *Parser) *ast.Statement {
	// to do, adapt with backtracking
	p.scanner.RollbackReset()
	bloc := block(p)
	if bloc != nil {
		return ast.NewStatement(bloc)
	}

	p.scanner.Rollback()
	assi := assignment(p)
	if assi != nil {
		return ast.NewStatement(assi)
	}

	p.scanner.Rollback()
	prin := print(p)
	if prin != nil {
		return ast.NewStatement(prin)
	}

	p.scanner.Rollback()
	cond := conditional(p)
	if cond != nil {
		return ast.NewStatement(cond)
	}

	p.scanner.Rollback()
	loopAst := loop(p)
	if loopAst != nil {
		return ast.NewStatement(loopAst)
	}

	p.scanner.Rollback()
	ret := returnStmt(p)
	if ret != nil {
		return ast.NewStatement(ret)
	}

	p.scanner.Rollback()
	readAst := read(p)
	if readAst != nil {
		return ast.NewStatement(readAst)
	}

	p.scanner.Rollback()
	invoc := invocation(p)
	if invoc != nil {
		return ast.NewStatement(invoc)
	}
	return nil
}

func block(p *Parser) *ast.Block {
	var lbraceTok ct.Token
	var lbraceMatch bool
	if lbraceTok, lbraceMatch = p.match(ct.LBRACE); !lbraceMatch {
		return nil
	}
	stmtsTok := statements(p)
	if stmtsTok == nil {
		return nil
	}
	if _, match := p.match(ct.RBRACE); !match {
		return nil
	}

	node := ast.NewBlock(stmtsTok)
	node.Token = &lbraceTok
	return node
}

func assignment(p *Parser) *ast.Assignment {
	lval := lValue(p)
	if lval == nil {
		return nil
	}
	if _, match := p.match(ct.ASSIGN); !match {
		return nil
	}
	expr := expression(p)
	if expr == nil {
		return nil
	}
	if _, match := p.match(ct.SEMICOLON); !match {
		return nil
	}
	node := ast.NewAssignment(lval, expr)
	return node
}

func read(p *Parser) *ast.Read {
	var fmtTok, idTok ct.Token
	var fmtMatch, idMatch bool

	if fmtTok, fmtMatch = p.match(ct.FMT); !fmtMatch {
		return nil
	}
	if _, match := p.match(ct.DOT); !match {
		return nil
	}
	if _, match := p.match(ct.SCAN); !match {
		return nil
	}
	if _, match := p.match(ct.LPAREN); !match {
		return nil
	}
	if _, match := p.match(ct.AMPERS); !match {
		return nil
	}
	if idTok, idMatch = p.match(ct.ID); !idMatch {
		return nil
	}
	if _, match := p.match(ct.RPAREN); !match {
		return nil
	}
	if _, match := p.match(ct.SEMICOLON); !match {
		return nil
	}

	node := ast.NewRead(ast.IdentLiteral{&idTok, idTok.Literal})
	node.Token = &fmtTok
	return node
}

func print(p *Parser) *ast.Print {
	var fmtTok, printTok, idTok ct.Token
	var fmtMatch, printMatch, idMatch bool

	if fmtTok, fmtMatch = p.match(ct.FMT); !fmtMatch {
		return nil
	}
	if _, match := p.match(ct.DOT); !match {
		return nil
	}
	printTok, printMatch = p.match(ct.PRINT)
	if !printMatch {
		printTok, printMatch = p.match(ct.PRINTLN)
	}
	if !printMatch {
		return nil
	}
	if _, match := p.match(ct.LPAREN); !match {
		return nil
	}
	if idTok, idMatch = p.match(ct.ID); !idMatch {
		return nil
	}
	if _, match := p.match(ct.RPAREN); !match {
		return nil
	}
	if _, match := p.match(ct.SEMICOLON); !match {
		return nil
	}

	node := ast.NewPrint(printTok.Literal, ast.IdentLiteral{&idTok, idTok.Literal})
	node.Token = &fmtTok
	return node
}

func conditional(p *Parser) *ast.Conditional {
	var ifTok ct.Token
	var ifMatch bool
	var node *ast.Conditional

	if ifTok, ifMatch = p.match(ct.IF); !ifMatch {
		return nil
	}
	if _, match := p.match(ct.LPAREN); !match {
		return nil
	}
	expr := expression(p)
	if expr == nil {
		return nil
	}
	if _, match := p.match(ct.RPAREN); !match {
		return nil
	}
	bloc := block(p)
	if bloc == nil {
		return nil
	}

	var elsBloc *ast.Block
	if _, match := p.match(ct.ELSE); match {
		elsBloc = block(p)
		if elsBloc == nil {
			return nil
		}
	}
	node = ast.NewConditional(expr, bloc, elsBloc)
	node.Token = &ifTok

	return node
}

func loop(p *Parser) *ast.Loop {
	var forTok ct.Token
	var forMatch bool

	if forTok, forMatch = p.match(ct.FOR); !forMatch {
		return nil
	}
	if _, match := p.match(ct.LPAREN); !match {
		return nil
	}
	expr := expression(p)
	if expr == nil {
		return nil
	}
	if _, match := p.match(ct.RPAREN); !match {
		return nil
	}
	bloc := block(p)
	if bloc == nil {
		return nil
	}

	node := ast.NewLoop(expr, bloc)
	node.Token = &forTok
	return node
}

func returnStmt(p *Parser) *ast.Return {
	var node *ast.Return
	var retTok ct.Token
	var retMatch bool

	if retTok, retMatch = p.match(ct.RETURN); !retMatch {
		return nil
	}
	expr := expression(p)
	if _, match := p.match(ct.SEMICOLON); !match {
		return nil
	}

	node = ast.NewReturn(expr)
	node.Token = &retTok

	return node
}

func invocation(p *Parser) *ast.Invocation {
	var idTok ct.Token
	var idMatch bool
	if idTok, idMatch = p.match(ct.ID); !idMatch {
		return nil
	}
	arg := arguments(p)
	if arg == nil {
		return nil
	}
	if _, match := p.match(ct.SEMICOLON); !match {
		return nil
	}

	node := ast.NewInvocation(ast.IdentLiteral{&idTok, idTok.Literal}, arg)
	node.Token = &idTok
	return node
}

func arguments(p *Parser) *ast.Arguments {
	var lParentok ct.Token
	var lParenMatch bool
	var exprs []ast.Expression

	if lParentok, lParenMatch = p.match(ct.LPAREN); !lParenMatch {
		return nil
	}
	exprFirst := expression(p)
	if exprFirst != nil {
		exprs = append(exprs, *exprFirst)
		for {
			if _, match := p.match(ct.COMMA); !match {
				break
			}
			expr := expression(p)
			if expr != nil {
				exprs = append(exprs, *expr)
			} else {
				return nil
			}
		}
	}
	if _, rParenMatch := p.match(ct.RPAREN); !rParenMatch {
		return nil
	}

	node := ast.NewArgs(exprs)
	node.Token = &lParentok

	return node
}

func lValue(p *Parser) *ast.LValue {
	var idTok ct.Token
	var idMatch bool
	var ids []ast.IdentLiteral

	if idTok, idMatch = p.match(ct.ID); !idMatch {
		return nil
	}
	for {
		if _, match := p.match(ct.DOT); !match {
			break
		}
		if id, match := p.match(ct.ID); match {
			ids = append(ids, ast.IdentLiteral{&id, id.Literal})
		} else {
			return nil
		}
	}

	node := ast.NewLvalue(ast.IdentLiteral{&idTok, idTok.Literal}, ids)
	node.Token = &idTok
	return node
}

func expression(p *Parser) *ast.Expression {
	var bts []ast.BoolTerm
	btLeft := boolTerm(p)
	if btLeft == nil {
		return nil
	}
	for {
		if _, match := p.match(ct.OR); !match {
			break
		}
		btRight := boolTerm(p)
		if btRight != nil {
			bts = append(bts, *btRight)
		} else {
			return nil
		}
	}

	node := ast.NewExpression(btLeft, bts)
	return node
}

func boolTerm(p *Parser) *ast.BoolTerm {
	var ets []ast.EqualTerm
	etLeft := equalTerm(p)

	if etLeft == nil {
		return nil
	}
	for {
		if _, match := p.match(ct.AND); !match {
			break
		}
		etRight := equalTerm(p)
		if etRight != nil {
			ets = append(ets, *etRight)
		} else {
			return nil
		}
	}

	node := ast.NewBoolTerm(etLeft, ets)
	return node
}

func equalTerm(p *Parser) *ast.EqualTerm {
	var eqOps []string
	var rts []ast.RelationTerm
	var eqTok ct.Token
	var match bool

	rtLeft := relationTerm(p)
	if rtLeft == nil {
		return nil
	}

	for {
		if eqTok, match = p.match(ct.EQUAL); !match {
			if eqTok, match = p.match(ct.NEQUAL); !match {
				break
			}
		}
		eqOps = append(eqOps, eqTok.Literal)
		rtRight := relationTerm(p)
		if rtRight != nil {
			rts = append(rts, *rtRight)
		} else {
			return nil
		}
	}

	node := ast.NewEqualTerm(rtLeft, eqOps, rts)
	return node
}

func relationTerm(p *Parser) *ast.RelationTerm {
	var rlOps []string
	var sts []ast.SimpleTerm
	var rlTok ct.Token
	var match bool

	stLeft := simpleTerm(p)
	if stLeft == nil {
		return nil
	}

	for {
		if rlTok, match = p.match(ct.GT); !match {
			if rlTok, match = p.match(ct.LT); !match {
				if rlTok, match = p.match(ct.GE); !match {
					if rlTok, match = p.match(ct.LE); !match {
						break
					}
				}
			}
		}
		rlOps = append(rlOps, rlTok.Literal)
		stRight := simpleTerm(p)
		if stRight != nil {
			sts = append(sts, *stRight)
		} else {
			return nil
		}
	}

	node := ast.NewRelationTerm(stLeft, rlOps, sts)
	return node
}

func simpleTerm(p *Parser) *ast.SimpleTerm {
	var stOps []string
	var tms []ast.Term
	var stTok ct.Token
	var match bool

	termLeft := term(p)
	if termLeft == nil {
		return nil
	}
	for {
		if stTok, match = p.match(ct.ADD); !match {
			if stTok, match = p.match(ct.MINUS); !match {
				break
			}
		}
		stOps = append(stOps, stTok.Literal)
		tmRight := term(p)
		if tmRight != nil {
			tms = append(tms, *tmRight)
		} else {
			return nil
		}
	}

	node := ast.NewSimpleTerm(termLeft, stOps, tms)
	return node
}

func term(p *Parser) *ast.Term {
	var tmOps []string
	var uts []ast.UnaryTerm
	var tmTok ct.Token
	var match bool

	utLeft := unaryTerm(p)
	if utLeft == nil {
		return nil
	}
	for {
		if tmTok, match = p.match(ct.MULTIPLY); !match {
			if tmTok, match = p.match(ct.DIVIDE); !match {
				break
			}
		}
		tmOps = append(tmOps, tmTok.Literal)
		utRight := unaryTerm(p)
		if utRight != nil {
			uts = append(uts, *utRight)
		} else {
			return nil
		}
	}

	node := ast.NewTerm(utLeft, tmOps, uts)
	return node
}

func unaryTerm(p *Parser) *ast.UnaryTerm {
	op := ""
	var uniOp ct.Token
	var match bool
	if uniOp, match = p.match(ct.NOT); match {
		op = uniOp.Literal
	} else if uniOp, match = p.match(ct.MINUS); match {
		op = uniOp.Literal
	}
	selTok := selectorTerm(p)
	if selTok == nil {
		return nil
	}
	node := ast.NewUnaryTerm(op, selTok)

	if op != "" {
		node.Token = &uniOp
	}

	return node
}

func selectorTerm(p *Parser) *ast.SelectorTerm {
	var ids []ast.IdentLiteral
	var idTok ct.Token
	var match bool

	facTok := factor(p)
	if facTok == nil {
		return nil
	}

	for {
		if _, match = p.match(ct.DOT); !match {
			break
		}
		if idTok, match = p.match(ct.ID); !match {
			return nil
		}
		ids = append(ids, ast.IdentLiteral{&idTok, idTok.Literal})
	}

	node := ast.NewSelectorTerm(facTok, ids)
	return node
}

func factor(p *Parser) *ast.Factor {
	var node ast.Expr

	if numTok, match := p.match(ct.NUM); match {
		val, _ := strconv.ParseInt(numTok.Literal, 10, 64)
		node = &ast.IntLiteral{Token: &numTok, Value: val}
	} else if truTok, match := p.match(ct.TRUE); match {
		node = &ast.IdentLiteral{Token: &truTok, Id: truTok.Literal}
	} else if flsTok, match := p.match(ct.FALSE); match {
		node = &ast.IdentLiteral{Token: &flsTok, Id: flsTok.Literal}
	} else if nilTok, match := p.match(ct.NIL); match {
		node = &ast.NilNode{Token: &nilTok}
	} else if identTok, match := p.match(ct.ID); match {
		argu := arguments(p)
		idl := &ast.IdentLiteral{Token: &identTok, Id: identTok.Literal}
		if argu == nil {
			node = idl
		} else {
			node = &ast.InvocExpr{Token: &identTok, Ident: *idl, InnerArgs: argu}
		}
	} else if lpTok, match := p.match(ct.LPAREN); match {
		expr := expression(p)
		if expr != nil {
			if _, match := p.match(ct.RPAREN); match {
				node = &ast.PriorityExpression{Token: &lpTok, InnerExpression: expr}
			}
		}
	}

	if node != nil {
		//return &node
		return ast.NewFactor(&node)
	} else {
		return nil
	}
}
