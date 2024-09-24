package parser

import (
	"testing"

	"github.com/flokkq/monkey-interpreter/lexer"
)

func TestParserErrors(t *testing.T) {
	tests := []struct {
		input         string
		expectedError string
	}{
		{
			input:         "let x 5;",
			expectedError: "expected next token to be =, got INT instead",
		},
		{
			input:         "let x",
			expectedError: "expected next token to be =, got EOF instead",
		},
		{
			input:         "let =",
			expectedError: "expected next token to be IDENT, got = instead",
		},
	}

	for idx, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		_ = p.ParseProgram()

		errors := p.Errors()
		if len(errors) == 0 {
			t.Errorf("Test case %d: Expected parser to have errors, but got none", idx)
			continue
		}

		if errors[0] != tt.expectedError {
			t.Errorf("Test case %d: Expected error message to be %q, got %q", idx, tt.expectedError, errors[0])
		}
	}
}
