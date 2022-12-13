package validator

import (
	"github.com/bacheha/horus/validator/validators"
	val "github.com/go-playground/validator"
)

var defaultValidators = map[string]val.Func{
	"oid": validators.ValidateObjectID,
}

type Validator struct {
	validate *val.Validate
}

func (v *Validator) ValidateStruct(o interface{}) error {
	return v.validate.Struct(o)
}

func (v *Validator) AddValidator(key string, fn val.Func) error {
	if err := v.validate.RegisterValidation(key, fn); err != nil {
		return err
	}
	return nil
}

func New() (*Validator, error) {
	v := val.New()
	validator := &Validator{validate: v}
	for key, fn := range defaultValidators {
		if err := validator.AddValidator(key, fn); err != nil {
			return nil, err
		}
	}
	return validator, nil
}
