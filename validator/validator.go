package validator

import (
	"github.com/bacheha/horus/validator/validators"
	val "github.com/go-playground/validator"
)

type Validator struct {
	validate *val.Validate
}

func (v *Validator) ValidateStruct(o interface{}) error {
	return v.validate.Struct(o)
}

func (v *Validator) AddValidator(key string, fn val.Func) {
	v.validate.RegisterValidation(key, fn)
}

func New() *Validator {
	v := val.New()
	validator := &Validator{validate: v}
	validator.AddValidator("oid", validators.ValidateObjectID)
	return validator
}
