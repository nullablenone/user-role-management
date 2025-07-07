package repository

import (
	"manajemen-user/internal/domain/role"

	"gorm.io/gorm"
)

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) role.Repository {
	return &roleRepository{DB: db}
}

func (r *roleRepository) GetAllRoles(user []role.Role) ([]role.Role, error) {
	if err := r.DB.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *roleRepository) GetRolesByID(id string) (*role.Role, error) {
	var user role.Role
	if err := r.DB.Where("ID = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *roleRepository) CreateRoles(user *role.Role) error {
	return r.DB.Create(user).Error
}

func (r *roleRepository) SaveRoles(user *role.Role) error {
	return r.DB.Save(user).Error
}

func (r *roleRepository) DeleteRoles(user *role.Role) error {
	return r.DB.Delete(user).Error
}
