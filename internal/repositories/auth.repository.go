package repositories

import (
	"api/internal/structs"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gorm.io/gorm"
)

type AuthRepository interface {
	GenerateToken(code string, challenge string) (*structs.GoogleTokenResponse, error)
	ValidateToken(token string) (*structs.GoogleTokenValidationResponse, error)
	RefreshToken(refreshToken string) (*structs.GoogleTokenRefreshResponse, error)
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		DB: db,
	}
}

func (r *authRepository) GenerateToken(code string, challenge string) (*structs.GoogleTokenResponse, error) {
	body := url.Values{}
	body.Set("client_id", os.Getenv("CLIENT_ID"))
	body.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	body.Set("redirect_uri", os.Getenv("REDIRECT_URI"))
	body.Set("grant_type", "authorization_code")
	body.Set("code", code)
	body.Set("code_verifier", challenge)

	resp, err := http.Post(os.Getenv("OAUTH_TOKEN_URL"), "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}

	var tokenRes structs.GoogleTokenResponse

	if err := json.Unmarshal(data, &tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}

func (r *authRepository) ValidateToken(token string) (*structs.GoogleTokenValidationResponse, error) {
	req, err := http.NewRequest("GET", os.Getenv("OAUTH_ID_URL"), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}

	var userInfo structs.GoogleTokenValidationResponse

	if err := json.Unmarshal(data, &userInfo); err != nil {
		return nil, err
	}

	if userInfo.Email == "" {
		return nil, errors.New("invalid user")
	}

	return &userInfo, nil
}

func (r *authRepository) RefreshToken(refreshToken string) (*structs.GoogleTokenRefreshResponse, error) {
	body := url.Values{}
	body.Set("client_id", os.Getenv("CLIENT_ID"))
	body.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	body.Set("refresh_token", refreshToken)
	body.Set("grant_type", "refresh_token")

	resp, err := http.Post(os.Getenv("OAUTH_TOKEN_URL"), "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}

	var tokenRes structs.GoogleTokenRefreshResponse

	if err := json.Unmarshal(data, &tokenRes); err != nil {
		return nil, err
	}

	return &tokenRes, nil
}
