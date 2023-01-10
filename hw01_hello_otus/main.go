package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	originalText := "Hello, OTUS!"

	fmt.Println(Reverse(originalText))
}

func Reverse(text string) string {
	resultText := stringutil.Reverse(text)
	return resultText
}
