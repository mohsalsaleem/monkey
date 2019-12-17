package lexer

import "github.com/mohsalsaleem/monkey/token"

// Lexer - The lexer
type Lexer struct {
	input        string // The input string
	position     int    // current position in input (points to current ch)
	readPosition int    // current reading position in input (points to next character of ch)
	ch           byte   // current character under examination
}

// New - Returns a new Lexer pointer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// If readPosition is greater than the length of the string,
	// ch will be 0 else, ch will be the current position that's being read
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken - Get next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// Create a token cased on the character recieved
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPARAN, l.ch)
	case ')':
		tok = newToken(token.RPARAN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}

	// Read next char to memory
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	// Get the current position of ch
	position := l.position
	// Loop until we encounter a 0,
	// meaning that it's not a valid character,
	// which meands there's a space or eof
	// In this operation, the current position is also update
	// via readChar()
	for isLetter(l.ch) {
		l.readChar()
	}

	// Return the substring from where we started,
	// till before encountered a 0
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	// If we encounter any whitespace, move forward a character
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
