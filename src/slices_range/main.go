package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"hola", "mundo"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	fmt.Println(esPalindromo("Ama"))
}

func esPalindromo(text string) bool {

	var textReverse string
	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		return true
	} else {
		return false
	}

}
