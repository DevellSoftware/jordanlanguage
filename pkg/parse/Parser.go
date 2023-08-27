package parse

import (
	"fmt"
	"strconv"

	"github.com/DevellSoftware/jordan/machine/pkg/script/lex"
)

type Parser struct {
	tokens  []lex.Token
	root    *NodeBlock
	current int
}

func NewParser(tokens []lex.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
		root:    NewNodeBlock(),
	}
}

func (p *Parser) Parse() *ParseResult {
	result := NewParseResult()

	blocks := make([]*NodeBlock, 0)
	arrays := make([]*NodeArray, 0)
	parentheses := make([]*NodeParentheses, 0)

	currentArray := func() *NodeArray {
		if len(arrays) == 0 {
			return nil
		}

		return arrays[len(arrays)-1]
	}

	currentParentheses := func() *NodeParentheses {
		if len(parentheses) == 0 {
			return nil
		}

		return parentheses[len(parentheses)-1]
	}

	currentBlock := func() *NodeBlock {
		if len(blocks) == 0 {
			return p.root
		}

		return blocks[len(blocks)-1]
	}

	for _, token := range p.tokens {
		if token.Type() == lex.TOKEN_KEYWORD {
			if currentArray() != nil {
				result.AddError(NewParseError("Unexpected keyword '" + token.Literal() + "'"))
			} else if currentParentheses() != nil {
				result.AddError(NewParseError("Unexpected keyword '" + token.Literal() + "'"))
			} else {
				currentBlock().Add(NewNodeKeyword(token.Literal()))
			}
		}

		if token.Type() == lex.TOKEN_IDENTIFIER {
			if currentArray() != nil {
				currentArray().Add(NewNodeIdentifier(token.Literal()))
			} else if currentParentheses() != nil {
				currentParentheses().Add(NewNodeIdentifier(token.Literal()))
			} else {
				currentBlock().Add(NewNodeIdentifier(token.Literal()))
			}
		}

		if token.Type() == lex.TOKEN_OPERATOR {
			if currentArray() != nil {
				result.AddError(NewParseError("Unexpected operator '" + token.Literal() + "'"))
			} else if currentParentheses() != nil {
				result.AddError(NewParseError("Unexpected operator '" + token.Literal() + "'"))
			} else {
				currentBlock().Add(NewNodeOperator(token.Literal()))
			}
		}

		if token.Type() == lex.TOKEN_NUMBER {
			value, err := strconv.ParseFloat(token.Literal(), 64)

			if err == nil {
				if currentArray() != nil {
					currentArray().Add(NewNodeNumber(value))
				} else if currentParentheses() != nil {
					currentParentheses().Add(NewNodeNumber(value))
				} else {
					currentBlock().Add(NewNodeNumber(value))
				}
			}
		}

		if token.Type() == lex.TOKEN_STRING {
			nodeString := NewNodeString(token.Literal()) // should we go with variables or direct to avoid GC

			if currentArray() != nil {
				currentArray().Add(nodeString)
			} else if currentParentheses() != nil {
				currentParentheses().Add(nodeString)
			} else {
				currentBlock().Add(nodeString)
			}
		}

		if token.Type() == lex.TOKEN_BOOLEAN {
			value, err := strconv.ParseBool(token.Literal())

			if err != nil {
				if currentArray() != nil {
					currentArray().Add(NewNodeBoolean(value))
				} else if currentParentheses() != nil {
					currentParentheses().Add(NewNodeBoolean(value))
				} else {
					currentBlock().Add(NewNodeBoolean(value))
				}
			}
		}

		if token.Type() == lex.TOKEN_LEFT_BRACE {
			if currentArray() != nil {
				result.AddError(NewParseError("Unexpected left brace"))
			} else if currentParentheses() != nil {
				result.AddError(NewParseError("Unexpected left brace"))
			} else {
				block := NewNodeBlock()
				blocks = append(blocks, block)
			}
		}

		if token.Type() == lex.TOKEN_RIGHT_BRACE {
			if currentArray() != nil {
				result.AddError(NewParseError("Unexpected right brace"))
			} else if currentParentheses() != nil {
				result.AddError(NewParseError("Unexpected right brace"))
			} else {
				lastBlock := currentBlock()
				blocks = blocks[:len(blocks)-1]
				currentBlock().Add(lastBlock)
			}
		}

		if token.Type() == lex.TOKEN_LEFT_BRACKET {
			array := NewNodeArray()

			/*
				if currentArray() != nil {
					currentArray().Add(array)
				} else {
					currentBlock().Add(array)
				}
			*/

			arrays = append(arrays, array)
		}

		if token.Type() == lex.TOKEN_RIGHT_BRACKET {
			if currentArray() != nil {
				currentArray().Add(currentArray())
			} else {
				currentBlock().Add(currentArray())
			}

			arrays = arrays[:len(arrays)-1]
		}

		if token.Type() == lex.TOKEN_LEFT_PARENTHESES {
			parentheses = append(parentheses, NewNodeParentheses())
		}

		if token.Type() == lex.TOKEN_RIGHT_PARENTHESES {
			if len(parentheses) == 0 {
				result.AddError(NewParseError("Unexpected right parentheses"))
			} else {
				currentBlock().Add(parentheses[len(parentheses)-1])
				parentheses = parentheses[:len(parentheses)-1]
			}
		}

		if token.Type() == lex.TOKEN_NEWLINE {
			currentBlock().Add(NewNodeLineEnd())
		}
	}

	return result
}

func PrintBlock(block Node, level int) {
	prefix := ""

	for i := 0; i < level; i++ {
		prefix += "  "
	}

	fmt.Println(prefix + "|_ " + block.ToString().Value() + " (" + block.Type() + ")")

	for _, element := range block.Children() {
		PrintBlock(element, level+1)
	}
}

func (p *Parser) PrintTree() {
	fmt.Println("ROOT")

	for _, node := range p.root.Children() {
		PrintBlock(node, 1)
	}
}

func (p *Parser) Root() Node {
	return p.root
}
