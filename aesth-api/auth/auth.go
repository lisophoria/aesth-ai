package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(jwtToken string) (jwt.MapClaims, error) {
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

func generateJWT(email string) (string, error) {
	var secret = []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func AuthHandler(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}

	err := c.ShouldBindJSON(&request)
	if err != nil || request.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing email"})
		return
	}

	token, err := generateJWT(request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
