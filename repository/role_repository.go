package repository

import "payroll/model/domain"

type RoleRepository interface {
	Create(role *domain.Role) *domain.Role
}
