package parser_test

import (
	"testing"

	"by.far.the.best.interpreter.ever.made.io/src/ast"
	"by.far.the.best.interpreter.ever.made.io/src/lexer"
	"by.far.the.best.interpreter.ever.made.io/src/parser"
	"github.com/stretchr/testify/assert"
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
		let kimchi = 198;
	`
	lex := lexer.NewLexer(input)
	p := parser.NewParser(lex)
	prog := p.ParseProgram()

	checkParserErrors(t, p)
	if prog == nil {
		t.Fatalf("Program is nil")
	}
	if (len(prog.Statements)) != 3 {
		t.Fatalf("Program statements do not contain 3 statements")
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

func TestReturnStatements(t *testing.T) {
	input := `
   return 5;
   return 10;
   return 993322;
   `
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestLetStatementsErrors(t *testing.T) {
	input := `
		let = 198;
	`
	lex := lexer.NewLexer(input)
	p := parser.NewParser(lex)
	p.ParseProgram()

	err := p.Errors()
	assert.Equal(t, 1, len(p.Errors()))
	assert.Equal(t, "expected next token to be IDENT, got = instead", err[0])
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
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
