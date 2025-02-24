package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
)

type RoleRepository interface {
	Create(role *domain.Role) *domain.Role
	Update(role *domain.Role) *domain.Role
	Delete(id int64)
	FindById(id int64) (*domain.Role, error)
	FindAll(search *dto.Search) []domain.Role
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.Role
	FindTotalItem(search *dto.Search) int64
}
