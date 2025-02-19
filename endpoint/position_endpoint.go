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

const POSITIONS = constant.PREFIX_API + "/positions"
const POSITIONS_ALL = POSITIONS + "/all"
const POSITIONS_BY_ID = POSITIONS + "/id/:" + constant.ID

func SetPositionEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	positionController := getPositionController(db, validate)

	fiberApp.Post(POSITIONS, positionController.Create)
	fiberApp.Put(POSITIONS_BY_ID, positionController.Update)
	fiberApp.Get(POSITIONS_BY_ID, positionController.FindById)
	fiberApp.Get(POSITIONS_ALL, positionController.FindAll)
	fiberApp.Get(POSITIONS, positionController.FindAllPagination)
	fiberApp.Delete(POSITIONS_BY_ID, positionController.Delete)
}

func getPositionController(db *gorm.DB, validate *validator.Validate) controller.PositionController {
	return controllerimpl.NewPositionControllerImpl(getPositionService(db, validate))
}

func getPositionService(db *gorm.DB, validate *validator.Validate) service.PositionService {
	return serviceimpl.NewPositionServiceImpl(
		getPositionRepository(db),
		getUserRepository(db),
		validate,
	)
}

func getPositionRepository(db *gorm.DB) repository.PositionRepository {
	return repositoryimpl.NewPositionRepositoryImpl(db)
}
