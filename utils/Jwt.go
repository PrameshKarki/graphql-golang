package utils

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenMetadata struct {
	ID    string
	Email string
}

func SignAccessToken(payload TokenMetadata) (string, error) {
	tokenSecret := GoDotEnv("TOKEN_SECRET")
	return signToken(map[string]interface{}{
		"id":    payload.ID,
		"email": payload.Email,
	}, tokenSecret, 15*time.Minute)

}

func signToken(Data map[string]interface{}, secret string, ExpiredAt time.Duration) (string, error) {
	claims := jwt.MapClaims{}

	for key, value := range Data {
		claims[key] = value
	}
	claims["exp"] = time.Now().Add(ExpiredAt).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(accessToken, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(accessToken *jwt.Token) TokenMetadata {
	var token TokenMetadata
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)
	return token
}
