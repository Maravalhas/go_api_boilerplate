package clients

import (
	"api/internal/config"
	"api/internal/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type googleTokenValidationResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type oauthClient struct {
	config *config.Config
}

func NewOAuthClient(config *config.Config) OAuthClient {
	return &oauthClient{
		config: config,
	}
}

func (c *oauthClient) GenerateToken(code string, challenge string) (*models.Token, error) {
	body := url.Values{}
	body.Set("client_id", c.config.ClientId)
	body.Set("client_secret", c.config.ClientSecret)
	body.Set("redirect_uri", c.config.RedirectUri)
	body.Set("grant_type", "authorization_code")
	body.Set("code", code)
	body.Set("code_verifier", challenge)

	resp, err := http.Post(c.config.OauthTokenUrl, "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}

	var result *models.Token
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *oauthClient) ValidateToken(token string) (*string, error) {
	req, err := http.NewRequest("GET", c.config.OauthIdUrl, nil)
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

	var formattedResponse *googleTokenValidationResponse

	if err := json.Unmarshal(data, &formattedResponse); err != nil {
		return nil, err
	}

	if formattedResponse.Email == "" {
		return nil, errors.New("invalid user")
	}

	return &formattedResponse.Email, nil
}

func (c *oauthClient) RefreshToken(refreshToken string) (*models.Token, error) {
	body := url.Values{}
	body.Set("client_id", c.config.ClientId)
	body.Set("client_secret", c.config.ClientSecret)
	body.Set("refresh_token", refreshToken)
	body.Set("grant_type", "refresh_token")

	resp, err := http.Post(c.config.OauthTokenUrl, "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(data))
	}

	var tokenRes *models.Token

	if err := json.Unmarshal(data, &tokenRes); err != nil {
		return nil, err
	}

	return tokenRes, nil
}
