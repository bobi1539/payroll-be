package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/model/search"
)

type BasicSalaryService interface {
	Create(request *request.BasicSalaryRequest, header dto.Header) response.BasicSalaryResponse
	Update(id int64, request *request.BasicSalaryRequest, header dto.Header) response.BasicSalaryResponse
	FindById(id int64) response.BasicSalaryResponse
	FindByIdDomain(id int64) *domain.BasicSalary
	FindAll(search *search.BasicSalarySearch) []response.BasicSalaryResponse
	FindAllPagination(search *search.BasicSalarySearch, pagination *dto.Pagination) response.PaginationResponse
	Delete(id int64) response.BasicSalaryResponse
}
