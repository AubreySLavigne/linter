package main

import (
	"testing"
)

func TestParseTokens(t *testing.T) {

	tests := []struct {
		input   string
		expects []token
	}{
		{input: "input {", expects: []token{
			{Type: identifier, Value: "input", Line: 0, Column: 0, Index: 0},
			{Type: separator, Value: "{", Line: 0, Column: 6, Index: 0},
		}},
		{input: `input {
    beats {
        port => "5044"
    }
}`, expects: []token{
			{Type: identifier, Value: "input", Line: 0, Column: 0, Index: 0},
			{Type: separator, Value: "{", Line: 0, Column: 6, Index: 6},
			{Type: identifier, Value: "beats", Line: 1, Column: 4, Index: 12},
			{Type: separator, Value: "{", Line: 1, Column: 10, Index: 18},
			{Type: identifier, Value: "port", Line: 2, Column: 8, Index: 28},
			{Type: operator, Value: "=>", Line: 2, Column: 13, Index: 33},
			{Type: literal, Value: "\"5044\"", Line: 2, Column: 16, Index: 36},
			{Type: separator, Value: "}", Line: 3, Column: 4, Index: 47},
			{Type: separator, Value: "}", Line: 4, Column: 0, Index: 49},
			{},
		}},
		{input: "1234", expects: []token{}},
	}

	tz := tokenizer{}
	tz.addRule(rule{name: identifier, pattern: "[A-Za-z]+"})
	tz.addRule(rule{name: separator, pattern: "[{}]"})
	tz.addRule(rule{name: literal, pattern: `"[^"]*"`})
	tz.addRule(rule{name: operator, pattern: "=>"})

	for _, test := range tests {
		// Setup Tokenizer
		tz.source = test.input
		tz.reset()

		// Run test
		for index, expectedToken := range test.expects {
			res, err := tz.next()
			if err != nil {
				t.Errorf("Unexpected Error during test %v [Index: %d]: %s", expectedToken, index, err)
			} else if !res.equal(expectedToken) {
				t.Errorf("Parsed Token does not match expected for input '%s' [Index: %d]. Got %v, expected %v", tz.source, index, res, expectedToken)
			}
		}
	}
}

func TestInvalidRegexpRule(t *testing.T) {
	tz := tokenizer{}
	tz.addRule(rule{name: operator, pattern: `\`})

	tz.source = "tokenize"

	_, err := tz.next()
	if err == nil {
		t.Errorf("Checking Invalid Regexp Rule. Regexp is not failing as expected")
	}
}
