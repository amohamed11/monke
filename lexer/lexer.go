package lexer

import "monke/token"

type Lexer struct {
	input        string
	position     int  // cursor for input (current char)
	readPosition int  // cursor for next to be read (after current char)
	chr          byte // current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.chr = 0
	} else {
		l.chr = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, chr byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(chr)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.chr {
	case '=':
		tok = newToken(token.ASSIGN, l.chr)
	case ';':
		tok = newToken(token.SEMICOLON, l.chr)
	case '(':
		tok = newToken(token.LPAREN, l.chr)
	case ')':
		tok = newToken(token.RPAREN, l.chr)
	case '{':
		tok = newToken(token.LBRACE, l.chr)
	case '}':
		tok = newToken(token.RBRACE, l.chr)
	case ',':
		tok = newToken(token.COMMA, l.chr)
	case '+':
		tok = newToken(token.PLUS, l.chr)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}
