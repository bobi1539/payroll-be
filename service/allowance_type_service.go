package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
)

type AllowanceTypeService interface {
	Create(request *request.AllowanceTypeRequest, header dto.Header) response.AllowanceTypeResponse
	Update(id int64, request *request.AllowanceTypeRequest, header dto.Header) response.AllowanceTypeResponse
	FindById(id int64) response.AllowanceTypeResponse
	FindByIdDomain(id int64) *domain.AllowanceType
	FindAll(search *dto.Search) []response.AllowanceTypeResponse
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse
	Delete(id int64) response.AllowanceTypeResponse
}
