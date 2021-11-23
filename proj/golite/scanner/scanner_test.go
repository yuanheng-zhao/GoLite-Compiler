package scanner

import (
	"bufio"
	"os"
	"proj/golite/token"
	"testing"
)

type ExpectedResult struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func VerifyTest(t *testing.T, tests []ExpectedResult, scanner *Scanner) {

	for i, tt := range tests {
		tok := scanner.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("FAILED[%d] - incorrect token.\nexpected=%v\ngot=%v\n",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("FAILED[%d] - incorrect token literal.\nexpected=%v\ngot=%v\n",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func Test1(t *testing.T) {
	// s1 := "This is a test string"

	fileObj, _ := os.Open("test1.golite")
	reader := bufio.NewReader(fileObj)
	scanner := New(reader)

	// 	package test1;
	// import "fmt";
	// a = 3;
	expected := []ExpectedResult{
		{token.PACK, "package"},
		{token.ID, "test1"},
		{token.SEMICOLON, ";"},
		{token.COMMENT, "//"},
		{token.IMPORT, "import"},
		{token.QTDMARK, "\""},
		{token.FMT, "fmt"},
		{token.QTDMARK, "\""},
		{token.SEMICOLON, ";"},
		{token.ID, "a"},
		{token.ASSIGN, "="},
		{token.NUM, "3"},
		{token.SEMICOLON, ";"},
	}

	VerifyTest(t, expected, scanner)
}

func Test2(t *testing.T) {
	// s1 := "This is a test string"

	fileObj, _ := os.Open("test2.golite")
	reader := bufio.NewReader(fileObj)
	scanner := New(reader)

	// 	package test1;
	// import "fmt";
	// a = 3;
	expected := []ExpectedResult{
		{token.PACK, "package"},
		{token.ID, "main"},
		{token.SEMICOLON, ";"},
		{token.IMPORT, "import"},
		{token.QTDMARK, "\""},
		{token.FMT, "fmt"},
		{token.QTDMARK, "\""},
		{token.SEMICOLON, ";"},
		{token.FUNC, "func"},
		{token.ID, "main"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.VAR, "var"},
		{token.ID, "a"},
		{token.INT, "int"},
		{token.SEMICOLON, ";"},
		{token.ID, "a"},
		{token.ASSIGN, "="},
		{token.NUM, "3"},
		{token.ADD, "+"},
		{token.NUM, "4"},
		{token.ADD, "+"},
		{token.NUM, "5"},
		{token.SEMICOLON, ";"},
		{token.FMT, "fmt"},
		{token.DOT, "."},
		{token.PRINT, "Print"},
		{token.LPAREN, "("},
		{token.ID, "a"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	VerifyTest(t, expected, scanner)
}

func Test3(t *testing.T) {
	// s1 := "This is a test string"

	fileObj, _ := os.Open("test3.golite")
	reader := bufio.NewReader(fileObj)
	scanner := New(reader)

	// 	package test1;
	// import "fmt";
	// a = 3;
	expected := []ExpectedResult{
		{token.ID, "printa"},
		{token.NUM, "2"},
		{token.PRINT, "Print"},
		{token.NUM, "2"},
		{token.ID, "print"},
		{token.ID, "let666print"},
		{token.SEMICOLON, ";"},
		{token.ID, "c"},
		{token.ASSIGN, "="},
		{token.ID, "a"},
		{token.AMPERS, "&"},
		{token.ID, "b"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.ID, "a"},
		{token.AND, "&&"},
		{token.ID, "b"},
		{token.EQUAL, "=="},
		{token.FALSE, "false"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.PRINTLN, "Println"},
		{token.LPAREN, "("},
		{token.ID, "c"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.NOT, "!"},
		{token.NEQUAL, "!="},
	}

	VerifyTest(t, expected, scanner)
}

func Test4(t *testing.T) {
	// s1 := "This is a test string"

	fileObj, _ := os.Open("test4.golite")
	reader := bufio.NewReader(fileObj)
	scanner := New(reader)

	expected := []ExpectedResult{
		{token.PACK, "package"},
		{token.ID, "main"},
		{token.FUNC, "func"},
		{token.ID, "getNum"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.MULTIPLY, "*"},
		{token.ID, "set"},
		{token.DOT, "."},
		{token.ID, "IntSet"},
		{token.LBRACE, "{"},
		{token.VAR, "var"},
		{token.ID, "res"},
		{token.ID, "set"},
		{token.DOT, "."},
		{token.ID, "IntSet"},
		{token.RETURN, "return"},
		{token.AMPERS, "&"},
		{token.ID, "res"},
		{token.RBRACE, "}"},
		{token.FUNC, "func"},
		{token.ID, "main"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.VAR, "var"},
		{token.ID, "setA"},
		{token.ID, "set"},
		{token.DOT, "."},
		{token.ID, "IntSet"},
		{token.ID, "setA"},
		{token.ASSIGN, "="},
		{token.MULTIPLY, "*"},
		{token.ID, "getNum"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.RETURN, "return"},
		{token.RBRACE, "}"},
	}

	VerifyTest(t, expected, scanner)
}

func Test5(t *testing.T) {

	fileObj, _ := os.Open("../parser/test1_parser.golite")
	reader := bufio.NewReader(fileObj)
	myScanner := New(reader)
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.LPAREN, "("},
		{token.ID, "a"},
		{token.OR, "||"},
		{token.ID, "b"},
		{token.AND, "&&"},
		{token.ID, "c"},
		{token.EQUAL, "=="},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
	}

	VerifyTest(t, expected, myScanner)
}

func TestRollback1(t *testing.T) {

	fileObj, _ := os.Open("../parser/test1_parser.golite")
	reader := bufio.NewReader(fileObj)
	myScanner := New(reader)
	// The expected result struct represents the token stream for the input source
	expected := []ExpectedResult{
		{token.LPAREN, "("},
		{token.ID, "a"},
		{token.OR, "||"},
		{token.ID, "b"},
		{token.AND, "&&"},
		{token.ID, "c"},
		{token.EQUAL, "=="},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
	}

	VerifyTest(t, expected, myScanner)
	if myScanner.rollbackOffset != 18 {
		t.Errorf("\nScan file (%v)\nExpected:%v\nGot:%v", fileObj.Name(), 18, myScanner.rollbackOffset)
	}
	myScanner.RollbackReset()
	if myScanner.rollbackOffset != 0 {
		t.Errorf("\nScan file (%v)\nReset RollbackOffset\nExpected:%v\nGot:%v", fileObj.Name(), 0, myScanner.rollbackOffset)
	}
}
