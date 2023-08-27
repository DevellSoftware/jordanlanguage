package lex

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	LEXER_STATE_DEFAULT LexerState = iota
	LEXER_STATE_STRING
)

type LexerState int

type Lexer struct {
	input             string
	position          int
	state             LexerState
	openedParentheses int
	openedBraces      int
	openedBrackets    int
	currentIntended   int
	tokens            []Token
}

var operators = []string{
	"=",
	"+",
	",",
	";",
	"==",
	"!=",
	"<",
	">",
	"!",
	"-",
	"*",
	"/",
}

var keywords = []string{
	"fun",
	"let",
	"true",
	"false",
	"if",
	"else",
	"return",
}

var regexpInteger = regexp.MustCompile(`^[0-9]+$`)
var regexpFloat = regexp.MustCompile(`^[0-9]+\.[0-9]+$`)
var regexpIdentifier = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

func filter(input string) string {
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "(", " ( ")
	input = strings.ReplaceAll(input, ")", " ) ")
	input = strings.ReplaceAll(input, "{", " { ")
	input = strings.ReplaceAll(input, "}", " } ")
	input = strings.ReplaceAll(input, "[", " [ ")
	input = strings.ReplaceAll(input, "]", " ] ")
	input = strings.ReplaceAll(input, ",", " , ")
	input = strings.ReplaceAll(input, ";", " ; ")

	return input
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:             filter(input),
		position:          0,
		state:             LEXER_STATE_DEFAULT,
		openedParentheses: 0,
		openedBraces:      0,
		openedBrackets:    0,
		currentIntended:   0,
		tokens:            make([]Token, 0),
	}
}

func (l *Lexer) isOperator(word string) bool {
	for _, operator := range operators {
		if operator == word {
			return true
		}
	}

	return false
}

func (l *Lexer) isKeyword(word string) bool {
	for _, keyword := range keywords {
		if keyword == word {
			return true
		}
	}

	return false
}

func (l *Lexer) isNumber(word string) bool {
	return regexpInteger.MatchString(word) || regexpFloat.MatchString(word)
}

func (l *Lexer) isIdentifier(word string) bool {
	return regexpIdentifier.MatchString(word)
}

func (l *Lexer) recognizeWord(word string) {
	word = strings.Trim(word, " ")

	if word == "" {
		return
	}

	if l.isOperator(word) {
		l.tokens = append(l.tokens, NewToken(TOKEN_OPERATOR, word, l.currentIntended))
		return
	}

	if l.isKeyword(word) {
		l.tokens = append(l.tokens, NewToken(TOKEN_KEYWORD, word, l.currentIntended))
		return
	}

	if l.isIdentifier(word) {
		l.tokens = append(l.tokens, NewToken(TOKEN_IDENTIFIER, word, l.currentIntended))
	}

	if l.isNumber(word) {
		l.tokens = append(l.tokens, NewToken(TOKEN_NUMBER, word, l.currentIntended))
	}
}

func (l *Lexer) Lex() {
	l.position = 0
	currentWord := ""
	currentString := ""

	var currentRune rune

	newLine := false

	for {
		if l.position >= len(l.input) {
			break
		}

		currentRune = rune(l.input[l.position])

		if l.state == LEXER_STATE_DEFAULT {
			if currentRune == '\t' && newLine {
				l.currentIntended += 2
				l.position++
				continue
			}

			if currentRune == '\n' || currentRune == '\r' {
				l.recognizeWord(currentWord)
				l.tokens = append(l.tokens, NewToken(TOKEN_NEWLINE, "\n", l.currentIntended))
				currentWord = ""
				newLine = true
				l.currentIntended = 0
				l.position++
				continue
			}

			if currentRune == ' ' {
				if newLine {
					l.currentIntended++
					l.position++
					continue
				}

				l.recognizeWord(currentWord)
				currentWord = ""
				l.position++
				continue
			}

			if currentRune == '"' {
				l.state = LEXER_STATE_STRING
				l.position++
				continue
			}

			if currentRune == '(' {
				l.openedParentheses++
				l.tokens = append(l.tokens, NewToken(TOKEN_LEFT_PARENTHESES, "(", l.currentIntended))
				l.position++
				continue
			}

			if currentRune == ')' {
				l.openedParentheses--
				l.tokens = append(l.tokens, NewToken(TOKEN_RIGHT_PARENTHESES, ")", l.currentIntended))
				l.position++
				continue
			}

			if currentRune == '{' {
				l.openedBraces++
				l.tokens = append(l.tokens, NewToken(TOKEN_LEFT_BRACE, "{", l.currentIntended))
				l.position++
				continue
			}

			if currentRune == '}' {
				fmt.Println("HERE")
				l.openedBraces--
				l.tokens = append(l.tokens, NewToken(TOKEN_RIGHT_BRACE, "}", l.currentIntended))
				l.position++
				continue
			}

			if currentRune == '[' {
				l.openedBrackets++
				l.tokens = append(l.tokens, NewToken(TOKEN_LEFT_BRACKET, "[", l.currentIntended))
				l.position++
				continue
			}

			if currentRune == ']' {
				l.openedBrackets--
				l.tokens = append(l.tokens, NewToken(TOKEN_RIGHT_BRACKET, "]", l.currentIntended))
				l.position++
				continue
			}

			if currentRune == ';' {
				l.tokens = append(l.tokens, NewToken(TOKEN_SEMILICON, ";", l.currentIntended))
				l.position++
				continue
			}

			if currentRune != ' ' && currentRune != '\t' && newLine {
				newLine = false
				continue
			}

			currentWord += string(currentRune)
		} else if l.state == LEXER_STATE_STRING {
			if currentRune == '"' {
				l.state = LEXER_STATE_DEFAULT

				l.tokens = append(l.tokens, NewToken(TOKEN_STRING, currentString, l.currentIntended))
			} else {
				currentString += string(currentRune)
			}
		}

		l.position++

		if l.position == len(l.input) {
			l.recognizeWord(currentWord)
		}
	}
}

func (l *Lexer) Tokens() []Token {
	return l.tokens
}
