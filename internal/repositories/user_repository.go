package repositories

import (
	"gorm.io/gorm"

	"go-api/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Save(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := repo.db.Find(&users).Error
	return users, err
}

func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user *models.User
	err := repo.db.First(&user, id).Error
	return user, err
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) DeleteUser(id uint) error {
	return repo.db.Delete(&models.User{}, id).Error
}
