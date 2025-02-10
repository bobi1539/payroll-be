package repositoryimpl

import (
	"payroll/model/domain"
	"payroll/repository"

	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &RoleRepositoryImpl{DB: db}
}

func (roleRepository *RoleRepositoryImpl) Create(role *domain.Role) *domain.Role {
	roleRepository.DB.Create(role)
	return role
}
