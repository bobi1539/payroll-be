package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
)

type RoleService interface {
	Create(request *request.RoleRequest, header dto.Header) response.RoleResponse
	Update(id int64, request *request.RoleRequest, header dto.Header) response.RoleResponse
	FindById(id int64) response.RoleResponse
	FindByIdDomain(id int64) *domain.Role
	FindAll(search *dto.Search) []response.RoleResponse
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse
	Delete(id int64) response.RoleResponse
}
