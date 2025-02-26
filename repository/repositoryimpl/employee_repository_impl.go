package repositoryimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/search"
	"payroll/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	DB *gorm.DB
}

func NewEmployeeRepositoryImpl(db *gorm.DB) repository.EmployeeRepository {
	return &EmployeeRepositoryImpl{DB: db}
}

func (employeeRepository *EmployeeRepositoryImpl) Create(employee *domain.Employee) *domain.Employee {
	employeeRepository.DB.Create(employee)
	return employee
}

func (employeeRepository *EmployeeRepositoryImpl) Update(employee *domain.Employee) *domain.Employee {
	employeeRepository.DB.Save(employee)
	return employee
}

func (employeeRepository *EmployeeRepositoryImpl) Delete(id int64) {
	result := employeeRepository.DB.Where("id = ?", id).Delete(&domain.Employee{})
	if result.Error != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.CANNOT_DELETE_THIS_DATA))
	}
}

func (employeeRepository *EmployeeRepositoryImpl) FindById(id int64) (*domain.Employee, error) {
	employee := &domain.Employee{}
	result := employeeRepository.DB.
		Preload(domain.POSITION).
		First(employee, "id = ?", id)

	if result.Error != nil {
		return nil, errors.New(constant.EMPLOYEE_NOT_FOUND)
	}
	return employee, nil
}

func (employeeRepository *EmployeeRepositoryImpl) FindAll(search *search.EmployeeSearch) []domain.Employee {
	var employees []domain.Employee
	employeeRepository.DB.
		Preload(domain.POSITION).
		Where(employeeRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Find(&employees)
	return employees
}

func (employeeRepository *EmployeeRepositoryImpl) FindAllPagination(search *search.EmployeeSearch, pagination *dto.Pagination) []domain.Employee {
	var employees []domain.Employee
	employeeRepository.DB.
		Preload(domain.POSITION).
		Where(employeeRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&employees)
	return employees
}

func (employeeRepository *EmployeeRepositoryImpl) FindTotalItem(search *search.EmployeeSearch) int64 {
	var totalItem int64
	employeeRepository.DB.
		Model(&domain.Employee{}).
		Where(employeeRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Count(&totalItem)
	return totalItem
}

func (employeeRepository *EmployeeRepositoryImpl) searchLike() string {
	return "LOWER(name) LIKE ?"
}
