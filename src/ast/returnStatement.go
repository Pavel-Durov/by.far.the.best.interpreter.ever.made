package ast

import "by.far.the.best.interpreter.ever.made.io/src/token"

type ReturnStatement struct {
	Token token.Token // the 'return' token ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
