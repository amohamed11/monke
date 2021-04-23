package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // Identifiers: x, y, foo
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "PLUS"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = ")"
	RPAREN = "("
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
