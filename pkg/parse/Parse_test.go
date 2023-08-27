package parse

import (
	"fmt"
	"testing"

	"github.com/DevellSoftware/jordan/machine/pkg/script/lex"
)

func TestSample1(t *testing.T) {
	code := `
let x = 10

func test() {
  print "test1"
}
  `

	lexer := lex.NewLexer(code)
	lexer.Lex()

	fmt.Println(lexer.Tokens())

	parser := NewParser(lexer.Tokens())
	result := parser.Parse()

	result.PrintErrors()

	parser.PrintTree()
}
