package scanner

import (
	"bufio"
	"proj/golite/token"

	// "io"
	"os"
	"testing"
)

// func Test1(t *testing.T) {
// 	// s1 := "This is a test string"

// 	f_obj, _ := os.Open("test.golite")
// 	var r *bufio.Reader
// 	r = bufio.NewReader(f_obj)
// 	c, _, _ := r.ReadRune()
// 	c, _, _ = r.ReadRune()
// 	c, _, _ = r.ReadRune()
// 	r.UnreadRune()

// 	c, _, _ = r.ReadRune()
// 	d := string(c)
// 	fmt.Println(d)
// }

// func Test2(t *testing.T) {
// 	f_obj, _ := os.Open("test.golite")
// 	var r *bufio.Reader
// 	r = bufio.NewReader(f_obj)
// 	s := New(r)
// 	x, _, err := s.reader.ReadRune()
// 	if err != nil && err == io.EOF {
// 		fmt.Println("EOF FOUND!")
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(x)
// 	}

// 	x, _, err = s.reader.ReadRune()
// 	if err != nil && err == io.EOF {
// 		fmt.Println("EOF FOUND!")
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(x)
// 	}
// }

// func Test3(t *testing.T) {
// 	test_map := make(map[string]int)
// 	test_map["\""] = 999
// 	f_obj, _ := os.Open("test.golite")
// 	var r *bufio.Reader
// 	r = bufio.NewReader(f_obj)
// 	c, _, _ := r.ReadRune()
// 	d := string(c)

// 	fmt.Println(test_map[d])

// }

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

func Test4(t *testing.T) {
	// s1 := "This is a test string"

	f_obj, _ := os.Open("test1.golite")
	reader := bufio.NewReader(f_obj)
	scanner := New(reader)

	// 	package test1;
	// import "fmt";
	// a = 3;
	expected := []ExpectedResult{
		{token.PACK, "package"},
		{token.ID, "test1"},
		{token.SEMICOLON, ";"},
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
