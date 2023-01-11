package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	cMixed = "a4bc2d5e"
	letter = "letter"
	number = "number"
)

var numberint int

var ErrInvalidString = errors.New("invalid string")

func Unpack(mixed string) (string, error) {
	var result strings.Builder
	var value, typ, value1, typ1 string
	runeline := []rune(mixed)
	fmt.Println(runeline)
	for _, c := range runeline {
		typ, value = Define(c)
		switch {
		case typ == letter:
			if typ1 == letter {
				result.WriteString(value1)
			}
		case typ == number && value1 != "":
			if s, err := strconv.Atoi(value); err == nil {
				numberint = s
			}
			switch {
			case typ1 == number:
				return "number after number", ErrInvalidString
			case typ1 == letter:
				if numberint == 0 {
					break
				}
				result.WriteString(strings.Repeat(value1, numberint))
			}
		case typ == number && value1 == "":
			return "first rune is not letter", ErrInvalidString
		}
		value1 = value
		typ1 = typ
	}
	if typ1 == letter {
		result.WriteString(value)
	}
	return result.String(), nil
}

func Define(c rune) (typ string, value string) {
	if unicode.IsDigit(c) {
		res := number
		return res, string(c)
	}
	res := letter
	return res, string(c)

}

func main() {
	fmt.Println(Unpack(cMixed))
}
