package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (j *JwtService) ValidateJWT(jwtToken string) (jwt.MapClaims, error) {
	var secret = []byte(os.Getenv("JWT_SECRET"))

	parsed, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !parsed.Valid {
		return nil, err
	}

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}

func (j *JwtService) GenerateJWT(email string) (string, error) {
	var secret = []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}