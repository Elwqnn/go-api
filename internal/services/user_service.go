package services

import (
	"golang.org/x/crypto/bcrypt"
	"strings"

	"go-api/internal/models"
	"go-api/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) Register(email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     strings.Split(email, "@")[0],
		Email:    email,
		Password: string(hashedPassword),
	}

	return service.repo.Save(user)
}

func (service *UserService) VerifyPassword(user *models.User, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func (service *UserService) CreateUser(user *models.User) error {
	return service.repo.CreateUser(user)
}

func (service *UserService) GetUsers() ([]models.User, error) {
	return service.repo.GetUsers()
}

func (service *UserService) GetUserByID(id uint) (*models.User, error) {
	return service.repo.GetUserByID(id)
}

func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	return service.repo.GetUserByEmail(email)
}

func (service *UserService) UpdateUser(user *models.User) error {
	return service.repo.UpdateUser(user)
}

func (service *UserService) DeleteUser(id uint) error {
	return service.repo.DeleteUser(id)
}
