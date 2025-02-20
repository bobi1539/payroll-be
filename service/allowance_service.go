package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
)

type AllowanceService interface {
	Create(request *request.AllowanceRequest, header dto.Header) response.AllowanceResponse
	Update(id int64, request *request.AllowanceRequest, header dto.Header) response.AllowanceResponse
	FindById(id int64) response.AllowanceResponse
	FindByIdDomain(id int64) *domain.Allowance
	FindAll(search *dto.Search) []response.AllowanceResponse
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse
	Delete(id int64) response.AllowanceResponse
}
