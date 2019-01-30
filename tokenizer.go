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

// TODO: Change this to an io.Reader type, for flexibility in input
type tokenizer struct {
	source string
}

func (tz *tokenizer) next() token {
	// TODO: Parse the input for this data. This is currently not functional
	// TODO: Allow the parsing rules to be customized
	return token{
		Type:   tokenType("Identifier"),
		Value:  tz.source,
		Line:   1,
		Column: 1,
	}
}
