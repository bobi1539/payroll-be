package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
)

type AllowanceTypeRepository interface {
	Create(allowanceType *domain.AllowanceType) *domain.AllowanceType
	Update(allowanceType *domain.AllowanceType) *domain.AllowanceType
	Delete(id int64)
	FindById(id int64) (*domain.AllowanceType, error)
	FindAll(search *dto.Search) []domain.AllowanceType
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.AllowanceType
	FindTotalItem(search *dto.Search) int64
}
