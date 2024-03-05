package ast

import (
	"lumen/token"
)

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (expression *ExpressionStatement) statementNode() {}

func (expression *ExpressionStatement) TokenLiteral() string {
	return expression.TokenLiteral()
}

func (expression *ExpressionStatement) String() string {
	if expression.Expression != nil {
		return expression.Expression.String()
	}

	return ""
}
