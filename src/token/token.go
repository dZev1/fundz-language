package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Special Tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	RARROW   = "->"
	LARROW   = "<-"
	DRARROW  = "=>"
	EXP      = "^"

	LT  = "<"
	LEQ = "<="
	GT  = ">"
	GEQ = ">="

	EQ  = "=="
	NEQ = "!="
	AND = "&&"
	OR  = "||"
	XOR = "^|"

	// Delimiters
	PIPE         = "|"
	COMMA        = ","
	DOT          = "."
	SEMICOLON    = ";"
	COLON        = ":"
	DOUBLE_COLON = "::"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	UNDERSCORE = "_"
	LET        = "LET"
	TRUE       = "TRUE"
	FALSE      = "FALSE"
	IF         = "IF"
	THEN       = "THEN"
	ELSE       = "ELSE"
	BOOL_TYPE  = "BOOL_TYPE"
	INT_TYPE   = "INT_TYPE"
	CHAR_TYPE  = "CHAR_TYPE"
	FLOAT_TYPE = "FLOAT_TYPE"
)

var keywords = map[string]TokenType{
	"let":   LET,
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"then":  THEN,
	"else":  ELSE,
	"bool":  BOOL_TYPE,
	"int":   INT_TYPE,
	"char":  CHAR_TYPE,
	"float": FLOAT_TYPE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
