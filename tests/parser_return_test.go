package tests

import (
	"lumen/ast"
	"lumen/lexer"
	"lumen/parser"
	"testing"
)

func TestReturnStatements(test *testing.T) {
	input := `
        return 5;
        return 10;
        return 993322;
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

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)

		if !ok {
			test.Errorf("statement not *ast.returnStatement. got=%T", statement)
			continue
		}

		if returnStatement.TokenLiteral() != "return" {
			test.Errorf(
				"returnStatement.TokenLiteral not 'return', got %q",
				statement.TokenLiteral(),
			)
		}
	}

	// tests := []struct {
	// 	expectedIndentifier string
	// }{
	// 	{"x"},
	// 	{"y"},
	// 	{"foobar"},
	// }

	// for i, tt := range tests {
	// 	statement := program.Statements[i]
	// 	if !testReturnStatement(test, statement, tt.expectedIndentifier) {
	// 		return
	// 	}
	// }
}
