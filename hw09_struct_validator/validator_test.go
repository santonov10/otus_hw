package hw09structvalidator

import (
	"encoding/json"
	"errors"
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
		Header    []byte `validate:"len:5"`
		Payload   []byte `validate:"notImplement:error"`
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "1",
				Name:   "1",
				Age:    2,
				Email:  "testtest.ru",
				Role:   "4",
				Phones: []string{"test", "12345678910", "123"},
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   errors.New("длина строки '1' должна быть '36', а не '1'"),
				},
				ValidationError{
					Field: "Age",
					Err:   errors.New("число '2' должно быть не меньше чем '18'"),
				},
				ValidationError{
					Field: "Email",
					Err:   errors.New("значение 'testtest.ru' не соответвует регулярному выражению '^\\w+@\\w+\\.\\w+$'"),
				},
				ValidationError{
					Field: "Role",
					Err:   errors.New("'4' должен находиться в массиве '[admin stuff]'"),
				},
				ValidationError{
					Field: "Phones",
					Err:   errors.New("длина строки 'test' должна быть '11', а не '4'; длина строки '123' должна быть '11', а не '3'"),
				},
			},
		},

		{
			in: User{
				ID:     "123456789012345678901234567890123456",
				Name:   "",
				Age:    44,
				Email:  "test@test.ru",
				Role:   "admin",
				Phones: nil, // nil check
				meta:   nil,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Phones",
					Err:   errors.New("длинна слайса '[]' = 0"),
				},
			},
		},

		{
			in: App{Version: "err"},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   errors.New("длина строки 'err' должна быть '5', а не '3'"),
				},
			},
		},
		{
			in:          App{Version: "1.2.3"},
			expectedErr: nil,
		},

		{
			in: Token{
				Header:    []byte{'q'},
				Payload:   []byte{'q', '2'},
				Signature: nil,
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Header",
					Err:   errors.New("отсутствует метод валидации для тэга: len и типа значения uint8"),
				},
				ValidationError{
					Field: "Payload",
					Err:   errors.New("отсутствует метод валидации для тэга: notImplement и типа значения uint8"),
				},
			},
		},

		{
			in: Response{
				Code: 200,
				Body: "body",
			},
			expectedErr: nil,
		},

		{
			in: Response{
				Code: 205,
				Body: "body",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   errors.New("'%!s(int=205)' должен находиться в массиве '[200 404 500]'"),
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			if err == nil || tt.expectedErr == nil {
				require.Equal(t, err, tt.expectedErr)
			} else {
				require.Equal(t, err.Error(), tt.expectedErr.Error())
			}
			// require.ErrorIs(t, err, tt.expectedErr) // - не работает такой вариант, наверное не достаточно создать ошибку с помощью errors.New()
		})
	}
}
