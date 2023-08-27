package lex

const (
	TOKEN_STRING TokenType = iota
	TOKEN_NUMBER
	TOKEN_BOOLEAN
	TOKEN_IDENTIFIER
	TOKEN_OPERATOR
	TOKEN_KEYWORD
	TOKEN_LEFT_PARENTHESES
	TOKEN_RIGHT_PARENTHESES
	TOKEN_LEFT_BRACE
	TOKEN_RIGHT_BRACE
	TOKEN_LEFT_BRACKET
	TOKEN_RIGHT_BRACKET
	TOKEN_SEMILICON
	TOKEN_NEWLINE
)

type TokenType int

type Token struct {
	tokenType TokenType
	literal   string
	intend    int
}

func NewToken(tokenType TokenType, literal string, intend int) Token {
	return Token{tokenType: tokenType, literal: literal, intend: intend}
}

func (t Token) Type() TokenType {
	return t.tokenType
}

func (t Token) Literal() string {
	return t.literal
}
