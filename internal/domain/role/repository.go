package role

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllRoles(user []Role) ([]Role, error)
	GetRolesByID(id string) (*Role, error)
	CreateRoles(user *Role) error
	SaveRoles(user *Role) error
	DeleteRoles(user *Role) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) GetAllRoles(user []Role) ([]Role, error) {
	if err := r.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetRolesByID(id string) (*Role, error) {
	var user Role
	if err := r.DB.Where("ID = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) CreateRoles(user *Role) error {
	return r.DB.Create(user).Error
}

func (r *repository) SaveRoles(user *Role) error {
	return r.DB.Save(user).Error
}

func (r *repository) DeleteRoles(user *Role) error {
	return r.DB.Delete(user).Error
}
