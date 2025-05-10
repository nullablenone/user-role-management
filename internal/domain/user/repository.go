package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllUsers(user []User) ([]User, error)
	GetUsersByID(id string) (*User, error)
	CreateUsers(user *User) error
	SaveUsers(user *User) error
	DeleteUsers(user *User) error
	FindByEmailWithRole(email string) (*User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllUsers(user []User) ([]User, error) {
	if err := r.DB.Preload("Role").Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetUsersByID(id string) (*User, error) {
	var user User
	if err := r.DB.Preload("Role").Where("ID = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) CreateUsers(user *User) error {
	return r.DB.Create(user).Error
}

func (r *repository) SaveUsers(user *User) error {
	return r.DB.Save(user).Error
}

func (r *repository) DeleteUsers(user *User) error {
	return r.DB.Delete(user).Error
}

func (r *repository) FindByEmailWithRole(email string) (*User, error) {
	var user User
	if err := r.DB.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
