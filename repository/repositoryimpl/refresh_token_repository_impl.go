package repositoryimpl

import (
	"payroll/model/domain"
	"payroll/repository"

	"gorm.io/gorm"
)

type RefreshTokenRepositoryImpl struct {
	DB *gorm.DB
}

func NewRefreshTokenRepositoryImpl(db *gorm.DB) repository.RefreshTokenRepository {
	return &RefreshTokenRepositoryImpl{DB: db}
}

func (refreshTokenRepository *RefreshTokenRepositoryImpl) Create(refreshToken *domain.RefreshToken) *domain.RefreshToken {
	refreshTokenRepository.DB.Create(refreshToken)
	return refreshToken
}

func (refreshTokenRepository *RefreshTokenRepositoryImpl) FindByTokenAndValidityIsValid(token string) (*domain.RefreshToken, error) {
	return &domain.RefreshToken{}, nil
}
