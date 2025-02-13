package service

type UserValidationService interface {
	ValidateCreateUsername(username string)
}
