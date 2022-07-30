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

func New(lexer *lexer.Lexer) *Parser {
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
	return nil
}
