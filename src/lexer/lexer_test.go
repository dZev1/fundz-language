package lexer

import (
	"testing"

	"github.com/dZev1/fundz-language/token"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}
	}{
		{
			name : "Test filter function lexing",
			input : `
				filter : (T -> bool) -> []T -> []T
				filter(_, []) { []. }
				filter(p, x::xs) {
					let rec : []T = filter(p, xs);
					if (p(x)) then { x :: rec. } else { rec. }
				}
			`,
			expected: []struct {
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
				{token.LPAREN, "("},
				{token.IDENT, "p"},
				{token.LPAREN, "("},
				{token.IDENT, "x"},
				{token.RPAREN, ")"},
				{token.RPAREN, ")"},
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
			},
		},
		{
			name : "Test reading every token type",
			input : `This is a test! 12345 + - * / < <= > >= == != && || ^| <- -> => , . ; : :: ( ) { } [ ] _`,
			expected: []struct {
				expectedType    token.TokenType
				expectedLiteral string
			}{
				{token.IDENT, "This"},
				{token.IDENT, "is"},
				{token.IDENT, "a"},
				{token.IDENT, "test"},
				{token.BANG, "!"},
				{token.INT, "12345"},
				{token.PLUS, "+"},
				{token.MINUS, "-"},
				{token.ASTERISK, "*"},
				{token.SLASH, "/"},
				{token.LT, "<"},
				{token.LEQ, "<="},
				{token.GT, ">"},
				{token.GEQ, ">="},
				{token.EQ, "=="},
				{token.NEQ, "!="},
				{token.AND, "&&"},
				{token.OR, "||"},
				{token.XOR, "^|"},
				{token.LARROW, "<-"},
				{token.RARROW, "->"},
				{token.DRARROW, "=>"},
				{token.COMMA, ","},
				{token.DOT, "."},
				{token.SEMICOLON, ";"},
				{token.COLON, ":"},
				{token.DOUBLE_COLON, "::"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.LBRACKET, "["},
				{token.RBRACKET, "]"},
				{token.UNDERSCORE, "_"},
				{token.EOF, ""},
			},
		},
	}

	
	for _, tt := range tests {
		l := New(tt.input)
		t.Run(tt.name, func(t *testing.T) {
			for i, expectedToken := range tt.expected {
				tok := l.NextToken()

				if tok.Type != expectedToken.expectedType {
					t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
						i, expectedToken.expectedType, tok.Type)
				}

				if tok.Literal != expectedToken.expectedLiteral {
					t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
						i, expectedToken.expectedLiteral, tok.Literal)
				}
			}
		})
	}
}
