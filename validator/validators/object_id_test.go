package validators

import (
	"testing"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type test struct {
	Name primitive.ObjectID `validate:"oid"`
}

func TestValidateObjectID(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("oid", ValidateObjectID)
	s := test{}
	err := validate.Struct(s)
	if err != nil {
		t.Error(err)
	}
}

func TestFailValidateObjectID(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("oid", ValidateObjectID)
	s := test{
		Name: primitive.NilObjectID,
	}
	err := validate.Struct(s)
	if err != nil {
		t.Error(err)
	}
}
