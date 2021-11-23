package scanner

import (
	"bufio"
	"io"
	"log"
	"proj/golite/token"
	"regexp"
)

type Scanner struct {
	reader *bufio.Reader

	// accumulated valid string, everytime the scanner returns a token
	// it will be cleaned up (re-set to an empty string)
	lexeme string

	numberCompiled *regexp.Regexp
	idCompiled     *regexp.Regexp
	whitespaces    *regexp.Regexp

	keywords map[string]token.TokenType
	symbols  map[string]token.TokenType

	isComment  bool
	lineNumber int

	// for Statement in rdp
	rollbackOffset int
}

func New(inputReader *bufio.Reader) *Scanner {
	scanner := &Scanner{reader: inputReader, lexeme: ""}
	scanner.numberCompiled, _ = regexp.Compile("^[0-9]+$")
	scanner.idCompiled, _ = regexp.Compile("^[a-zA-Z][a-zA-Z0-9]*$")
	scanner.whitespaces, _ = regexp.Compile("\\s+")
	scanner.isComment = false
	scanner.lineNumber = 1

	keywordsMap := map[string]token.TokenType{
		"int":    token.INT,
		"number": token.NUM,
		"bool":   token.BOOL,
		"true":   token.TRUE,
		"false":  token.FALSE,
		"id":     token.ID,
		"nil":    token.NIL,

		"let":     token.LET,
		"Print":   token.PRINT,
		"Println": token.PRINTLN,
		"return":  token.RETURN,
		"package": token.PACK,
		"import":  token.IMPORT,
		"fmt":     token.FMT,
		"type":    token.TYPE,
		"struct":  token.STRUCT,
		"Scan":    token.SCAN,
		"if":      token.IF,
		"else":    token.ELSE,
		"for":     token.FOR,
		"func":    token.FUNC,
		"var":     token.VAR,
	}

	symbolsMap := map[string]token.TokenType{
		".":  token.DOT,
		",":  token.COMMA,
		"\"": token.QTDMARK,
		"{":  token.LBRACE,
		"}":  token.RBRACE,
		"(":  token.LPAREN,
		")":  token.RPAREN,

		"=":  token.ASSIGN,
		"&":  token.AMPERS,
		";":  token.SEMICOLON,
		"+":  token.ADD,
		"-":  token.MINUS,
		"*":  token.MULTIPLY,
		"/":  token.DIVIDE,
		"||": token.OR,
		"&&": token.AND,
		"!":  token.NOT,
		"==": token.EQUAL,
		"!=": token.NEQUAL,
		">":  token.GT,
		">=": token.GE,
		"<":  token.LT,
		"<=": token.LE,
		"//": token.COMMENT,
	}

	scanner.keywords = keywordsMap
	scanner.symbols = symbolsMap

	scanner.rollbackOffset = 0

	return scanner
}

func (l *Scanner) NextToken() token.Token {
	for {
		r, _, err := l.reader.ReadRune()
		l.rollbackOffset += 1
		if err != nil {
			if err == io.EOF {
				// return 'eof' if we have not processed any chars as current literal (lexeme)
				if len(l.lexeme) == 0 {
					return token.Token{Type: token.EOF, Literal: "eof", LineNum: l.lineNumber}
				}
				if l.numberCompiled.MatchString(l.lexeme) {
					tempLexeme := l.lexeme
					l.lexeme = ""
					return token.Token{Type: token.NUM, Literal: tempLexeme, LineNum: l.lineNumber}
				}
				if tok, exist := l.symbols[l.lexeme]; exist {
					tempLexeme := l.lexeme
					l.lexeme = ""
					return token.Token{Type: tok, Literal: tempLexeme, LineNum: l.lineNumber}
				}
				if l.idCompiled.MatchString(l.lexeme) {
					tempLexeme := l.lexeme
					// check if it matches with some keywords (e.g. print, var)
					if tok, exist := l.keywords[l.lexeme]; exist {
						l.lexeme = ""
						return token.Token{Type: tok, Literal: tempLexeme, LineNum: l.lineNumber}
					}
					l.lexeme = ""
					return token.Token{Type: token.ID, Literal: tempLexeme, LineNum: l.lineNumber}
				}

			} else {
				// unknown error
				log.Fatal(err)
			}
		} else {
			tempLineNum := l.lineNumber
			if l.isComment {
				if r == '\n' || r == '\r' { // newline
					l.isComment = false
					rNext, _, _ := l.reader.ReadRune()
					l.rollbackOffset += 1
					if rNext != '\n' {
						l.reader.UnreadRune()
						l.rollbackOffset -= 1
					}
					l.lineNumber++
				}
				continue
			}
			currLexeme := l.lexeme + string(r)
			if currLexeme == "|" { // for the special case "||"
				l.lexeme = currLexeme
				continue
			}
			_, exist := l.symbols[currLexeme]
			if l.numberCompiled.MatchString(currLexeme) || l.idCompiled.MatchString(currLexeme) || exist {
				l.lexeme = currLexeme
				continue
			}

			if len(l.lexeme) == 0 && !l.whitespaces.MatchString(string(r)) {
				return token.Token{Type: token.ILLEGAL, Literal: "ILLEGAL", LineNum: tempLineNum}
			}

			// if r != ' ' && r != '\n' && r != '\t' {
			// 	l.reader.UnreadRune() // rollback
			// }
			// rollback if not a whitespace/newline/carriage return/etc
			if !l.whitespaces.MatchString(string(r)) {
				l.reader.UnreadRune()
				l.rollbackOffset -= 1
			} else if r == '\n' || r == '\r' { // newline
				rNext, _, _ := l.reader.ReadRune()
				l.rollbackOffset += 1
				if rNext != '\n' {
					l.reader.UnreadRune()
					l.rollbackOffset -= 1
				}
				l.lineNumber++
			}

			tempLexeme := l.lexeme
			l.lexeme = ""
			if l.numberCompiled.MatchString(tempLexeme) {
				return token.Token{Type: token.NUM, Literal: tempLexeme, LineNum: tempLineNum}
			}
			if l.idCompiled.MatchString(tempLexeme) {
				if tok, exist := l.keywords[tempLexeme]; exist {
					return token.Token{Type: tok, Literal: tempLexeme, LineNum: tempLineNum}
				}
				return token.Token{Type: token.ID, Literal: tempLexeme, LineNum: tempLineNum}
			}
			if tok, exist := l.symbols[tempLexeme]; exist {
				if tok == token.COMMENT {
					l.isComment = true
				}
				return token.Token{Type: tok, Literal: tempLexeme, LineNum: tempLineNum}
			}
		}
	}
}

func (l *Scanner) RollbackReset() {
	l.rollbackOffset = 0
}

func (l *Scanner) Rollback() {
	if l.rollbackOffset < 1 {
		// have to reset first
		return
	}
	for l.rollbackOffset > 0 {
		l.reader.UnreadRune()
		l.rollbackOffset -= 1
	}
}
