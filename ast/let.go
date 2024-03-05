package ast

import "lumen/token"

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {
	//TODO
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
