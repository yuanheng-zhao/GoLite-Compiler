package parser

import (
	"fmt"
	ct "proj/golite/context"
	"proj/golite/scanner"
	tk "proj/golite/token"
	"testing"
)

func Test1(t *testing.T) {
	ctx := ct.New(false, false, "test1_parser.golite")
	myScanner := scanner.New(*ctx)

	// The expected result struct represents the token stream for the input source
	tokens := []tk.Token{
		{tk.PACK, "package", 1},
		{tk.ID, "main", 1},
		{tk.SEMICOLON, ";", 1},

		{tk.IMPORT, "import", 2},
		{tk.QTDMARK, "\"", 2},
		{tk.FMT, "fmt", 2},
		{tk.QTDMARK, "\"", 2},
		{tk.SEMICOLON, ";", 2},

		{tk.FUNC, "func", 3},
		{tk.ID, "main", 3},
		{tk.LPAREN, "(", 3},
		{tk.RPAREN, ")", 3},
		{tk.LBRACE, "{", 3},

		{tk.VAR, "var", 4},
		{tk.ID, "a", 4},
		{tk.INT, "int", 4},
		{tk.SEMICOLON, ";", 4},

		{tk.ID, "a", 5},
		{tk.ASSIGN, "=", 5},
		{tk.NUM, "1", 5},
		{tk.ADD, "+", 5},
		{tk.NUM, "1", 5},
		{tk.SEMICOLON, ";", 5},

		{tk.RBRACE, "}", 6},
	}

	parser := New(*myScanner)
	fmt.Println(parser.tokens)
	ast := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}

	fmt.Println("AST Printout:")
	fmt.Println(ast.String())
}

func Test2(t *testing.T) {
	ctx := ct.New(false, false, "test2_parser.golite")
	myScanner := scanner.New(*ctx)

	// The expected result struct represents the token stream for the input source
	tokens := []tk.Token{
		{tk.PACK, "package", 1},
		{tk.ID, "main", 1},
		{tk.SEMICOLON, ";", 1},

		{tk.IMPORT, "import", 2},
		{tk.QTDMARK, "\"", 2},
		{tk.FMT, "fmt", 2},
		{tk.QTDMARK, "\"", 2},
		{tk.SEMICOLON, ";", 2},

		{tk.FUNC, "func", 3},
		{tk.ID, "main", 3},
		{tk.LPAREN, "(", 3},
		{tk.RPAREN, ")", 3},
		{tk.LBRACE, "{", 3},

		{tk.VAR, "var", 4},
		{tk.ID, "b", 4},
		{tk.INT, "bool", 4},
		{tk.SEMICOLON, ";", 4},

		{tk.ID, "b", 5},
		{tk.ASSIGN, "=", 5},
		{tk.TRUE, "true", 5},
		{tk.OR, "||", 5},
		{tk.FALSE, "false", 5},
		{tk.AND, "&&", 5},
		{tk.LPAREN, "(", 5},
		{tk.TRUE, "true", 5},
		{tk.EQUAL, "==", 5},
		{tk.TRUE, "true", 5},
		{tk.RPAREN, ")", 5},
		{tk.SEMICOLON, ";", 5},

		{tk.RBRACE, "}", 6},
	}

	parser := New(*myScanner)
	ast := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
	fmt.Println("AST Printout:")
	fmt.Println(ast.String())
}

func Test3(t *testing.T) {
	ctx := ct.New(false, false, "test3_parser.golite")
	myScanner := scanner.New(*ctx)

	parser := New(*myScanner)
	ast := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", "Ignore (Too long)", "Valid AST", nil)
	}

	fmt.Println(ast.String())
}
