package handlers

import (
	"api/internal/config"
	"api/internal/services"
	"api/internal/utils"
	"net/http"

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

	if valid := utils.BindAndValidateQuery(context, &query); !valid {
		return
	}

	response, err := h.authService.GenerateToken(query.Code, query.Challenge)

	if err != nil {
		HandleError(context, err)
		return
	}
	context.SetCookie("refresh_token", response.RefreshToken, config.Current.RefreshTokenTTL, "/", "", false, true)
	context.SetCookie("id_token", response.IDToken, config.Current.IdTokenTTL, "/", "", false, true)
	context.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{
		"id_token":   response.IDToken,
		"expires_in": config.Current.IdTokenTTL,
	}})
}

func (h *AuthHandler) PostRefreshToken(context *gin.Context) {
	refreshToken, err := context.Cookie("refresh_token")

	if err != nil || refreshToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	response, err := h.authService.RefreshToken(refreshToken)

	if err != nil {
		HandleError(context, err)
		return
	}
	context.SetCookie("id_token", response.IDToken, config.Current.IdTokenTTL, "/", "", false, true)
	context.JSON(http.StatusOK, gin.H{"data": map[string]interface{}{
		"id_token":   response.IDToken,
		"expires_in": config.Current.IdTokenTTL,
	}})
}

func (h *AuthHandler) GetTokenData(context *gin.Context) {
	token, err := context.Cookie("id_token")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	response, err := h.authService.ValidateToken(token)

	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": response})
}
