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

const AUTHS = constant.PREFIX_API + "/auths"
const AUTHS_LOGIN = AUTHS + "/login"
const AUTHS_REFRESH_TOKEN = AUTHS + "/refresh-token"

func SetAuthEndpoint(fiberApp *fiber.App, db *gorm.DB, validate *validator.Validate) {
	authController := getAuthController(db, validate)

	fiberApp.Post(AUTHS_LOGIN, authController.Login)
	fiberApp.Post(AUTHS_REFRESH_TOKEN, authController.LoginRefreshToken)
}

func getAuthController(db *gorm.DB, validate *validator.Validate) controller.AuthController {
	return controllerimpl.NewAuthControllerImpl(getAuthService(db, validate))
}

func getAuthService(db *gorm.DB, validate *validator.Validate) service.AuthService {
	return serviceimpl.NewAuthServiceImpl(
		getUserRepository(db),
		getRefreshTokenRepository(db),
		getJwtService(),
		validate,
	)
}

func getJwtService() service.JwtService {
	return serviceimpl.NewJwtServiceImpl()
}

func getRefreshTokenRepository(db *gorm.DB) repository.RefreshTokenRepository {
	return repositoryimpl.NewRefreshTokenRepositoryImpl(db)
}
