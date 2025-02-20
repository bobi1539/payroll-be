package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
)

type AllowanceRepository interface {
	Create(allowance *domain.Allowance) *domain.Allowance
	Update(allowance *domain.Allowance) *domain.Allowance
	Delete(id int64)
	FindById(id int64) (*domain.Allowance, error)
	FindAll(search *dto.Search) []domain.Allowance
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.Allowance
	FindTotalItem() int64
}
