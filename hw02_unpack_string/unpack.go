package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const c_mixed = "g2n0bb"

var its_slash bool

var its_doubleslash bool

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
		case typ == "letter":
			switch {
			case typ1 == "letter":
				result.WriteString(string(value1))

			case typ1 == "\\":
				return "letter after \\", ErrInvalidString
			}

		case typ == "\\n" && value1 != "":
			result.WriteString(string(value1))
			value = "\\n"

		case typ == "\\" && value1 != "" && !its_slash:

			its_slash = true
			if typ1 == "letter" {
				result.WriteString(string(value1))
			}
		case typ == "\\" && typ1 == "\\" && !its_doubleslash:

			its_doubleslash = true
			its_slash = false
		case typ == "number" && value1 != "":
			if s, err := strconv.Atoi(string(value)); err == nil {
				numberint = s

			}

			switch {

			case typ1 == "number" && !its_slash && !its_doubleslash:
				return "number after number", ErrInvalidString
			case typ1 == "letter" || typ1 == "\\n":
				if numberint == 0 {

					break
				}
				result.WriteString(string(strings.Repeat(string(value1), numberint)))

			case its_doubleslash:
				if numberint == 0 {

					break
				}
				result.WriteString(string(strings.Repeat(string(value1), numberint)))
				its_doubleslash = false

			case its_slash:
				if typ == "number" && typ1 == "number" {
					if numberint == 0 {

						break
					}
					result.WriteString(string(strings.Repeat(string(value1), numberint-1)))
					its_slash = false
					break
				}
				result.WriteString(string(value))
			}

		case typ == "\\n" && value1 == "" || typ == "\\" && value1 == "" || typ == "number" && value1 == "":
			return "first rune is not letter", ErrInvalidString
		}
		value1 = value
		typ1 = typ
	}
	if typ1 == "letter" {

		result.WriteString(string(value))
	}
	return result.String(), nil
}
func Define(c rune) (typ string, value string) {
	switch {
	case !unicode.IsDigit(c):
		switch {
		case c != 10 && c != 92:
			res := "letter"
			return res, string(c)
		case c == 10: //\n
			res := "\\n"
			return res, string(c)
		case c == 92: //\
			res := "\\"
			return res, string(c)
		}
	case unicode.IsDigit(c):
		res := "number"
		return res, string(c)
	}
	return
}

func main() {
	fmt.Println(Unpack(c_mixed))

}
