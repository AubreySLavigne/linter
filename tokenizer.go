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

func (tz *tokenizer) next() (token, error) {

	matches, err := tz.findUpcomingMatches()
	if err != nil {
		return token{}, err
	}
	if len(matches) == 0 {
		return token{}, nil
	}

	t := matches[0]

	tz.pointer += len(t.Value)

	return t, nil
}

func (tz *tokenizer) findUpcomingMatches() ([]token, error) {
	searchSpace := tz.source[tz.pointer:]

	var res []token
	for _, r := range tz.rules {
		re, err := regexp.Compile(r.pattern)
		if err != nil {
			return nil, err
		}

		loc := re.FindIndex([]byte(searchSpace))
		if loc == nil {
			continue
		}

		res = append(res, token{
			Type:   tokenType(r.name),
			Value:  searchSpace[loc[0]:loc[1]],
			Line:   0,
			Column: tz.pointer + loc[0],
		})
	}

	return res, nil
}

func (tz *tokenizer) addRule(r rule) {
	tz.rules = append(tz.rules, r)
}

func (tz *tokenizer) reset() {
	tz.pointer = 0
}
