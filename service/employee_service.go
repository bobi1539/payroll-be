package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/model/search"
)

type EmployeeService interface {
	Create(request *request.EmployeeRequest, header dto.Header) response.EmployeeResponse
	Update(id int64, request *request.EmployeeRequest, header dto.Header) response.EmployeeResponse
	FindById(id int64) response.EmployeeResponse
	FindByIdDomain(id int64) *domain.Employee
	FindAll(search *search.EmployeeSearch) []response.EmployeeResponse
	FindAllPagination(search *search.EmployeeSearch, pagination *dto.Pagination) response.PaginationResponse
	Delete(id int64) response.EmployeeResponse
}
