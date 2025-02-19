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

const BASIC_SALARIES = constant.PREFIX_API + "/basic-salaries"
const BASIC_SALARIES_ALL = BASIC_SALARIES + "/all"
const BASIC_SALARIES_BY_ID = BASIC_SALARIES + "/id/:" + constant.ID

func SetBasicSalaryEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	basicSalaryController := geBasicSalaryController(db, validate)

	fiberApp.Post(BASIC_SALARIES, basicSalaryController.Create)
	fiberApp.Put(BASIC_SALARIES_BY_ID, basicSalaryController.Update)
	fiberApp.Get(BASIC_SALARIES_BY_ID, basicSalaryController.FindById)
	fiberApp.Get(BASIC_SALARIES_ALL, basicSalaryController.FindAll)
	fiberApp.Get(BASIC_SALARIES, basicSalaryController.FindAllPagination)
	fiberApp.Delete(BASIC_SALARIES_BY_ID, basicSalaryController.Delete)
}

func geBasicSalaryController(db *gorm.DB, validate *validator.Validate) controller.BasicSalaryController {
	return controllerimpl.NewBasicSalaryControllerImpl(getBasicSalaryService(db, validate))
}

func getBasicSalaryService(db *gorm.DB, validate *validator.Validate) service.BasicSalaryService {
	return serviceimpl.NewBasicSalaryServiceImpl(
		getBasicSalaryRepository(db),
		getUserService(db, validate),
		getPositionService(db, validate),
		validate,
	)
}

func getBasicSalaryRepository(db *gorm.DB) repository.BasicSalaryRepository {
	return repositoryimpl.NewBasicSalaryRepositoryImpl(db)
}
