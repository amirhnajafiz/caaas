package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	errSigningMethod = errors.New("error in signing method")
	errInvalidToken  = errors.New("token is invalid")
)

type Auth struct {
	key    string
	salt   string
	expire int
}

// New builds a new auth struct for handling JWT tokens.
func New(cfg Config) *Auth {
	return &Auth{
		key:    cfg.PrivateKey,
		expire: cfg.TokensExpireTime,
		salt:   cfg.EncryptionSalt,
	}
}

// GenerateJWT creates a new JWT token.
func (a *Auth) GenerateJWT(username string) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// create claims
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(a.expire)).Unix()
	claims["authorized"] = true

	// generate token string
	tokenString, err := token.SignedString([]byte(a.salt + a.key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT gets a token string and extracts the data.
func (a *Auth) ParseJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errSigningMethod
		}

		return []byte(a.salt + a.key), nil
	})
	if err != nil {
		return "", err
	}

	// taking out claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		key := claims["username"].(string)

		return key, nil
	}

	return "", errInvalidToken
}
