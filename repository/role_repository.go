package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
)

type RoleRepository interface {
	Create(role *domain.Role) *domain.Role
	Update(role *domain.Role) *domain.Role
	FindById(id int64) (*domain.Role, error)
	FindAll(search *dto.Search) []domain.Role
}
