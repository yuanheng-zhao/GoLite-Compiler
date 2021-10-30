package token

import (
	"fmt"
)

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	INT     = "INT"
	NUM     = "NUM" // an integer containing one or more digits (0-9)
	BOOL    = "BOOL"
	TRUE    = "TRUE"
	FALSE   = "FALSE"
	ID      = "ID"
	NIL     = "NIL"

	LET     = "LET"
	PRINT   = "PRINT"
	PRINTLN = "PRINTLN"
	RETURN  = "RETURN"
	PACK    = "PACK"
	IMPORT  = "IMPORT"
	FMT     = "FMT"
	TYPE    = "TYPE"
	STRUCT  = "STRUCT"
	SCAN    = "SCAN"
	IF      = "IF"
	ELSE    = "ELSE"
	FOR     = "FOR"
	FUNC    = "FUNC"
	VAR     = "VAR"

	DOT     = "DOT"
	COMMA   = "COMMA"
	QTDMARK = "QTDMARK"
	LBRACE  = "LBRACE"
	RBRACE  = "RBRACE"
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"

	ASSIGN    = "ASSIGN"
	AMPERS    = "AMPERS" // for getting address
	SEMICOLON = "SEMICOLON"
	ADD       = "ADD"
	MINUS     = "MINUS"
	MULTIPLY  = "MULTIPLY"
	DIVIDE    = "DIVIDE"
	OR        = "OR"
	AND       = "AND"
	NOT       = "NOT"
	EQUAL     = "EQUAL"
	NEQUAL    = "NEQUAL"
	GT        = "GT"
	GE        = "GE"
	LT        = "LT"
	LE        = "LE"
	COMMENT   = "COMMENT"
)

type Token struct {
	Type    TokenType
	Literal string
	//type_lookup string
}

func (tok Token) String() string {
	return fmt.Sprintf("Token.%v(%v)", tok.Type, tok.Literal)
}
