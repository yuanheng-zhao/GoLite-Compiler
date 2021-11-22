package parser

import (
	"bufio"
	"os"
	"proj/golite/scanner"
	ct "proj/golite/token"
	"testing"
)

func Test1(t *testing.T) {

	f_obj, _ := os.Open("test1_parser.golite")
	reader := bufio.NewReader(f_obj)
	myScanner := scanner.New(reader)
	// The expected result struct represents the token stream for the input source
	tokens := []ct.Token{
		{ct.LPAREN, "(", 1},
		{ct.ID, "a", 1},
		{ct.OR, "||", 1},
		{ct.ID, "b", 1},
		{ct.AND, "&&", 1},
		{ct.ID, "c", 1},
		{ct.EQUAL, "==", 1},
		{ct.TRUE, "true", 1},
		{ct.RPAREN, ")", 1},
	}

	// Define  a new scanner for some Cal program
	parser := New(*myScanner)
	ast := parser.Parse()
	if ast == nil {
		t.Errorf("\nParse(%v)\nExpected:%v\nGot:%v", tokens, "Valid AST", nil)
	}
}