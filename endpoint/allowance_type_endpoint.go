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

const ALLOWANCE_TYPES = constant.PREFIX_API + "/allowance-types"
const ALLOWANCE_TYPES_ALL = ALLOWANCE_TYPES + "/all"
const ALLOWANCE_TYPES_BY_ID = ALLOWANCE_TYPES + "/id/:" + constant.ID

func SetAllowanceTypeEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	atController := getAllowanceTypeController(db, validate)

	fiberApp.Post(ALLOWANCE_TYPES, atController.Create)
	fiberApp.Put(ALLOWANCE_TYPES_BY_ID, atController.Update)
	fiberApp.Get(ALLOWANCE_TYPES_BY_ID, atController.FindById)
	fiberApp.Get(ALLOWANCE_TYPES_ALL, atController.FindAll)
	fiberApp.Get(ALLOWANCE_TYPES, atController.FindAllPagination)
	fiberApp.Delete(ALLOWANCE_TYPES_BY_ID, atController.Delete)
}

func getAllowanceTypeController(db *gorm.DB, validate *validator.Validate) controller.AllowanceTypeController {
	return controllerimpl.NewAllowanceTypeControllerImpl(getAllowanceTypeService(db, validate))
}

func getAllowanceTypeService(db *gorm.DB, validate *validator.Validate) service.AllowanceTypeService {
	return serviceimpl.NewAllowanceTypeServiceImpl(
		getAllowanceTypeRepository(db),
		getUserService(db, validate),
		validate,
	)
}

func getAllowanceTypeRepository(db *gorm.DB) repository.AllowanceTypeRepository {
	return repositoryimpl.NewAllowanceTypeRepositoryImpl(db)
}
