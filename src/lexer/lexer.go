package lexer

import "by.far.the.best.interpreter.ever.made.io/src/token"

const (
	NUL = 0
)

type Lexer struct {
	input            string
	currentPosition  int
	nextPositionRead int
	ch               byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) getNotEqToken() token.Token {
	ch := l.ch
	l.readChar()
	literal := string(ch) + string(l.ch)
	return token.Token{Type: token.NOT_EQ, Literal: literal}
}

func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok = l.getNotEqToken()
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case NUL:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		unident := l.readUnidentifiedChar()
		if unident != nil {
			return *unident
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return tok
}

func (l *Lexer) readUnidentifiedChar() *token.Token {
	tok := &token.Token{}
	if l.isCurrentCharLetter() {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	} else if l.isCurrentCharDigit() {
		tok.Type = token.INT
		tok.Literal = l.readNumber()
		return tok
	}
	return nil
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) isEnd() bool {
	return l.nextPositionRead >= len(l.input)
}

func (l *Lexer) readChar() {
	if l.isEnd() {
		l.ch = NUL
	} else {
		l.ch = l.input[l.nextPositionRead]
	}
	l.currentPosition = l.nextPositionRead
	l.nextPositionRead += 1
}

func (l *Lexer) peekChar() byte {
	if l.nextPositionRead >= len(l.input) {
		return NUL
	} else {
		return l.input[l.nextPositionRead]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for l.isCurrentCharLetter() {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for l.isCurrentCharDigit() {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) isCurrentCharLetter() bool {
	return 'a' <= l.ch && l.ch <= 'z' || 'A' <= l.ch && l.ch <= 'Z' || l.ch == '_'
}

func (l *Lexer) isCurrentCharDigit() bool {
	return '0' <= l.ch && l.ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
