package middlewares

import (
	"api/internal/errors"
	"api/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middlewares struct {
	authService  *services.AuthService
	usersService *services.UsersService
}

func New(authService *services.AuthService, usersService *services.UsersService) *Middlewares {
	return &Middlewares{
		authService:  authService,
		usersService: usersService,
	}
}

func (m *Middlewares) IsAuthenticated(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(errors.NewResponseError(errors.Unauthorized))
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")
	if len(token) == 0 {
		context.AbortWithStatusJSON(errors.NewResponseError(errors.Unauthorized))
		return
	}

	user, err := m.authService.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(errors.NewResponseError(errors.Unauthorized))
		context.Abort()
		return
	}

	context.Set("user", user)
	context.Next()
}
