package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/search"
)

type BasicSalaryRepository interface {
	Create(basicSalary *domain.BasicSalary) *domain.BasicSalary
	Update(basicSalary *domain.BasicSalary) *domain.BasicSalary
	Delete(id int64)
	FindById(id int64) (*domain.BasicSalary, error)
	FindAll(search *search.BasicSalarySearch) []domain.BasicSalary
	FindAllPagination(search *search.BasicSalarySearch, pagination *dto.Pagination) []domain.BasicSalary
	FindTotalItem(search *search.BasicSalarySearch) int64
}
