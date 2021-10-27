package scanner

import (
	"bufio"
	"go/token"
	"io"
	"regexp"
)

type Scanner struct {
	reader *bufio.Reader

	// accumulated valid string, everytime the scanner returns a token
	// it will be cleaned up (re-set to an empty string)
	curr_literal string

	number_compiled *regexp.Regexp
	id_compiled     *regexp.Regexp
}

func New(input_reader *bufio.Reader) *Scanner {
	scanner := &Scanner{reader: input_reader, curr_literal: ""}
	scanner.number_compiled, _ = regexp.Compile("[0-9]+")
	scanner.id_compiled, _ = regexp.Compile("[a-zA-Z][a-zA-Z0-9]*")

	return scanner
}

func (l *Scanner) NextToken() token.Token {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// if temp exists, return temp
				// else return token type eof
			} else {
				// report the unknown error
			}
		} else {
			new_string := l.curr_literal + string(r)
			// check new_string fulfills number_compiled or id_complied
			// or check the map of special characters

			if l.number_compiled.MatchString(new_string) {

			}

		}
	}

}
