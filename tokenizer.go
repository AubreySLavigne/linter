package main

import (
	"regexp"
)

// TODO: Can this use the regexp.syntax package instead? "Package syntax
//       parses regular expressions into parse trees and compiles parse trees
//       into programs."

// TODO: Change the source to an io.Reader type, for flexibility in input
type tokenizer struct {
	source  string
	rules   []rule
	pointer int
}

// TODO: Can I rewrite this with the io.Read interface?
func (tz *tokenizer) next() (token, error) {
	searchSpace := tz.source[tz.pointer:]

	for _, r := range tz.rules {

		re, err := regexp.Compile(r.pattern)
		if err != nil {
			return token{}, err
		}

		match := re.FindString(searchSpace)
		matchLen := len(match)
		if matchLen > 0 {
			// TODO: This should be moved forward until after the newly found
			//       token, rather than just moving incrementally forward
			tz.pointer++

			// TODO: Parse for line and column numbers
			return token{
				Type:   tokenType(r.name),
				Value:  match,
				Line:   0,
				Column: 0,
			}, nil
		}
	}

	return token{}, nil
}

func (tz *tokenizer) addRule(r rule) {
	tz.rules = append(tz.rules, r)
}

func (tz *tokenizer) reset() {
	tz.pointer = 0
}
