package main

import "fmt"

// The standardized output format for linters:
// filename.ft:line:column: message
//
// Example:
// pipelines/logstash.conf:157:11: 'upd' is unknown input plugin

type message struct {
	filename string
	line     int
	column   int
	msg      string
}

func (m *message) format() string {
	return fmt.Sprintf("%s:%d:%d: %s", m.filename, m.line, m.column, m.msg)
}
