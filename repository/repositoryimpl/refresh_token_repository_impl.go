package repositoryimpl

import (
	"errors"
	"payroll/constant"
	"payroll/model/domain"
	"payroll/repository"
	"time"

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
	refreshToken := &domain.RefreshToken{}

	result := refreshTokenRepository.DB.
		Preload(domain.USER).
		Where("token = ?", token).
		Where("validity > ?", time.Now()).
		First(refreshToken)

	if result.Error != nil {
		return refreshToken, errors.New(constant.DATA_NOT_FOUND)
	}
	return refreshToken, nil
}
