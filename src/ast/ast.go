package ast

import (
	"github.com/Exegetech/monkey/src/token"
)

type Node interface {
  TokenLiteral() string
}

type Statement interface {
  Node
  statementNode() // dummy method for type checking
}

type Expression interface {
  Node
  expressionNode() // dummy method for type checking
}

type Program struct {
  Statements []Statement
}

func (p *Program) TokenLiteral() string {
  if len(p.Statements) > 0 {
    return p.Statements[0].TokenLiteral()
  }

  return ""
}

type LetStatement struct {
  Token token.Token
  Name *Identifier
  Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
  return ls.Token.Literal
}

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
