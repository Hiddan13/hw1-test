package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	TextDefauilt = "♥5x▷"
	letter       = "letter"
	number       = "number"
)

var (
	ErrNumAfterNum    = errors.New("number after number")
	ErrFirstNotLetter = errors.New("first rune is not letterr")
)

func Unpack(text string) (string, error) {
	var numberint int
	var result strings.Builder
	var value, typ, value1, typ1 string
	runeline := []rune(text)
	for _, c := range string(runeline) {
		typ, value = DefineTypeOfLetter(c)
		switch {
		case typ == letter:
			if typ1 == letter {
				result.WriteString(value1)
			}
		case typ == number && value1 != "":
			if num, err := strconv.Atoi(value); err == nil {
				numberint = num
			}
			switch {
			case typ1 == number:
				return "", ErrNumAfterNum
			case typ1 == letter:
				if numberint == 0 {
					break
				}
				result.WriteString(strings.Repeat(value1, numberint))
			}
		case typ == number && value1 == "":
			return "", ErrFirstNotLetter
		}
		value1 = value
		typ1 = typ
	}
	if typ1 == letter {
		result.WriteString(value)
	}
	return result.String(), nil
}

func DefineTypeOfLetter(let rune) (typ string, value string) {
	if unicode.IsDigit(let) {
		typeOfLetter := number
		return typeOfLetter, string(let)
	}
	resultType := letter
	return resultType, string(let)
}

func main() {
	fmt.Println(Unpack(TextDefauilt))
}
