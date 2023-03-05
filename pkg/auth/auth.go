package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	errSigningMethod = errors.New("error in signing method")
	errInvalidToken  = errors.New("token is invalid")
)

type Auth struct {
	key    string
	expire int
}

// New builds a new auth struct for handling JWT tokens.
func New(cfg Config) *Auth {
	return &Auth{
		key:    cfg.PrivateKey,
		expire: cfg.ExpireTime,
	}
}

// GenerateJWT creates a new JWT token.
func (a *Auth) GenerateJWT(clientID, appKey string) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// create claims
	claims := token.Claims.(jwt.MapClaims)
	claims["client_id"] = clientID
	claims["app_key"] = appKey

	// generate token string
	tokenString, err := token.SignedString([]byte(a.key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT gets a token string and extracts the data.
func (a *Auth) ParseJWT(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errSigningMethod
		}

		return []byte(a.key), nil
	})
	if err != nil {
		return "", "", err
	}

	// taking out claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["client_id"].(string)
		key := claims["app_key"].(string)

		return id, key, nil
	}

	return "", "", errInvalidToken
}
