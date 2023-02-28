package auth

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
func (a *Auth) GenerateJWT(username string, password string) (string, time.Time, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)
	expireTime := time.Now().Add(time.Duration(a.expire) * time.Minute)

	// create claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = expireTime.Unix()
	claims["username"] = username
	claims["password"] = password

	// generate token string
	tokenString, err := token.SignedString([]byte(a.key))
	if err != nil {
		return "", expireTime, err
	}

	return tokenString, expireTime, nil
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
		username := claims["username"].(string)
		password := claims["password"].(string)

		return username, password, nil
	}

	return "", "", errInvalidToken
}
