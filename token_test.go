package main

import "testing"

func TestTokenEqual(t *testing.T) {
	tests := []struct {
		token1  token
		token2  token
		expects bool
	}{
		{
			token1:  token{Type: tokenType("type1"), Value: "test", Line: 15, Column: 23},
			token2:  token{Type: tokenType("type1"), Value: "test", Line: 15, Column: 23},
			expects: true,
		},
		{
			token1:  token{Type: tokenType("type1"), Value: "test", Line: 15, Column: 23},
			token2:  token{Type: tokenType("different"), Value: "test", Line: 15, Column: 23},
			expects: false,
		},
		{
			token1:  token{Type: tokenType("type1"), Value: "test", Line: 15, Column: 23},
			token2:  token{Type: tokenType("type1"), Value: "different", Line: 15, Column: 23},
			expects: false,
		},
		{
			token1:  token{Type: tokenType("type1"), Value: "test", Line: 15, Column: 23},
			token2:  token{Type: tokenType("type1"), Value: "test", Line: 200, Column: 23},
			expects: false,
		},
		{
			token1:  token{Type: tokenType("type1"), Value: "test", Line: 15, Column: 23},
			token2:  token{Type: tokenType("type1"), Value: "test", Line: 15, Column: 11},
			expects: false,
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