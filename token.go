package main

type token struct {
	Type  tokenType
	Value string

	// Line and column are for 2-dimensional indexes. Index is for 1-dimensional
	Line   int
	Column int
	Index  int
}

func (t *token) equal(t2 token) bool {
	return t.Type == t2.Type &&
		t.Value == t2.Value &&
		t.Line == t2.Line &&
		t.Column == t2.Column
}

func (t *token) isEarlier(t2 token) bool {
	if t.Line < t2.Line {
		return true
	}
	if t.Line == t2.Line && t.Column < t2.Column {
		return true
	}

	return false
}

func earliestToken(tokens []token) token {
	var res = tokens[0]
	for _, t := range tokens {
		if t.isEarlier(res) {
			res = t
		}
	}
	return res
}
