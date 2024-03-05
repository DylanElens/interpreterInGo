package tests

import (
	"lumen/ast"
	"lumen/lexer"
	"lumen/parser"
	"testing"
)

func TestIdentifierExpression(test *testing.T) {
	input := `
        foobar;
    `

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	program := parser.ParseProgram()
	checkParseErrors(test, parser)

	if program == nil {
		test.Fatalf("ParsePogram() returned nil")
	}

	if len(program.Statements) != 1 {
		test.Fatalf(
			"program.Statements does not contain 1 statements. got=%d",
			len(program.Statements),
		)
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		test.Fatalf(
			"program.Statements[0] is not an ast.ExpressionStatement. got=%T",
			program.Statements[0],
		)
	}

	ident, ok := statement.Expression.(*ast.Identifier)

	if !ok {
		test.Fatalf("exp not *ast.Identifier. got=%T", statement.Expression)
	}

	if !ok {
		test.Fatalf("ident.Value not %s. got=%T", "foobar", statement.Expression)
	}

	if ident.TokenLiteral() != "foobar" {
		test.Fatalf("ident.TokenLiteral not %s. got=%T", "foobar", statement.TokenLiteral())
	}

}
