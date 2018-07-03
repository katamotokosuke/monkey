package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const(
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "ASSIGN"
	PLUS = "PLUS"

	COMMA = ","
	SEMICOLON = ";"
	BANG = "!"
	MINUS = "-"
	SLASH = "/"
	ASTERISK = "*"
	LT = "<"
	GT = ">"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
	IF = "IF"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
	ELSE = "ELSE"
	EQ = "EQ"
	NOT_EQ = "NOT_EQ"
)

var keywords = map[string] TokenType {
	"fn": FUNCTION,
	"let": LET,
	"if": IF,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
	"else": ELSE,
	"==": EQ,
	"!=": NOT_EQ,
}

func LookupIdent(ident string) TokenType {
	if tok,  ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}