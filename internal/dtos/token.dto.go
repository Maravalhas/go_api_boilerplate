package dtos

import "api/internal/models"

type TokenDTO struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}

func ToTokenDTO(token *models.Token) *TokenDTO {
	return &TokenDTO{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	}
}
