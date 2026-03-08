package middlewares

import (
	"api/internal/config"
	"api/internal/dtos"
	"api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	if config.Current.Env == config.Test {
		context.Next()
		return
	}

	token, err := context.Cookie("id_token")
	if err != nil || token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	user, err := m.authService.ValidateToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	context.Set("user", user)
	context.Next()
}

func (m *Middlewares) IsOwn(context *gin.Context) {
	if config.Current.Env == config.Test {
		context.Next()
		return
	}

	user, exists := GetUserFromContext(context)

	if !exists {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	id := context.Param("id")

	paramId, err := uuid.Parse(id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if user.ID != paramId {
		context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}
}

func GetUserFromContext(context *gin.Context) (*dtos.UserDto, bool) {
	user, exists := context.Get("user")
	if !exists {
		return nil, false
	}

	userData, ok := user.(*dtos.UserDto)
	if !ok {
		return nil, false
	}

	return userData, true
}
