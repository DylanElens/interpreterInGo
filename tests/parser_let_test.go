package tests

import (
	"lumen/ast"
	"lumen/lexer"
	"lumen/parser"
	"testing"
)

func TestLetStatements(test *testing.T) {
	input := `
        let x = 5;
        let y = 10;
        let foobar = 838383;
    `

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	program := parser.ParseProgram()
	checkParseErrors(test, parser)

	if program == nil {
		test.Fatalf("ParsePogram() returned nil")
	}

	if len(program.Statements) != 3 {
		test.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	tests := []struct {
		expectedIndentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(test, statement, tt.expectedIndentifier) {
			return
		}
	}
}

func checkParseErrors(test *testing.T, parser *parser.Parser) {
	errors := parser.Errors()

	if len(errors) == 0 {
		return
	}

	test.Errorf("parser has %d errors", len(errors))

	for _, message := range errors {
		test.Errorf("parser error: %q", message)
	}

	test.FailNow()
}

func testLetStatement(test *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		test.Errorf("statement.Tokenliteral() not 'let', got=%q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)

	if !ok {
		test.Errorf("s not *ast.LetStatement. got =%T", statement)
		return false
	}

	if letStatement.Name.Value != name {
		test.Errorf("LetStatement.name.Value not '%s'. got =%s", name, letStatement.Name.Value)
	}

	if letStatement.Name.TokenLiteral() != name {
		test.Errorf("LetStatement.name.Value not '%s'. got =%s", name, letStatement.Name.Value)
	}

	return true
}
