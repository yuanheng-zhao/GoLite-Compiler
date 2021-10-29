package main

import (
	"bufio"
	"fmt"
	"os"
	"proj/golite/scanner"
	"proj/golite/token"
	s "strings"
)

func main() {
	//IMPLEMENT ME!
	usageStatement := "To read in a Golite program and print out each token\n	Usage: go run golite.go -<FLAG> <FILENAME>"
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(usageStatement)
		os.Exit(0)
	}
	flag := s.TrimSpace(args[0])
	filename := s.TrimSpace(args[1])

	if flag == "-lex" {
		f_obj, err := os.Open(filename)
		if err != nil { // the filename should be valid
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		reader := bufio.NewReader(f_obj)
		my_scanner := scanner.New(reader)
		var tok token.TokenType
		for tok != token.EOF {
			tok = my_scanner.NextToken().Type
			fmt.Println(tok)
		}
	}

}
