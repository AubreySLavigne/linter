package main

import (
	"testing"
)

func TestParseTokens(t *testing.T) {

	tests := []struct {
		input   string
		expects []token
	}{
		{input: "input", expects: []token{{Type: tokenType("Identifier"), Value: "input", Line: 0, Column: 0}}},
		{input: "{", expects: []token{{Type: tokenType("OpenCurly"), Value: "{", Line: 0, Column: 0}, {}}},
		{input: "input {", expects: []token{
			{Type: tokenType("Identifier"), Value: "input", Line: 0, Column: 0},
			{Type: tokenType("OpenCurly"), Value: "{", Line: 0, Column: 6},
		}},
		{input: "1234", expects: []token{}},
	}

	tz := tokenizer{}
	tz.addRule(rule{name: "Identifier", pattern: "[A-Za-z]+"})
	tz.addRule(rule{name: "OpenCurly", pattern: "{"})

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
	tz.addRule(rule{name: "Invalid Rule", pattern: `\`})

	tz.source = "tokenize"

	_, err := tz.next()
	if err == nil {
		t.Errorf("Checking Invalid Regexp Rule. Regexp is not failing as expected")
	}
}
