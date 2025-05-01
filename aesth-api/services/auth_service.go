package services

import (
	"aesth-api/models"
	"aesth-api/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo}
}

func (s *AuthService) Login(email string, password string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) Register(email string, password string) (*models.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email: email,
		Password: string(hashed),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
