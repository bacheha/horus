package validators

import (
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateObjectID(fl validator.FieldLevel) bool {
	return primitive.IsValidObjectID(fl.Field().String())
}
