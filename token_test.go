package main

import (
	"testing"
)

func TestTokenEqual(t *testing.T) {
	tests := []struct {
		token1  token
		token2  token
		expects bool
	}{
		{
			token1:  token{Type: identifier, Value: "test", Line: 15, Column: 23, Index: 205},
			token2:  token{Type: identifier, Value: "test", Line: 15, Column: 23, Index: 205},
			expects: true,
		},
		{
			token1:  token{Type: identifier, Value: "test", Line: 15, Column: 23, Index: 205},
			token2:  token{Type: literal, Value: "test", Line: 15, Column: 23, Index: 205},
			expects: false,
		},
		{
			token1:  token{Type: identifier, Value: "test", Line: 15, Column: 23, Index: 205},
			token2:  token{Type: identifier, Value: "different", Line: 15, Column: 23, Index: 205},
			expects: false,
		},
		{
			token1:  token{Type: identifier, Value: "test", Line: 15, Column: 23, Index: 205},
			token2:  token{Type: identifier, Value: "test", Line: 200, Column: 23, Index: 205},
			expects: false,
		},
		{
			token1:  token{Type: identifier, Value: "test", Line: 15, Column: 23, Index: 205},
			token2:  token{Type: identifier, Value: "test", Line: 15, Column: 11, Index: 205},
			expects: false,
		},
		{
			token1:  token{Type: identifier, Value: "test", Line: 15, Column: 23, Index: 205},
			token2:  token{Type: identifier, Value: "test", Line: 15, Column: 23, Index: 134},
			expects: true,
		},
	}

	for _, test := range tests {
		if res := test.token1.equal(test.token2); res != test.expects {
			if test.expects == true {
				t.Errorf("Token %v does not match token %v when it should.", test.token1, test.token2)
			} else {
				t.Errorf("Token %v matches token %v when it should not.", test.token1, test.token2)
			}
		}
	}
}

func TestIsEarlier(t *testing.T) {
	tests := []struct {
		t1      token
		t2      token
		expects bool
	}{
		{
			t1:      token{Type: identifier, Value: "beats", Line: 1, Column: 5},
			t2:      token{Type: identifier, Value: "beats", Line: 1, Column: 5},
			expects: false,
		},
		{
			t1:      token{Type: separator, Value: "{", Line: 0, Column: 6},
			t2:      token{Type: identifier, Value: "beats", Line: 1, Column: 5},
			expects: true,
		},
		{
			t1:      token{Type: separator, Value: "}", Line: 3, Column: 5},
			t2:      token{Type: separator, Value: "{", Line: 0, Column: 6},
			expects: false,
		},
	}

	for _, test := range tests {
		if res := test.t1.isEarlier(test.t2); res != test.expects {
			t.Errorf("%v.isEarlier(%v) return unexpected. Got %t, Expected %t", test.t1, test.t2, res, test.expects)
		}
	}
}

func TestEarliestToken(t *testing.T) {
	tokens := []token{
		{Type: identifier, Value: "beats", Line: 1, Column: 5},
		{Type: separator, Value: "{", Line: 0, Column: 6},
		{Type: separator, Value: "}", Line: 3, Column: 5},
		{Type: literal, Value: "5044", Line: 2, Column: 17},
	}
	expected := token{Type: separator, Value: "{", Line: 0, Column: 6}

	if res := earliestToken(tokens); !res.equal(expected) {
		t.Errorf("Earliest token does not match expected. Got %v, Expected %v", res, expected)
	}
}
