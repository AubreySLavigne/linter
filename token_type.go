package main

type tokenType int

const (
	identifier tokenType = iota + 1
	keyword
	separator
	operator
	literal
	comment
)
