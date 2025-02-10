package repositoryimpl

import (
	"errors"
	"payroll/constant"
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

func (repository *RoleRepositoryImpl) Create(role *domain.Role) *domain.Role {
	repository.DB.Create(role)
	return role
}

func (repository *RoleRepositoryImpl) Update(role *domain.Role) *domain.Role {
	repository.DB.Save(role)
	return role
}

func (repository *RoleRepositoryImpl) FindById(id int64) (*domain.Role, error) {
	role := &domain.Role{}
	result := repository.DB.First(&role, "id = ? ", id)

	if result.Error == nil {
		return role, nil
	}
	return role, errors.New(constant.DATA_NOT_FOUND)
}
