package main

import "testing"

func TestAddMessage(t *testing.T) {
	tests := []struct {
		msg     message
		expects []message
	}{
		{
			msg: message{filename: "abc", line: 12, column: 34, msg: "Basic Message"},
			expects: []message{
				message{filename: "abc", line: 12, column: 34, msg: "Basic Message"},
			},
		},
		{
			msg: message{filename: "def", line: 45, column: 56, msg: "Second Message"},
			expects: []message{
				message{filename: "abc", line: 12, column: 34, msg: "Basic Message"},
				message{filename: "def", line: 45, column: 56, msg: "Second Message"},
			},
		},
	}

	lint := linter{}

	for _, test := range tests {
		lint.addMessage(test.msg)
		assertMessages(t, lint.messages, test.expects)
	}
}

func assertMessages(t *testing.T, messages, expected []message) {
	countExpected := len(expected)
	if count := len(messages); count != countExpected {
		t.Errorf("Number of messages do not match. Got %d, Expected %d", count, countExpected)
		return
	}

	for i := range expected {
		if messages[i].filename != expected[i].filename ||
			messages[i].line != expected[i].line ||
			messages[i].column != expected[i].column ||
			messages[i].msg != expected[i].msg {
			t.Errorf("Message at index %d does not match expected. Got %v, Expected %v", i, messages[i], expected[i])
		}
	}
}
