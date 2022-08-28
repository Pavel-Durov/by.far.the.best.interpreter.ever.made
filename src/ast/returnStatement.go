package ast

import (
	"bytes"

	"by.far.the.best.interpreter.ever.made.io/src/token"
)

type ReturnStatement struct {
	Token       token.Token // the 'return' token ReturnValue Expression
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}
