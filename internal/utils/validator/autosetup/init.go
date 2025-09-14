package autosetup

import (
	validator2 "github.com/Wuchieh/go-server/internal/utils/validator"
	"github.com/duke-git/lancet/v2/pointer"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ginValidator struct {
	validate *validator.Validate
}

func (g ginValidator) ValidateStruct(a any) error {
	a = pointer.ExtractPointer(a)
	return g.validate.Struct(a)
}

func (g ginValidator) Engine() any {
	return g.validate
}

func init() {
	v := new(ginValidator)
	v.validate = validator2.GetValidate()
	binding.Validator = v
}
