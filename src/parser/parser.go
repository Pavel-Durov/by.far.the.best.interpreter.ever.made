package parser

import (
	"by.far.the.best.interpreter.ever.made.io/src/ast"
	"by.far.the.best.interpreter.ever.made.io/src/lexer"
	"by.far.the.best.interpreter.ever.made.io/src/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	}
	return nil
}

func (p *Parser) isEOF() bool {
	return p.currentToken.Type == token.EOF
}
func NewParser(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() *ast.Program {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
	return nil
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := ast.NewProgram()
	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()

	}
	return prog
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := ast.LetStatement{Token: p.currentToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: skipping till semicolon for now
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return &stmt
}
