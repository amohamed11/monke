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

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.chr {
	// Handle known tokens
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
	default:
		if isLetter(l.chr) {
			// Handle identifiers
			tok.Literal = l.readIdenitifer()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.chr) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			// Token is not know and is therefore illegal
			tok = newToken(token.ILLEGAL, l.chr)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, chr byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(chr)}
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

func (l *Lexer) readIdenitifer() string {
	position := l.position

	for isLetter(l.chr) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.chr) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.chr == ' ' || l.chr == '\t' || l.chr == '\n' || l.chr == '\r' {
		l.readChar()
	}
}

func isLetter(chr byte) bool {
	return 'a' <= chr && chr <= 'z' || 'A' <= chr && chr <= 'Z' || chr == '_'
}

func isDigit(chr byte) bool {
	return '0' <= chr && chr <= '9'
}
