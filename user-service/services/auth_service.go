package services

import (
	"errors"
	"gorm.io/gorm"
	"strings"
	"user-service/models"
	"user-service/repositories"
	"user-service/utils"
)

type AuthService interface {
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{userRepo: repo}
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashedPassword, hashErr := utils.HashPassword(password)
			if hashErr != nil {
				return "", errors.New("failed to hash password")
			}

			newUser := &models.User{
				Name:     strings.Split(email, "@")[0],
				Email:    email,
				Password: hashedPassword,
			}

			if createErr := s.userRepo.Create(newUser); createErr != nil {
				return "", errors.New("failed to create user")
			}

			user = newUser
		} else {
			return "", errors.New("invalid credentials")
		}
	}

	if err := utils.VerifyPassword(user.Password, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateJWT(user.ID)
}
