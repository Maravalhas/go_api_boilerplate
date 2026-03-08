package handlers

import (
	"api/internal/errors"
	"api/internal/logger"
	stdErrors "errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var log = logger.New("Handlers")

func HandleError(c *gin.Context, err error) {
	var appErr *errors.AppError
	if stdErrors.As(err, &appErr) {
		c.AbortWithStatusJSON(appErr.StatusCode, gin.H{"message": appErr.Message})
		return
	}
	log.Errorf("Unexpected error: %v", err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
}
