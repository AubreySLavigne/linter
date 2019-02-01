package main

type token struct {
	Type   tokenType
	Value  string
	Line   int
	Column int
}

func (t *token) equal(t2 token) bool {
	return t.Type == t2.Type &&
		t.Value == t2.Value &&
		t.Line == t2.Line &&
		t.Column == t2.Column
}

type tokenType string
