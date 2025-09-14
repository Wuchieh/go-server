package validator

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type Validation struct {
	fn                       validator.Func
	callValidationEvenIfNull bool
}

var (
	validate = validator.New(func(v *validator.Validate) {
		v.SetTagName("binding")
	})

	validation = map[string]Validation{
		"gmail": {
			fn:                       isGmail,
			callValidationEvenIfNull: false,
		},
	}

	gmailRegex = regexp.MustCompile(`^[a-z0-9.]+@gmail\.com$`)
)

func isGmail(fl validator.FieldLevel) bool {
	field := fl.Field()

	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			return false
		}
		field = field.Elem()
	}

	if field.Kind() != reflect.String {
		return false
	}

	email := field.String()
	return gmailRegex.MatchString(email)
}

func init() {
	for tag, v := range validation {
		if err := validate.RegisterValidation(tag, v.fn, v.callValidationEvenIfNull); err != nil {
			panic(err)
		}
	}
}

func GetValidate() *validator.Validate {
	return validate
}
