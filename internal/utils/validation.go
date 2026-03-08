package utils

import (
	"api/internal/errors"
	"api/internal/logger"

	"github.com/gin-gonic/gin"
)

var log = logger.New("Validations")

func BindAndValidateQuery[T any](context *gin.Context, obj *T) bool {
	if err := context.ShouldBindQuery(obj); err != nil {
		log.Debugf("Query validation error: %v", err)
		context.AbortWithStatusJSON(errors.ErrBadRequest.StatusCode, gin.H{"message": errors.ErrBadRequest.Message})
		return false
	}
	return true
}

func BindAndValidateBody[T any](context *gin.Context, obj *T) bool {
	if err := context.ShouldBindJSON(obj); err != nil {
		log.Debugf("Body validation error: %v", err)
		context.AbortWithStatusJSON(errors.ErrBadRequest.StatusCode, gin.H{"message": errors.ErrBadRequest.Message})
		return false
	}
	return true
}
