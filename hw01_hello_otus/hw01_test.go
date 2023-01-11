package main

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	var textTesting, result = "Hello, OTUS!", "!SUTO ,olleH"
	realResult := Reverse(textTesting)
	if realResult != result {
		t.Errorf("%s != %s", realResult, result)
		fmt.Println("invalid output:", realResult)
	} else {
		fmt.Println(realResult, "=", result)
	}
}
