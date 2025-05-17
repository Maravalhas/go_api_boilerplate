package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func BindAndValidateQuery[T any](context *gin.Context, obj *T) error {
	if err := context.ShouldBindQuery(obj); err != nil {
		return errors.New("invalid query parameters")
	}
	if err := validate.Struct(obj); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func BindAndValidateBody[T any](context *gin.Context, obj *T) error {
	if err := context.ShouldBind(obj); err != nil {
		return errors.New("invalid body parameters")
	}
	if err := validate.Struct(obj); err != nil {
		return errors.New(err.Error())
	}
	return nil
}
