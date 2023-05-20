package ast

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode() // dummy method for type checking
}

type Expression interface {
	Node
	expressionNode() // dummy method for type checking
}
