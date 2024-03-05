package parser

import (
	"fmt"
	"lumen/token"
)

// This is an "assertion" function. They enforce the correctness of the order
// of toekns by checking the type of the next token.
func (parser *Parser) expectPeek(inputToken token.TokenType) bool {
	//We check if the token our parser is on is equal to the inputToken
	// we are expecting
	if parser.peekTokenIs(inputToken) {
		//Only if the token matches we will move on
		parser.nextToken()
		return true
	} else {
		parser.peekError(inputToken)
		return false
	}
}

func (parser *Parser) currentTokenIs(inputToken token.TokenType) bool {
	return parser.currentToken.Type == inputToken
}

func (parser *Parser) peekTokenIs(inputToken token.TokenType) bool {
	return parser.peekToken.Type == inputToken
}

func (parser *Parser) peekError(inputToken token.TokenType) {
	message := fmt.Sprintf(
		"expected next token to be %s, got %s instead",
		inputToken,
		parser.peekToken.Type,
	)
	parser.errors = append(parser.errors, message)
}
