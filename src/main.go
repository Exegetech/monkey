package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// type Lexer struct {
//   input string

// }

// // Token represents a token in the input
// type Token struct {
// 	Type  string
// 	Value string
// }

// // Lexer tokenizes the input string
// type Lexer struct {
// 	input  string
// 	tokens []Token
// 	pos    int
// }

// // NewLexer creates a new Lexer instance
// func NewLexer(input string) *Lexer {
// 	return &Lexer{
// 		input:  input,
// 		tokens: make([]Token, 0),
// 		pos:    0,
// 	}
// }

// // Tokenize processes the input string and generates tokens
// func (l *Lexer) Tokenize() []Token {
// 	for l.pos < len(l.input) {
// 		switch l.input[l.pos] {
// 		case '+':
// 			l.tokens = append(l.tokens, Token{"PLUS", "+"})
// 		case '-':
// 			l.tokens = append(l.tokens, Token{"MINUS", "-"})
// 		default:
// 			if isDigit(l.input[l.pos]) {
// 				value := l.parseNumber()
// 				l.tokens = append(l.tokens, Token{"NUMBER", value})
// 			} else if !isWhitespace(l.input[l.pos]) {
// 				panic("Invalid character: " + string(l.input[l.pos]))
// 			}
// 		}
// 		l.pos++
// 	}
// 	return l.tokens
// }

// // parseNumber extracts the numeric value from the input string
// func (l *Lexer) parseNumber() string {
// 	start := l.pos
// 	for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
// 		l.pos++
// 	}
// 	return l.input[start:l.pos]
// }

// // isDigit checks if the given character is a digit
// func isDigit(ch byte) bool {
// 	return '0' <= ch && ch <= '9'
// }

// // isWhitespace checks if the given character is a whitespace
// func isWhitespace(ch byte) bool {
// 	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
// }

// // Interpreter interprets the input tokens and evaluates expressions
// type Interpreter struct {
// 	tokens []Token
// 	pos    int
// }

// // NewInterpreter creates a new Interpreter instance
// func NewInterpreter(tokens []Token) *Interpreter {
// 	return &Interpreter{
// 		tokens: tokens,
// 		pos:    0,
// 	}
// }

// // Expr represents an arithmetic expression
// type Expr struct {
// 	Op    string
// 	Left  *Expr
// 	Right *Expr
// 	Value int
// }

// // ParseExpression parses an arithmetic expression
// func (i *Interpreter) ParseExpression() *Expr {
// 	token := i.tokens[i.pos]
// 	i.pos++

// 	left := &Expr{
// 		Op:    token.Type,
// 		Value: atoi(token.Value),
// 	}

// 	for i.pos < len(i.tokens) && (i.tokens[i.pos].Type == "PLUS" || i.tokens[i.pos].Type == "MINUS") {
// 		token = i.tokens[i.pos]
// 		i.pos++

// 		right := &Expr{
// 			Op:    token.Type,
// 			Value: atoi(i.tokens[i.pos].Value),
// 		}

// 		left = &Expr{
// 			Op:    token.Type,
// 			Left:  left,
// 			Right: right,
// 		}
// 		i.pos++
// 	}

// 	return left
// }

// // Evaluate evaluates the given arithmetic expression
// func (e *Expr) Evaluate() int {
// 	if e.Op == "NUMBER" {
// 		return e.Value
// 	}

// 	leftValue := e.Left.Evaluate()
// 	rightValue := e.Right.Evaluate()

// 	switch e.Op {
// 	case "PLUS":
// 		return
