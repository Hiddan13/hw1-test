package main

import (
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	var t1, result = "Hello, OTUS!", "!SUTO ,olleH"
	realResult := HeadText(t1)
	if realResult != result {
		t.Errorf("%s != %s", realResult, result)
		fmt.Println("invalid output:", realResult)
	} else {
		fmt.Println(realResult, "=", result)
	}
}
