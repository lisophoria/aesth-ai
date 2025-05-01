package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"aesth-api/dto"
	"aesth-api/repositories"
)

type UserHandler struct {
	repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{repo}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.repo.GetMany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []dto.UserResponseDTO
	for _, u := range users {
		result = append(result, dto.UserResponseDTO{ID: u.ID, Email: u.Email})
	}
	c.JSON(http.StatusOK, result)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, dto.UserResponseDTO{ID: user.ID, Email: user.Email})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var input dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		user.Password = string(hashed)
	}

	if err := h.repo.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponseDTO{ID: user.ID, Email: user.Email})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}