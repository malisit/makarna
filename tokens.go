package main

import (
	"bytes"
)

type Token struct {
	type_, value_ string
}

func (t *Token) toString() string {
	var buffer bytes.Buffer

	buffer.WriteString("Token({")
	buffer.WriteString(t.type_)
	buffer.WriteString("}, {")
	buffer.WriteString(t.value_)
	buffer.WriteString("})")

	return buffer.String()
}