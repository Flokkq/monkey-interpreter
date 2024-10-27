package evaluator

import (
	"testing"

	"github.com/flokkq/monkey-interpreter/lexer"
	"github.com/flokkq/monkey-interpreter/object"
	"github.com/flokkq/monkey-interpreter/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testEval(s string) object.Object {
	l := lexer.New(s)
	p := parser.New(l)

	return Eval(p.ParseProgram())
}
