package parser

import (
	"lumen/ast"
	"lumen/token"
)

type prefixParseFn func() ast.Expression
type infixParseFn func(ast.Expression) ast.Expression

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: parser.currentToken}

	statement.Expression = parser.parseExpression(LOWEST)

	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) parseExpression(precendence int) ast.Expression {
	prefix := parser.prefixParseFns[parser.currentToken.Type]

	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
}
