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
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

const ROLES = constant.PREFIX_API + "/roles"
const ROLES_BY_ID = ROLES + "/:" + constant.ROLE_ID

func SetRoleEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	roleController := getRoleController(db, validate)

	fiberApp.Post(ROLES, roleController.Create)
	fiberApp.Put(ROLES, roleController.Update)
	fiberApp.Get(ROLES_BY_ID, roleController.FindById)
}

func getRoleController(db *gorm.DB, validate *validator.Validate) controller.RoleController {
	return controllerimpl.NewRoleController(getRoleService(db, validate))
}

func getRoleService(db *gorm.DB, validate *validator.Validate) service.RoleService {
	return serviceimpl.NewRoleService(getRoleRepository(db), validate)
}

func getRoleRepository(db *gorm.DB) repository.RoleRepository {
	return repositoryimpl.NewRoleRepository(db)
}
