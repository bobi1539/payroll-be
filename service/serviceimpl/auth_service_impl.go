package serviceimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/helper"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/repository"
	"payroll/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	JwtService     service.JwtService
	Validate       *validator.Validate
}

func NewAuthServiceImpl(
	userRepository repository.UserRepository,
	jwtService service.JwtService,
	validate *validator.Validate,
) service.AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		JwtService:     jwtService,
		Validate:       validate,
	}
}

func (authService *AuthServiceImpl) Login(request *request.LoginRequest) response.LoginResponse {
	err := authService.Validate.Struct(request)
	exception.PanicErrorBusiness(fiber.StatusBadRequest, err)

	user, err := authService.UserRepository.FindByUsername(request.Username)
	if err != nil {
		wrongUsernamePassword()
	}

	validatePassword(user.Password, request.Password)

	jwtToken := authService.JwtService.GenerateJwtToken(user)
	refreshToken := helper.GenerateRandomString(50)
	return buildLoginResponse(jwtToken, refreshToken)
}

func validatePassword(hashPassword string, password string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		wrongUsernamePassword()
	}
}

func wrongUsernamePassword() {
	exception.PanicErrorBusiness(fiber.StatusUnauthorized, errors.New(constant.WRONG_USERNAME_PASSWORD))
}

func buildLoginResponse(jwtToken string, refrestToken string) response.LoginResponse {
	return response.LoginResponse{
		JwtToken:     jwtToken,
		RefreshToken: refrestToken,
	}
}
