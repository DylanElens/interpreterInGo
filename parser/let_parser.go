package parser

import (
	"lumen/ast"
	"lumen/token"
)

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	//We construct a *ast.LetStatement node with the token the
	//parser is currently sitting on (hopefully a let token)
	statement := &ast.LetStatement{Token: parser.currentToken}

	//Since a let statement looks like "let foo = 8;"
	//the next token should be an identifier thus "foo"
	// otherwise it is not a let statement
	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	//We can now set the statement name
	//a statement name is a combination of the token type
	//And the string value "foo"
	statement.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}

	//The token after the "foo" should be an assign
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	//TODO: Expressions!
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}
