package auth

import (
	"projeto_chat_backend/pkg/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secreto") // Modifique isso para sua chave secreta!

func GenerateToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &model.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tkStr string) (*model.Claims, error) {
	claims := &model.Claims{}

	token, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func GenerateRefreshToken(user model.User) (*model.RefreshToken, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 7 dias

	claims := &model.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &model.RefreshToken{
		Token:     tokenString,
		UserID:    user.ID,
		Expiry:    expirationTime,
		CreatedAt: time.Now(),
	}, nil
}
