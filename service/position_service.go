package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
)

type PositionService interface {
	Create(request *request.PositionRequest, header dto.Header) response.PositionResponse
	Update(id int64, request *request.PositionRequest, header dto.Header) response.PositionResponse
	FindById(id int64) response.PositionResponse
	FindByIdDomain(id int64) *domain.Position
	FindAll(search *dto.Search) []response.PositionResponse
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) response.PaginationResponse
	Delete(id int64) response.PositionResponse
}
