package parser

import (
	"github.com/Exegetech/monkey/src/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // < or =< or > or >=
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x or !x
	CALL        // myFunction(x)
)

var precedence = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.LTE:      LESSGREATER,
	token.GT:       LESSGREATER,
	token.GTE:      LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}
