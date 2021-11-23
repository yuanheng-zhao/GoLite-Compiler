package parser

import (
	"bufio"
	"fmt"
	"os"
	"proj/golite/scanner"
	ct "proj/golite/token"
	"testing"
)

func Test1(t *testing.T) {
	fileObj, _ := os.Open("test1_parser.golite")
	reader := bufio.NewReader(fileObj)
	myScanner := scanner.New(reader)
	// The expected result struct represents the token stream for the input source
	tokens := []ct.Token{
		{ct.PACK, "package", 1},
		{ct.ID, "main", 1},
		{ct.SEMICOLON, ";", 1},

		{ct.IMPORT, "import", 2},
		{ct.QTDMARK, "\"", 2},
		{ct.FMT, "fmt", 2},
		{ct.QTDMARK, "\"", 2},
		{ct.SEMICOLON, ";", 2},

		{ct.FUNC, "func", 3},
		{ct.ID, "main", 3},
		{ct.LPAREN, "(", 3},
		{ct.RPAREN, ")", 3},
		{ct.LBRACE, "{", 3},

		{ct.VAR, "var", 4},
		{ct.ID, "a", 4},
		{ct.INT, "int", 4},
		{ct.SEMICOLON, ";", 4},

		{ct.ID, "a", 5},
		{ct.ASSIGN, "=", 5},
		{ct.NUM, "1", 5},
		{ct.ADD, "+", 5},
		{ct.NUM, "1", 5},
		{ct.SEMICOLON, ";", 5},

		{ct.RBRACE, "}", 6},
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

	fileObj, _ := os.Open("test2_parser.golite")
	reader := bufio.NewReader(fileObj)
	myScanner := scanner.New(reader)
	// The expected result struct represents the token stream for the input source
	tokens := []ct.Token{
		{ct.PACK, "package", 1},
		{ct.ID, "main", 1},
		{ct.SEMICOLON, ";", 1},

		{ct.IMPORT, "import", 2},
		{ct.QTDMARK, "\"", 2},
		{ct.FMT, "fmt", 2},
		{ct.QTDMARK, "\"", 2},
		{ct.SEMICOLON, ";", 2},

		{ct.FUNC, "func", 3},
		{ct.ID, "main", 3},
		{ct.LPAREN, "(", 3},
		{ct.RPAREN, ")", 3},
		{ct.LBRACE, "{", 3},

		{ct.VAR, "var", 4},
		{ct.ID, "b", 4},
		{ct.INT, "bool", 4},
		{ct.SEMICOLON, ";", 4},

		{ct.ID, "b", 5},
		{ct.ASSIGN, "=", 5},
		{ct.TRUE, "true", 5},
		{ct.OR, "||", 5},
		{ct.FALSE, "false", 5},
		{ct.AND, "&&", 5},
		{ct.LPAREN, "(", 5},
		{ct.TRUE, "true", 5},
		{ct.EQUAL, "==", 5},
		{ct.TRUE, "true", 5},
		{ct.RPAREN, ")", 5},
		{ct.SEMICOLON, ";", 5},

		{ct.RBRACE, "}", 6},
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
	fileObj, _ := os.Open("test3_parser.golite")
	reader := bufio.NewReader(fileObj)
	myScanner := scanner.New(reader)

	parser := New(*myScanner)
	ast := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", "Ignore (Too long)", "Valid AST", nil)
	}

	fmt.Println(ast.String())
}
