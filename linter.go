package main

type linter struct {
	messages  []message
	tokenizer tokenizer
}

func (l *linter) addMessage(msg message) {
	if l.messages == nil {
		l.messages = make([]message, 0, 10)
	}

	l.messages = append(l.messages, msg)
}
