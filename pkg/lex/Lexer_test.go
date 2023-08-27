package lex

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `let five = 5`

	lexer := NewLexer(input)

	lexer.Lex()

	fmt.Println(lexer.Tokens())
}

func TestIntend(t *testing.T) {
	input := `
let five = 5
  let two = 2
    let three = 3
      let four = 4
`

	lexer := NewLexer(input)

	for i := 0; i < 100000; i++ {
		lexer.Lex()
	}
}
