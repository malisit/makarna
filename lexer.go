package main

import (
	"fmt"
	"bytes"
	"unicode"
	"log"
)

type Lexer struct {
	text string
	position int
	currentChar string
}

// func (l *Lexer) currentChar() string {
// 	return string([]rune(l.text)[l.position])
// }

func (l *Lexer) invCharError() Error {
	return Error{fmt.Sprint("Invalid character at ", l.position)}
}

func (l *Lexer) move() {
	l.position++
	if l.position > len(l.text) -1 {
		l.currentChar = "NONE"
	} else {
		l.currentChar = string([]rune(l.text)[l.position])
	}
}

func (l *Lexer) wsHandler() {
	for l.currentChar!="NONE" && l.currentChar==" " {
		l.move()
	}
}

func (l *Lexer) num() string {
	var buffer bytes.Buffer
	for l.currentChar!="NONE" && (unicode.IsNumber([]rune(l.currentChar)[0])||l.currentChar==".") {
		buffer.WriteString(l.currentChar)
		l.move()
	}

	return buffer.String()
}

func (l *Lexer) getNextToken() Token {
	for l.currentChar != "NONE" {
		if l.currentChar == " " {
			l.wsHandler()
			continue
		}
		
		if unicode.IsNumber([]rune(l.currentChar)[0]) {
			return Token{NUM, l.num()}
		}

		if l.currentChar == "+" {
			l.move()
			return Token{ADD, "+"}
		}

		if l.currentChar == "-" {
			l.move()
			return Token{SUB, "-"}
		}

		if l.currentChar == "*" {
			l.move()
			return Token{MUL, "*"}
		}

		if l.currentChar == "/" {
			l.move()
			return Token{DIV, "/"}
		}

		if l.currentChar == "(" {
			l.move()
			return Token{LPAR, "("}
		}

		if (l.currentChar == ")") {
			l.move()
			return Token{RPAR, ")"}
		}

		log.Fatalln(l.invCharError().text)
	}

	return Token{EOF, "NONE"}
	
}
