package ast

import (
	"github.com/Exegetech/monkey/src/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

// TODO: Identifier doesn't produce a value
// so it shouldn't be an expression
// this is here to make things simple. Refactor
func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
