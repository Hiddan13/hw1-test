package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	cMixed = "n0bb"
	letter = "letter"
	number = "number"
)

var itsSlash bool

var itsDoubleslash bool

var numberint int

var ErrInvalidString = errors.New("invalid string")

func Unpack(mixed string) (string, error) {
	var result strings.Builder
	var value, typ string
	var value1, typ1 string
	var runeline = []rune(mixed)
	fmt.Println(runeline)
	for _, c := range runeline {
		typ, value = Define(c)
		switch {
		case typ == letter:
			switch {
			case typ1 == letter:
				result.WriteString(value1)
			case typ1 == "\\":
				return "letter after \\", ErrInvalidString
			}
		case typ == "\\n" && value1 != "":
			result.WriteString(value1)
			value = "\\n"
		case typ == "\\" && value1 != "" && !itsSlash:
			itsSlash = true
			if typ1 == letter {
				result.WriteString(value1)
			}
		case typ == "\\" && typ1 == "\\" && !itsDoubleslash:
			itsDoubleslash = true
			itsSlash = false
		case typ == number && value1 != "":
			if s, err := strconv.Atoi(string(value)); err == nil {
				numberint = s
			}
			switch {
			case typ1 == number && !itsSlash && !itsDoubleslash:
				return "number after number", ErrInvalidString
			case typ1 == letter || typ1 == "\\n":
				if numberint == 0 {
					break
				}
				result.WriteString(strings.Repeat(string(value1), numberint))
			case itsDoubleslash:
				if numberint == 0 {
					break
				}
				result.WriteString(strings.Repeat(string(value1), numberint))
				itsDoubleslash = false

			case itsSlash:
				if typ == number && typ1 == number {
					if numberint == 0 {
						break
					}
					result.WriteString(strings.Repeat(string(value1), numberint-1))
					itsSlash = false
					break
				}
				result.WriteString(value)
			}
		case typ == "\\n" && value1 == "" || typ == "\\" && value1 == "" || typ == number && value1 == "":
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
	switch {
	case !unicode.IsDigit(c):
		switch {
		case c != 10 && c != 92:
			res := letter
			return res, string(c)
		case c == 10: // \n
			res := "\\n"
			return res, string(c)
		case c == 92: // \
			res := "\\"
			return res, string(c)
		}
	case unicode.IsDigit(c):
		res := number
		return res, string(c)
	}
	return
}

func main() {
	fmt.Println(Unpack(cMixed))

}
