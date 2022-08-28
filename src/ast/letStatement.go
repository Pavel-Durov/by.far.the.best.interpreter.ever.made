package ast

import "by.far.the.best.interpreter.ever.made.io/src/token"

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value *Expression
}

func (s *LetStatement) statementNode()       {}
func (s *LetStatement) TokenLiteral() string { return s.Token.Literal }
