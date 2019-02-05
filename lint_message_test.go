package main

import "testing"

func TestFormat(t *testing.T) {
	m := message{
		filename: "pipelines/logstash.conf",
		line:     157,
		column:   11,
		msg:      "'upd' is unknown input plugin",
	}
	expected := "pipelines/logstash.conf:157:11: 'upd' is unknown input plugin"

	if res := m.format(); res != expected {
		t.Errorf("Formatted Message does not match expected.\nGot: \"%s\"\nExpected: \"%s\"", res, expected)
	}
}
