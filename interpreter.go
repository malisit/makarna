package main

import (
	"fmt"
	"strconv"
	"log"
)

type Interpreter struct {
	lexer Lexer
	currentToken Token
}

func (i *Interpreter) invalidSyntax() Error {
	return Error{fmt.Sprint("Syntax Error at ", i.lexer.position)}
}

func (i *Interpreter) consume(tokenType string) {
	if i.currentToken.type_ == tokenType {
		i.currentToken = i.lexer.getNextToken()
	} else {
		log.Fatalln(i.invalidSyntax().text)
	}
}

func (i *Interpreter) factor() float64 {
	token := i.currentToken

	if token.type_ == NUM {
		i.consume(NUM)
		result, err := strconv.ParseFloat(token.value_, 64)
		if err == nil {
			return result
		} else {
			log.Fatalln(i.invalidSyntax().text)
		}

	} else if token.type_ == LPAR {
		i.consume(LPAR)
		result := i.expression()
		i.consume(RPAR)
		return result
	}

	return 0

}

func (i *Interpreter) term() float64 {
	result := i.factor()

	for i.currentToken.type_ == MUL || i.currentToken.type_ == DIV {
		token := i.currentToken

		if token.type_ == MUL {
			i.consume(MUL)
			result = result * i.factor()
		} else if token.type_ == DIV {
			i.consume(DIV)
			result = result / i.factor()
		}

	}
	return result
}

func (i *Interpreter) expression() float64 {
	result := i.term()

	for i.currentToken.type_ == SUB || 	i.currentToken.type_ == ADD {
		token := i.currentToken

		if token.type_ == ADD {
			i.consume(ADD)
			result = result + i.term()
		} else if token.type_ == SUB {
			i.consume(SUB)
			result = result - i.term()
		}
	}

	return result
}

