package repository

import "payroll/model/domain"

type RefreshTokenRepository interface {
	Create(refreshToken *domain.RefreshToken) *domain.RefreshToken
	FindByTokenAndValidityIsValid(token string) (*domain.RefreshToken, error)
}
