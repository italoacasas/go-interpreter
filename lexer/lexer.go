package lexer

import "gointerpreter/token"

// Lexer structure
type Lexer struct {
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position in input (after current char)
	ch byte // current char under examination
}

// New initialize the Lexer structure
func New(input  string) *Lexer {
	// Setup input
	l := &Lexer{input: input}
	// Setup position, readPosition and ch
	l.readChar()
	return l
}

// readChar makes sure the Lexer structure keeps
// a valid state
func (l *Lexer) readChar() {
	// If
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}