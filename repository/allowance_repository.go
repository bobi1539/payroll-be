package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/search"
)

type AllowanceRepository interface {
	Create(allowance *domain.Allowance) *domain.Allowance
	Update(allowance *domain.Allowance) *domain.Allowance
	Delete(id int64)
	FindById(id int64) (*domain.Allowance, error)
	FindAll(search *search.AllowanceSearch) []domain.Allowance
	FindAllPagination(search *search.AllowanceSearch, pagination *dto.Pagination) []domain.Allowance
	FindTotalItem(search *search.AllowanceSearch) int64
}
