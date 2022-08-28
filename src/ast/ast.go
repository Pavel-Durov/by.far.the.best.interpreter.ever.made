package ast

// Every node in AST has to implement this interface
type Node interface {
	// returns associated token literal value (used only for debugging and testing)
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}
