package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
)

type BasicSalaryRepository interface {
	Create(basicSalary *domain.BasicSalary) *domain.BasicSalary
	Update(basicSalary *domain.BasicSalary) *domain.BasicSalary
	Delete(id int64)
	FindById(id int64) (*domain.BasicSalary, error)
	FindAll() []domain.BasicSalary
	FindAllPagination(pagination *dto.Pagination) []domain.BasicSalary
	FindTotalItem() int64
}
