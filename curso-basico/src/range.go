package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"hola", "que", "hace"}
	for i, element := range slice {
		fmt.Println(i, element)
	}
	isPalindrome("Neuquen")
}

func isPalindrome(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if strings.ToLower(text) == strings.ToLower(textReverse) {
		fmt.Println("Es un Palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}
