package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/model/search"
)

type AllowanceService interface {
	Create(request *request.AllowanceRequest, header dto.Header) response.AllowanceResponse
	Update(id int64, request *request.AllowanceRequest, header dto.Header) response.AllowanceResponse
	FindById(id int64) response.AllowanceResponse
	FindByIdDomain(id int64) *domain.Allowance
	FindAll(search *search.AllowanceSearch) []response.AllowanceResponse
	FindAllPagination(search *search.AllowanceSearch, pagination *dto.Pagination) response.PaginationResponse
	Delete(id int64) response.AllowanceResponse
}
