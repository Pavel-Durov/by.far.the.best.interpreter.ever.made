package ast

type Program struct {
	Statements []Statement
}

func NewProgram() *Program {
	return &Program{
		Statements: []Statement{},
	}
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
