package handlers

import (
	"api/internal/errors"
	"api/internal/services"
	"api/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) PostToken(context *gin.Context) {
	var query struct {
		Code      string `form:"code"`
		Challenge string `form:"challenge"`
	}

	if err := utils.BindAndValidateQuery(context, &query); err != nil {
		context.AbortWithStatusJSON(errors.NewResponseError(errors.BadRequest))
		return
	}

	response, err := h.authService.GenerateToken(query.Code, query.Challenge)

	if err != nil {
		context.AbortWithStatusJSON(errors.NewResponseError(errors.BadRequest))
		return
	}
	context.SetCookie("refresh_token", response.RefreshToken, 86400, "/", "", false, true)
	context.JSON(http.StatusOK, gin.H{"data": *response})
}

func (h *AuthHandler) PostRefreshToken(context *gin.Context) {
	refreshToken, err := context.Cookie("refresh_token")

	if err != nil || refreshToken == "" {
		context.AbortWithStatusJSON(errors.NewResponseError(errors.Unauthorized))
		return
	}

	response, err := h.authService.RefreshToken(refreshToken)

	if err != nil {
		context.AbortWithStatusJSON(errors.NewResponseError(errors.InternalError))
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": *response})
}

func (h *AuthHandler) GetTokenData(context *gin.Context) {
	token := context.GetHeader("Authorization")

	token = strings.TrimPrefix(token, "Bearer ")

	if len(token) <= 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	response, err := h.authService.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": response})
}
