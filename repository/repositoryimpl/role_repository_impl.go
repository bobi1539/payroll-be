package repositoryimpl

import (
	"errors"
	"payroll/constant"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
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

func (roleRepository *RoleRepositoryImpl) Update(role *domain.Role) *domain.Role {
	roleRepository.DB.Save(role)
	return role
}

func (roleRepository *RoleRepositoryImpl) FindById(id int64) (*domain.Role, error) {
	role := &domain.Role{}
	result := roleRepository.DB.
		Where(constant.IS_DELETED_FALSE).
		First(&role, "id = ?", id)

	if result.Error == nil {
		return role, nil
	}
	return role, errors.New(constant.DATA_NOT_FOUND)
}

func (roleRepository *RoleRepositoryImpl) FindAll(search *dto.Search) []domain.Role {
	var roles []domain.Role
	roleRepository.DB.
		Where("LOWER(name) LIKE ?", helper.StringQueryLike(search.Value)).
		Where(constant.IS_DELETED_FALSE).
		Order("id ASC").
		Find(&roles)
	return roles
}

func (roleRepository *RoleRepositoryImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.Role {
	var roles []domain.Role
	roleRepository.DB.
		Where("LOWER(name) LIKE ?", helper.StringQueryLike(search.Value)).
		Where(constant.IS_DELETED_FALSE).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&roles)
	return roles
}

func (roleRepository *RoleRepositoryImpl) FindTotalItem() int64 {
	var totalItem int64
	roleRepository.DB.
		Model(&domain.Role{}).
		Where(constant.IS_DELETED_FALSE).
		Count(&totalItem)
	return totalItem
}
