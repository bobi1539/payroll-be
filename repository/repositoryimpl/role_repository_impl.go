package repositoryimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoleRepositoryImpl(db *gorm.DB) repository.RoleRepository {
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

func (roleRepository *RoleRepositoryImpl) Delete(id int64) {
	result := roleRepository.DB.Where("id = ?", id).Delete(&domain.Role{})
	if result.Error != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.CANNOT_DELETE_THIS_DATA))
	}
}

func (roleRepository *RoleRepositoryImpl) FindById(id int64) (*domain.Role, error) {
	role := &domain.Role{}
	result := roleRepository.DB.First(role, "id = ?", id)

	if result.Error != nil {
		return role, errors.New(constant.ROLE_NOT_FOUND)
	}
	return role, nil
}

func (roleRepository *RoleRepositoryImpl) FindAll(search *dto.Search) []domain.Role {
	var roles []domain.Role
	roleRepository.DB.
		Where(roleRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Find(&roles)
	return roles
}

func (roleRepository *RoleRepositoryImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.Role {
	var roles []domain.Role
	roleRepository.DB.
		Where(roleRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&roles)
	return roles
}

func (roleRepository *RoleRepositoryImpl) FindTotalItem(search *dto.Search) int64 {
	var totalItem int64
	roleRepository.DB.
		Model(&domain.Role{}).
		Where(roleRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Count(&totalItem)
	return totalItem
}

func (roleRepository *RoleRepositoryImpl) searchLike() string {
	return "LOWER(name) LIKE ?"
}
