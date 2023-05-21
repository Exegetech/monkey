package parser_test

import (
	"testing"

	"github.com/Exegetech/monkey/src/ast"
	"github.com/Exegetech/monkey/src/lexer"
	"github.com/Exegetech/monkey/src/parser"
)

func TestLetStatement(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain 1 statements. got = %d",
				len(program.Statements),
			)
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	tests := []struct {
		input         string
		expectedValue interface{}
	}{
		{"return 5;", 5},
		{"return true;", true},
		{"return foobar;", "foobar"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain 1 statements. got = %d",
				len(program.Statements),
			)
		}

		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf(
				"stmt not *ast.ReturnStatement, got = %T",
				stmt,
			)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf(
				"returnStmt.TokenLiteral() not 'return' got = %q",
				returnStmt.TokenLiteral(),
			)
		}
	}
}
