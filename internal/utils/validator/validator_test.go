package validator_test

import (
	"testing"

	"github.com/Wuchieh/go-server/internal/utils/validator"
	"github.com/duke-git/lancet/v2/pointer"
)

func TestValidate(t *testing.T) {
	v := validator.GetValidate()

	type testCase struct {
		name  string
		data  any
		error bool
	}

	TestCase := []testCase{
		{
			name:  "base",
			data:  "qwer@gmail.com",
			error: false,
		},
		{
			name:  "pointer string",
			data:  pointer.Of("qwer@gmail.com"),
			error: false,
		},
		{
			name:  "data error",
			data:  "asd+asd@gmail.com",
			error: true,
		},
		{
			name:  "email supplier error",
			data:  "qwer@yahoo.com",
			error: true,
		},
		{
			name:  "type error",
			data:  12345,
			error: true,
		},
	}

	for _, a := range TestCase {
		t.Run(a.name, func(t *testing.T) {
			err := v.Var(a.data, "gmail")
			if (err != nil) != a.error {
				t.Errorf("got error: %v, expected error: %v", err, a.error)
			}
		})
	}
}
