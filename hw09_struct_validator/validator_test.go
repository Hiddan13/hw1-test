package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	Check struct {
		Ms []string `validate:"len:1"`
		Mi []int    `validate:"max:1"`
	}
)

func TestValidate(t *testing.T) {
	var ve ValidationErrors
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          &User{ID: "333", Name: "Bil", Age: 8, meta: nil},
			expectedErr: ValidationErrors{ValidationError{Field: "Age", Err: ErrMin}},
		},
		{
			in:          &User{ID: "0", Name: "Bod", Age: 44},
			expectedErr: ve,
		},
		{
			in:          &Check{Mi: []int{1, 22}},
			expectedErr: ValidationErrors{ValidationError{Field: "Mi", Err: ErrMax}},
		},
		{
			in:          &[]int{1, 2},
			expectedErr: ErrStruct,
		},
		{
			in:          &User{ID: "1234566", Email: "ii@gmail"},
			expectedErr: ValidationErrors{ValidationError{Field: "Email", Err: ErrRegExp}},
		},
		{
			in:          &App{Version: "1.2.2.222"},
			expectedErr: ValidationErrors{ValidationError{Field: "Version", Err: ErrLen}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			if err := Validate(tt.in); err != nil {
				require.Equal(t, tt.expectedErr, err)
			}
		})
	}
}
