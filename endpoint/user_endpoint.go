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

const USERS = constant.PREFIX_API + "/users"
const USERS_ALL = USERS + "/all"
const USERS_BY_ID = USERS + "/id/:" + constant.ID

func SetUserEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	userController := getUserController(db, validate)

	fiberApp.Post(USERS, userController.Create)
	fiberApp.Put(USERS_BY_ID, userController.Update)
	fiberApp.Get(USERS_BY_ID, userController.FindById)
	fiberApp.Get(USERS_ALL, userController.FindAll)
	fiberApp.Get(USERS, userController.FindAllPagination)
	fiberApp.Delete(USERS_BY_ID, userController.Delete)
}

func getUserController(db *gorm.DB, validate *validator.Validate) controller.RoleController {
	return controllerimpl.NewUserController(getUserService(db, validate))
}

func getUserService(db *gorm.DB, validate *validator.Validate) service.UserService {
	return serviceimpl.NewUserService(
		getUserRepository(db),
		validate,
		getRoleService(db, validate),
	)
}

func getUserRepository(db *gorm.DB) repository.UserRepository {
	return repositoryimpl.NewUserRepository(db)
}
