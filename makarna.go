package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("makarna > ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		if text != "exit" {
			result := runCode(text)
			fmt.Println(result)
		} else {
			break
		}
	}

}

func runCode(text string) float64 {
	lexer := Lexer{text, 0, string([]rune(text)[0])}
	interpreter := Interpreter{lexer, lexer.getNextToken()}
	result := interpreter.expression()
	
	return result
}