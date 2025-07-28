package repository

import (
	"errors"
	"manajemen-user/internal/domain/role"
	"manajemen-user/internal/domain/user"
	appErrors "manajemen-user/internal/errors"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetAllUsers() ([]user.User, error) {
	var userModels []UserModel
	if err := r.DB.Preload("Role").Find(&userModels).Error; err != nil {
		return nil, err
	}

	var domainUsers []user.User

	for _, dbUser := range userModels {

		domainUser := user.User{
			ID:        dbUser.ID,
			Name:      dbUser.Name,
			Email:     dbUser.Email,
			Password:  dbUser.Password,
			RoleID:    dbUser.RoleID,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,

			Role: role.Role{
				ID:          dbUser.Role.ID,
				Name:        dbUser.Role.Name,
				Deskription: dbUser.Role.Deskription,
				CreatedAt:   dbUser.Role.CreatedAt,
				UpdatedAt:   dbUser.Role.UpdatedAt,
			},
		}

		domainUsers = append(domainUsers, domainUser)
	}

	return domainUsers, nil
}

func (r *userRepository) GetUsersByID(id string) (*user.User, error) {
	var userModel UserModel
	if err := r.DB.Preload("Role").Where("ID = ?", id).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
	}

	domainUser := user.User{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		RoleID:    userModel.RoleID,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,

		Role: role.Role{
			ID:          userModel.Role.ID,
			Name:        userModel.Role.Name,
			Deskription: userModel.Role.Deskription,
			CreatedAt:   userModel.Role.CreatedAt,
			UpdatedAt:   userModel.Role.UpdatedAt,
		},
	}
	return &domainUser, nil
}

func (r *userRepository) CreateUsers(user *user.User) error {
	userModel := UserModel{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   user.RoleID,
	}

	if err := r.DB.Create(&userModel).Error; err != nil {
		return err
	}

	user.ID = userModel.ID

	return nil
}

func (r *userRepository) SaveUsers(user *user.User) error {
	return r.DB.Save(user).Error
}

func (r *userRepository) DeleteUsers(user *user.User) error {
	return r.DB.Delete(&user).Error
}

func (r *userRepository) FindByEmailWithRole(email string) (*user.User, error) {
	var userModel UserModel
	if err := r.DB.Preload("Role").Where("email = ?", email).First(&userModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrNotFound
		}
		return nil, err
	}

	domainUser := user.User{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		RoleID:    userModel.RoleID,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,

		Role: role.Role{
			ID:          userModel.Role.ID,
			Name:        userModel.Role.Name,
			Deskription: userModel.Role.Deskription,
			CreatedAt:   userModel.Role.CreatedAt,
			UpdatedAt:   userModel.Role.UpdatedAt,
		},
	}
	return &domainUser, nil
}
