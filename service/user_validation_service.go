package service

import "payroll/model/domain"

type UserValidationService interface {
	ValidateCreateUsername(username string)
	ValidateUpdateUsername(username string, userExisting *domain.User)
	ValidatePassword(password string)
}
