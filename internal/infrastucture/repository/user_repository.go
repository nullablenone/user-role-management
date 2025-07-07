package repository

import (
	"manajemen-user/internal/domain/user"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetAllUsers(user []user.User) ([]user.User, error) {
	if err := r.DB.Preload("Role").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUsersByID(id string) (*user.User, error) {
	var user user.User
	if err := r.DB.Preload("Role").Where("ID = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUsers(user *user.User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) SaveUsers(user *user.User) error {
	return r.DB.Save(user).Error
}

func (r *userRepository) DeleteUsers(user *user.User) error {
	return r.DB.Delete(user).Error
}

func (r *userRepository) FindByEmailWithRole(email string) (*user.User, error) {
	var user user.User
	if err := r.DB.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
