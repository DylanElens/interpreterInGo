package parser

import (
	"lumen/ast"
	"lumen/token"
)

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	//We construct a *ast.returnStatement node with the token the
	//parser is currently sitting on (hopefully a return token)
	statement := &ast.ReturnStatement{Token: parser.currentToken}

	parser.nextToken()

	//TODO: expressions
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement

}
