package lexer

import (
	"testing"

	"github.com/dZev1/fundz-language/token"
)

func TestNextToken(t *testing.T) {
	input := `filter : (T -> bool) -> []T -> []T
filter(_, []) { []. }
filter(p, x::xs) {
	let rec : []T = filter(p, xs);
	if { p(x) } then { x :: rec. } else { rec. } }
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "filter"},
		{token.COLON, ":"},
		{token.LPAREN, "("},
		{token.IDENT, "T"},
		{token.RARROW, "->"},
		{token.BOOL_TYPE, "bool"},
		{token.RPAREN, ")"},
		{token.RARROW, "->"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.IDENT, "T"},
		{token.RARROW, "->"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.IDENT, "T"},
		{token.IDENT, "filter"},
		{token.LPAREN, "("},
		{token.UNDERSCORE, "_"},
		{token.COMMA, ","},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.DOT, "."},
		{token.RBRACE, "}"},
		{token.IDENT, "filter"},
		{token.LPAREN, "("},
		{token.IDENT, "p"},
		{token.COMMA, ","},
		{token.IDENT, "x"},
		{token.DOUBLE_COLON, "::"},
		{token.IDENT, "xs"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.LET, "let"},
		{token.IDENT, "rec"},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.IDENT, "T"},
		{token.ASSIGN, "="},
		{token.IDENT, "filter"},
		{token.LPAREN, "("},
		{token.IDENT, "p"},
		{token.COMMA, ","},
		{token.IDENT, "xs"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LBRACE, "{"},
		{token.IDENT, "p"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.THEN, "then"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.DOUBLE_COLON, "::"},
		{token.IDENT, "rec"},
		{token.DOT, "."},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.IDENT, "rec"},
		{token.DOT, "."},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
