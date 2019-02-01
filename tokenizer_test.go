package main

import (
	"testing"
)

func TestParseTokens(t *testing.T) {

	tests := []struct {
		input   string
		expects token
	}{
		{input: "input", expects: token{Type: tokenType("Identifier"), Value: "input", Line: 1, Column: 1}},
		{input: "{", expects: token{Type: tokenType("OpenCurly"), Value: "{", Line: 1, Column: 1}},
		{input: "1234", expects: token{}},
	}

	tz := tokenizer{}
	tz.addRule(rule{name: "Identifier", pattern: "[A-Za-z]+"})
	tz.addRule(rule{name: "OpenCurly", pattern: "{"})

	for _, test := range tests {
		// Setup Tokenizer
		tz.source = test.input
		tz.reset()

		// Run test
		res, err := tz.next()
		if err != nil {
			t.Errorf("Unexpected Error during test %v: %s", test, err)
		} else if !res.equal(test.expects) {
			t.Errorf("Parsed Token does not match expected for input '%s'. Got %v, expected %v", tz.source, res, test.expects)
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
