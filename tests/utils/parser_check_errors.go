package tests

import (
	"lumen/parser"
	"testing"
)

func CheckParseErrors(test *testing.T, parser *parser.Parser) {
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
