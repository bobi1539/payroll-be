package endpoint

import (
	"payroll/constant"
	"payroll/controller"
	"payroll/controller/controllerimpl"
	"payroll/repository"
	"payroll/repository/repositoryimpl"
	"payroll/service"
	"payroll/service/serviceimpl"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const EMPLOYEES = constant.PREFIX_API + "/employees"
const EMPLOYEES_ALL = EMPLOYEES + "/all"
const EMPLOYEES_BY_ID = EMPLOYEES + "/id/:" + constant.ID

func SetEmployeeEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	employeeController := geEmployeeController(db, validate)

	fiberApp.Post(EMPLOYEES, employeeController.Create)
	fiberApp.Put(EMPLOYEES_BY_ID, employeeController.Update)
	fiberApp.Get(EMPLOYEES_BY_ID, employeeController.FindById)
	fiberApp.Get(EMPLOYEES_ALL, employeeController.FindAll)
	fiberApp.Get(EMPLOYEES, employeeController.FindAllPagination)
	fiberApp.Delete(EMPLOYEES_BY_ID, employeeController.Delete)
}

func geEmployeeController(db *gorm.DB, validate *validator.Validate) controller.EmployeeController {
	return controllerimpl.NewEmployeeControllerImpl(getEmployeeService(db, validate))
}

func getEmployeeService(db *gorm.DB, validate *validator.Validate) service.EmployeeService {
	return serviceimpl.NewEmployeeServiceImpl(
		getEmployeeRepository(db),
		getUserService(db, validate),
		getPositionService(db, validate),
		validate,
	)
}

func getEmployeeRepository(db *gorm.DB) repository.EmployeeRepository {
	return repositoryimpl.NewEmployeeRepositoryImpl(db)
}
