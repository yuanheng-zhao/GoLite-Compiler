package scanner

import (
	"proj/golite/token"
)

type Scanner struct {
	test map[rune]token.TokenType
}
