package service

import (
	"payroll/model/domain"
	"payroll/model/request"
	"payroll/model/response"
)

type RoleService interface {
	Create(request *request.RoleRequest) response.RoleResponse
	Update(id int64, request *request.RoleRequest) response.RoleResponse
	FindById(id int64) response.RoleResponse
	FindByIdDomain(id int64) *domain.Role
}
