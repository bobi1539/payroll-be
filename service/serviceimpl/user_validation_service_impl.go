package serviceimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/repository"
	"payroll/service"

	"github.com/gofiber/fiber/v2"
)

type UserValidationServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserValidationServiceImpl(userRepository repository.UserRepository) service.UserValidationService {
	return &UserValidationServiceImpl{
		UserRepository: userRepository,
	}
}

func (userValidationService *UserValidationServiceImpl) ValidateCreateUsername(username string) {
	user, _ := userValidationService.UserRepository.FindByUsername(username)
	if user != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.USERNAME_ALREADY_REGISTERED))
	}
}

func (userValidationService *UserValidationServiceImpl) ValidateUpdateUsername(username string, userExisting *domain.User) {
	user, _ := userValidationService.UserRepository.FindByUsername(username)
	if user != nil && user.ID != userExisting.ID {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.USERNAME_ALREADY_REGISTERED))
	}
}

func (userValidationService *UserValidationServiceImpl) ValidatePassword(password string) {
	if len(password) < constant.MIN_PASSWORD_LEN {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.PASSWORD_MIN_10))
	}
	if !helper.ContainsDigitUpperLower(password) {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.PASSWORD_DIGIT_UPPER_LOWER))
	}
}
