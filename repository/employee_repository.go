package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/search"
)

type EmployeeRepository interface {
	Create(employee *domain.Employee) *domain.Employee
	Update(employee *domain.Employee) *domain.Employee
	Delete(id int64)
	FindById(id int64) (*domain.Employee, error)
	FindAll(search *search.EmployeeSearch) []domain.Employee
	FindAllPagination(search *search.EmployeeSearch, pagination *dto.Pagination) []domain.Employee
	FindTotalItem(search *search.EmployeeSearch) int64
}
