package service

import (
	"payroll/model/domain"
	"payroll/model/dto"
)

type JwtService interface {
	GenerateJwtToken(user *domain.User) string
	ExtractJwtClaims(jwtToken string) dto.JwtClaims
}
