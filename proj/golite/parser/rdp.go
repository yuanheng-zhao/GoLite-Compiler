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
func (p *Parser) Parse() *ast.Factor {
	return factor(p) // QI TOU starting
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
		if argu == nil {
			node = &ast.IdentLiteral{Token: &identTok, Id: identTok.Literal}
		} else {
			node = &ast.InvocExpr{Token: &identTok, InnerArgs: argu}
		}
	} else if lpTok, match := p.match(ct.LPAREN); match {
		expr := expression(p)
		if expr != nil {
			if _, match := p.match(ct.RPAREN); match {
				node = &ast.PriorityExpression{Token: &lpTok, InnerExpression: expr}
			}
		}
	}

	// to do: verify the value of declared node
	if node != nil {
		//return &node
		return ast.NewFactor(node)
	} else {
		return nil
	}
}
