package scanner

import (
	ct "proj/golite/context"
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

type ExpectedResultFull struct {
	expectedType    token.TokenType
	expectedLiteral string
	expectedLineNum int
}

func VerifyTestFull(t *testing.T, tests []ExpectedResultFull, scanner *Scanner) {

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

		if tok.LineNum != tt.expectedLineNum {
			t.Fatalf("FAILED[%d] - incorrect token lineNum.\nexpected=%v\ngot=%v\n",
				i, tt.expectedLineNum, tok.LineNum)
		}
	}
}

func Test1(t *testing.T) {
	ctx := ct.New(false, false, "test1.golite")
	scanner := New(*ctx)

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

	expectedFull := []ExpectedResultFull{
		{token.PACK, "package", 1},
		{token.ID, "test1", 1},
		{token.SEMICOLON, ";", 1},
		{token.COMMENT, "//", 1},
		{token.IMPORT, "import", 3},
		{token.QTDMARK, "\"", 3},
		{token.FMT, "fmt", 3},
		{token.QTDMARK, "\"", 3},
		{token.SEMICOLON, ";", 3},
		{token.ID, "a", 4},
		{token.ASSIGN, "=", 4},
		{token.NUM, "3", 4},
		{token.SEMICOLON, ";", 4},
	}

	VerifyTest(t, expected, scanner)

	// test with line number
	scanner = New(*ctx)
	VerifyTestFull(t, expectedFull, scanner)
}

func Test2(t *testing.T) {
	ctx := ct.New(false, false, "test2.golite")
	scanner := New(*ctx)

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
	ctx := ct.New(false, false, "test3.golite")
	scanner := New(*ctx)

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
	ctx := ct.New(false, false, "test4.golite")
	scanner := New(*ctx)

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
