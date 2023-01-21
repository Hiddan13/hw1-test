package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{input: "a2b1nn0", expected: "aabn"},
		{input: "mm2r3b0", expected: "mmmrrr"},
		{input: "e3", expected: "eee"},
		{input: "b5x4", expected: "bbbbbxxxx"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	testsError := []struct {
		input         string
		expectedError error
	}{
		{input: "b2m41i", expectedError: errors.New("number after number")},
		{input: "aaa10b", expectedError: errors.New("number after number")},
		{input: "nn23e", expectedError: errors.New("number after number")},
		{input: "3n5g", expectedError: errors.New("first rune is not letter")},
		{input: "3abc", expectedError: errors.New("first rune is not letter")},
		{input: "45", expectedError: errors.New("first rune is not letter")},
	}
	for _, tc := range testsError {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			_, errorUnpack := Unpack(tc.input)
			errors.Is(errorUnpack, tc.expectedError)
		})
	}
}
