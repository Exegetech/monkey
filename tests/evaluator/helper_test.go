package evaluator_test

import (
	"testing"

	"github.com/Exegetech/monkey/src/evaluator"
	"github.com/Exegetech/monkey/src/lexer"
	"github.com/Exegetech/monkey/src/object"
	"github.com/Exegetech/monkey/src/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)

	program := p.ParseProgram()

	return evaluator.Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf(
			"object is not Integer. got = %T (%+v)",
			obj,
			obj,
		)

		return false
	}

	if result.Value != expected {
		t.Errorf(
			"object has wrong value. got = %d, want = %d",
			result.Value,
			expected,
		)

		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf(
			"object is not Boolean. got = %T (%+v)",
			obj,
			obj,
		)

		return false
	}

	if result.Value != expected {
		t.Errorf(
			"object has wrong value. got = %t, want = %t",
			result.Value,
			expected,
		)

		return false
	}

	return true
}
