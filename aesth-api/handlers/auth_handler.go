package handlers

import (
	"aesth-api/auth"
	"aesth-api/dto"
	"aesth-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	jwtService *auth.JwtService
	authService *services.AuthService
}

func NewAuthHandler(jwtService *auth.JwtService, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{jwtService, authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input dto.RegisterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Register(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, _ := h.jwtService.GenerateJWT(user.Email)
	c.JSON(http.StatusCreated, dto.AuthResponseDTO{Token: token})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	token, _ := h.jwtService.GenerateJWT(user.Email)
	c.JSON(http.StatusOK, dto.AuthResponseDTO{Token: token})
}