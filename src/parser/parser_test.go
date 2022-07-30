package parser_test

import (
	"testing"

	"by.far.the.best.interpreter.ever.made.io/src/ast"
	"by.far.the.best.interpreter.ever.made.io/src/lexer"
	"by.far.the.best.interpreter.ever.made.io/src/parser"
)

type TestCase struct {
	input              string
	expectedIdentifier string
	expectedValue      interface{}
}

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let kimchi = 1985;
	`
	lex := lexer.NewLexer(input)
	p := parser.NewParser(lex)
	prog := p.ParseProgram()
	if prog == nil {
		t.Fatalf("Program is nil")
	}
	if (len(prog.Statements)) != 3 {
		t.Fatalf("Program statements do nor contain 3 statements")
	}
	test := []struct {
		identifier string
	}{
		{"x"}, {"y"}, {"kimchi"},
	}

	for i, testCase := range test {
		stmt := prog.Statements[i]
		if !testLetStatement(t, stmt, testCase.identifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, identifier string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("token literal is not let 'let'. got=%q", stmt.TokenLiteral())
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("unexpected type. expected: *ast.LetStatement, got=%T", stmt)
		return false
	}
	if letStmt.Name.Value != identifier {
		t.Errorf("unexpected let statement name value not '%s'. got=%s", identifier, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != identifier {
		t.Errorf("unexpected let statement token literal not '%s'. got=%s", identifier, letStmt.Name)
		return false
	}
	return true
}
