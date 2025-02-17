package serviceimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/request"
	"payroll/model/response"
	"payroll/repository"
	"payroll/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository         repository.UserRepository
	RefreshTokenRepository repository.RefreshTokenRepository
	JwtService             service.JwtService
	Validate               *validator.Validate
}

func NewAuthServiceImpl(
	userRepository repository.UserRepository,
	refreshTokenRepository repository.RefreshTokenRepository,
	jwtService service.JwtService,
	validate *validator.Validate,
) service.AuthService {
	return &AuthServiceImpl{
		UserRepository:         userRepository,
		RefreshTokenRepository: refreshTokenRepository,
		JwtService:             jwtService,
		Validate:               validate,
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
	refreshToken := authService.createRefreshToken(user)

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

func (authService *AuthServiceImpl) createRefreshToken(user *domain.User) string {
	token := helper.GenerateRandomString(50)

	refreshToken := &domain.RefreshToken{
		Token:    token,
		Validity: time.Now().AddDate(0, 1, 0),
		User:     user,
	}
	authService.RefreshTokenRepository.Create(refreshToken)

	return token
}

func buildLoginResponse(jwtToken string, refrestToken string) response.LoginResponse {
	return response.LoginResponse{
		JwtToken:     jwtToken,
		RefreshToken: refrestToken,
	}
}
