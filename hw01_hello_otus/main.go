package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	text := "Hello, OTUS!"
	fmt.Println(HeadText(text))
}
func HeadText(t string) string {
	text01 := stringutil.Reverse(t)
	return text01
}
