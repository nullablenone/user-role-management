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

func (r *roleRepository) GetAllRoles() ([]role.Role, error) {
	var roleModels []RoleModel
	if err := r.DB.Find(&roleModels).Error; err != nil {
		return nil, err
	}

	var roles []role.Role

	for _, item := range roleModels {
		domainRole := role.Role{
			ID:          item.ID,
			Name:        item.Name,
			Deskription: item.Deskription,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}

		roles = append(roles, domainRole)
	}

	return roles, nil
}

func (r *roleRepository) GetRolesByID(id string) (*role.Role, error) {
	var roleModel RoleModel
	if err := r.DB.Where("ID = ?", id).First(&roleModel).Error; err != nil {
		return nil, err
	}

	domainRole := role.Role{
		ID:          roleModel.ID,
		Name:        roleModel.Name,
		Deskription: roleModel.Deskription,
		CreatedAt:   roleModel.CreatedAt,
		UpdatedAt:   roleModel.UpdatedAt,
	}
	return &domainRole, nil
}

func (r *roleRepository) CreateRoles(role *role.Role) error {
	roleModel := RoleModel{
		Name:        role.Name,
		Deskription: role.Deskription,
	}

	if err := r.DB.Create(&roleModel).Error; err != nil {
		return err
	}
	role.ID = roleModel.ID

	return nil
}

func (r *roleRepository) SaveRoles(role *role.Role) error {
	return r.DB.Save(role).Error
}

func (r *roleRepository) DeleteRoles(role *role.Role) error {
	return r.DB.Delete(role).Error
}
