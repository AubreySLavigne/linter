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

	res := earliestToken(matches)

	// TODO: This likely needs to be improved to consider the offset for the new space
	tz.pointer = res.Index + len(res.Value)

	return res, nil
}

// TODO: This maybe possible to optimize/organize better. Refactor this to be
//       more manageable.
func (tz *tokenizer) findUpcomingMatches() ([]token, error) {
	if tz.pointer >= len(tz.source) {
		// TODO: Should this return error, or nil?
		return []token{}, nil
	}
	searchSpace := tz.source[tz.pointer:]

	newlines := newlineIndexes(tz.source)

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

		line, col := lineAndColumn(newlines, tz.pointer+loc[0])

		res = append(res, token{
			Type:   tokenType(r.name),
			Value:  searchSpace[loc[0]:loc[1]],
			Line:   line,
			Column: col,
			Index:  tz.pointer + loc[0],
		})
	}

	return res, nil
}

func lineAndColumn(newlines []int, pos int) (line, col int) {
	hasNewline := false
	highestNlIndex := 0
	for _, nlIndex := range newlines {
		if nlIndex < pos {
			hasNewline = true
			line++
			if nlIndex > highestNlIndex {
				highestNlIndex = nlIndex
			}
		}
	}

	col = pos - highestNlIndex
	if hasNewline {
		// We won't count newline characters as the 0-index for that line, so
		// we need to reduce the value by 1 to reflect that
		col--
	}

	return line, col
}

func newlineIndexes(text string) []int {
	re, _ := regexp.Compile("\n")

	var res []int
	indexes := re.FindAllIndex([]byte(text), -1)
	for _, index := range indexes {
		res = append(res, index[0])
	}

	return res
}

func (tz *tokenizer) addRule(r rule) {
	tz.rules = append(tz.rules, r)
}

func (tz *tokenizer) reset() {
	tz.pointer = 0
}
