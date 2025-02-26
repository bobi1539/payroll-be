package serviceimpl

import (
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/model/search"
	"payroll/repository"
	"payroll/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
	UserService        service.UserService
	PositionService    service.PositionService
	Validate           *validator.Validate
}

func NewEmployeeServiceImpl(
	employeeRepository repository.EmployeeRepository,
	userService service.UserService,
	positionService service.PositionService,
	validate *validator.Validate,
) service.EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		UserService:        userService,
		PositionService:    positionService,
		Validate:           validate,
	}
}

func (employeeService *EmployeeServiceImpl) Create(request *request.EmployeeRequest, header dto.Header) response.EmployeeResponse {
	employeeService.validateRequest(request)
	user := employeeService.UserService.FindByIdDomain(header.UserId)

	employee := &domain.Employee{}
	employeeService.setEmployee(employee, request)
	helper.SetCreated(&employee.BaseDomain, user)
	helper.SetUpdated(&employee.BaseDomain, user)

	return response.ToEmployeeResponse(employeeService.EmployeeRepository.Create(employee))
}

func (employeeService *EmployeeServiceImpl) Update(id int64, request *request.EmployeeRequest, header dto.Header) response.EmployeeResponse {
	employeeService.validateRequest(request)
	user := employeeService.UserService.FindByIdDomain(header.UserId)

	employee := employeeService.FindByIdDomain(id)
	employeeService.setEmployee(employee, request)
	helper.SetUpdated(&employee.BaseDomain, user)

	return response.ToEmployeeResponse(employeeService.EmployeeRepository.Update(employee))
}

func (employeeService *EmployeeServiceImpl) FindById(id int64) response.EmployeeResponse {
	employee := employeeService.FindByIdDomain(id)
	return response.ToEmployeeResponse(employee)
}

func (employeeService *EmployeeServiceImpl) FindByIdDomain(id int64) *domain.Employee {
	employee, err := employeeService.EmployeeRepository.FindById(id)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
	return employee
}

func (employeeService *EmployeeServiceImpl) FindAll(search *search.EmployeeSearch) []response.EmployeeResponse {
	employees := employeeService.EmployeeRepository.FindAll(search)
	return response.ToEmployeeResponses(employees)
}

func (employeeService *EmployeeServiceImpl) FindAllPagination(search *search.EmployeeSearch, pagination *dto.Pagination) response.PaginationResponse {
	employees := employeeService.EmployeeRepository.FindAllPagination(search, pagination)
	totalItem := employeeService.EmployeeRepository.FindTotalItem(search)

	responses := response.ToEmployeeResponses(employees)
	return response.ToPaginationResponse(responses, pagination.PageNumber, pagination.PageSize, totalItem)
}

func (employeeService *EmployeeServiceImpl) Delete(id int64) response.EmployeeResponse {
	employee := employeeService.FindByIdDomain(id)
	employeeService.EmployeeRepository.Delete(id)
	return response.ToEmployeeResponse(employee)
}

func (employeeService *EmployeeServiceImpl) validateRequest(request *request.EmployeeRequest) {
	err := employeeService.Validate.Struct(request)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)
}

func (employeeService *EmployeeServiceImpl) setEmployee(employee *domain.Employee, request *request.EmployeeRequest) {
	employee.Name = request.Name
	employee.PhoneNumber = request.PhoneNumber
	employee.Email = request.Email
	employee.Address = request.Address
	employee.WorkStatus = request.WorkStatus
	employee.BankAccountNumber = request.BankAccountNumber
	employee.BankAccountName = request.BankAccountName
	employee.Npwp = request.Npwp
	employee.DateOfBirth = helper.FromStringToTime(request.DateOfBirth)
	employee.JoinDate = helper.FromStringToTime(request.JoinDate)
	employee.IsMarried = *request.IsMarried
	employee.TotalChild = *request.TotalChild
	employee.Position = employeeService.PositionService.FindByIdDomain(request.PositionId)
}
