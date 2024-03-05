package tests

import (
	"lumen/ast"
	"lumen/lexer"
	"lumen/parser"
	"testing"
)

func TestIntegerLiteralExpression(test *testing.T) {
	input := `
        5;
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

	literal, ok := statement.Expression.(*ast.IntegerLiteral)

	if !ok {
		test.Fatalf("exp not *ast.Identifier. got=%T", statement.Expression)
	}

	if literal.Value != 5 {
		test.Fatalf("literal.Value not %d. got=%d", "5", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		test.Fatalf("literal.TokenLiteral not %s. got=%s", "5", statement.TokenLiteral())
	}

}
