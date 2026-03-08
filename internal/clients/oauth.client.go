package clients

import (
	"api/internal/config"
	"api/internal/models"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwksCache struct {
	mu        sync.RWMutex
	keys      map[string]*rsa.PublicKey
	fetchedAt time.Time
	ttl       time.Duration
}

type jwksResponse struct {
	Keys []jwkKey `json:"keys"`
}

type jwkKey struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func newJWKSCache(ttl time.Duration) *jwksCache {
	return &jwksCache{
		keys: make(map[string]*rsa.PublicKey),
		ttl:  ttl,
	}
}

func (c *jwksCache) getKey(kid string, jwksUrl string) (*rsa.PublicKey, error) {
	c.mu.RLock()
	if time.Since(c.fetchedAt) < c.ttl {
		if key, ok := c.keys[kid]; ok {
			c.mu.RUnlock()
			return key, nil
		}
	}
	c.mu.RUnlock()

	if err := c.refresh(jwksUrl); err != nil {
		return nil, err
	}

	c.mu.RLock()
	defer c.mu.RUnlock()
	key, ok := c.keys[kid]
	if !ok {
		return nil, errors.New("signing key not found")
	}
	return key, nil
}

func (c *jwksCache) refresh(jwksUrl string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if time.Since(c.fetchedAt) < c.ttl {
		return nil
	}

	resp, err := http.Get(jwksUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var jwks jwksResponse
	if err := json.Unmarshal(data, &jwks); err != nil {
		return err
	}

	keys := make(map[string]*rsa.PublicKey, len(jwks.Keys))
	for _, k := range jwks.Keys {
		if k.Kty != "RSA" {
			continue
		}

		nBytes, err := base64.RawURLEncoding.DecodeString(k.N)
		if err != nil {
			continue
		}
		eBytes, err := base64.RawURLEncoding.DecodeString(k.E)
		if err != nil {
			continue
		}

		n := new(big.Int).SetBytes(nBytes)
		e := new(big.Int).SetBytes(eBytes)

		keys[k.Kid] = &rsa.PublicKey{
			N: n,
			E: int(e.Int64()),
		}
	}

	c.keys = keys
	c.fetchedAt = time.Now()
	return nil
}

type oauthClient struct {
	jwks *jwksCache
}

func NewOAuthClient() OAuthClient {
	return &oauthClient{
		jwks: newJWKSCache(1 * time.Hour),
	}
}

func (c *oauthClient) GenerateToken(code string, challenge string) (*models.Token, error) {
	body := url.Values{}
	body.Set("client_id", config.Current.ClientId)
	body.Set("client_secret", config.Current.ClientSecret)
	body.Set("redirect_uri", config.Current.RedirectUri)
	body.Set("grant_type", "authorization_code")
	body.Set("code", code)
	body.Set("code_verifier", challenge)

	resp, err := http.Post(config.Current.OauthTokenUrl, "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))

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

func (c *oauthClient) ValidateToken(idToken string) (*string, error) {
	token, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("missing kid in token header")
		}

		return c.jwks.getKey(kid, config.Current.OauthJWKSUrl)
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	iss, _ := claims.GetIssuer()
	validIss := false
	for _, allowed := range config.Current.OauthIssuers {
		if iss == allowed {
			validIss = true
			break
		}
	}
	if !validIss {
		return nil, errors.New("invalid issuer")
	}

	aud, _ := claims.GetAudience()
	validAud := false
	for _, a := range aud {
		if a == config.Current.ClientId {
			validAud = true
			break
		}
	}
	if !validAud {
		return nil, errors.New("invalid audience")
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return nil, errors.New("email not found in token")
	}

	emailVerified, ok := claims["email_verified"].(bool)
	if !ok || !emailVerified {
		return nil, errors.New("email not verified")
	}

	return &email, nil
}

func (c *oauthClient) RefreshToken(refreshToken string) (*models.Token, error) {
	body := url.Values{}
	body.Set("client_id", config.Current.ClientId)
	body.Set("client_secret", config.Current.ClientSecret)
	body.Set("refresh_token", refreshToken)
	body.Set("grant_type", "refresh_token")

	resp, err := http.Post(config.Current.OauthTokenUrl, "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))

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
