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

const ALLOWANCES = constant.PREFIX_API + "/allowances"
const ALLOWANCES_ALL = ALLOWANCES + "/all"
const ALLOWANCES_BY_ID = ALLOWANCES + "/id/:" + constant.ID

func SetAllowanceEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	allowanceController := getAllowanceController(db, validate)

	fiberApp.Post(ALLOWANCES, allowanceController.Create)
	fiberApp.Put(ALLOWANCES_BY_ID, allowanceController.Update)
	fiberApp.Get(ALLOWANCES_BY_ID, allowanceController.FindById)
	fiberApp.Get(ALLOWANCES_ALL, allowanceController.FindAll)
	fiberApp.Get(ALLOWANCES, allowanceController.FindAllPagination)
	fiberApp.Delete(ALLOWANCES_BY_ID, allowanceController.Delete)
}

func getAllowanceController(db *gorm.DB, validate *validator.Validate) controller.AllowanceController {
	return controllerimpl.NewAllowanceControllerImpl(getAllowanceService(db, validate))
}

func getAllowanceService(db *gorm.DB, validate *validator.Validate) service.AllowanceService {
	return serviceimpl.NewAllowanceServiceImpl(
		getAllowanceRepository(db),
		getUserService(db, validate),
		getPositionService(db, validate),
		getAllowanceTypeService(db, validate),
		validate,
	)
}

func getAllowanceRepository(db *gorm.DB) repository.AllowanceRepository {
	return repositoryimpl.NewAllowanceRepositoryImpl(db)
}
