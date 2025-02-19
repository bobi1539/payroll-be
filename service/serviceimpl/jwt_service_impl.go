package serviceimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var config = helper.NewViper()

type JwtServiceImpl struct {
}

func NewJwtServiceImpl() service.JwtService {
	return &JwtServiceImpl{}
}

func (jwtService *JwtServiceImpl) GenerateJwtToken(user *domain.User) string {
	jwtClaims := buildJwtClaims(user)
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims).SignedString(getJwtSignatureKey())
	helper.PanicIfError(err)
	return jwtToken
}

func (jwtService *JwtServiceImpl) ExtractJwtClaims(jwtToken string) dto.JwtClaims {
	token, err := jwt.ParseWithClaims(jwtToken, &dto.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getJwtSignatureKey(), nil
	})

	if err != nil {
		exception.PanicErrorBusiness(fiber.StatusUnauthorized, errors.New(constant.UNAUTHORIZED))
	}

	claims, ok := token.Claims.(*dto.JwtClaims)
	if !ok || !token.Valid {
		exception.PanicErrorBusiness(fiber.StatusUnauthorized, errors.New(constant.UNAUTHORIZED))
	}

	return *claims
}

func buildJwtClaims(user *domain.User) dto.JwtClaims {
	return dto.JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.GetString(constant.APP_NAME),
			ExpiresAt: getJwtExpired(),
		},
		UserId: user.Id,
	}
}

func getJwtExpired() *jwt.NumericDate {
	jwtExpiredInHour := time.Duration(helper.StringToInt(config.GetString(constant.JWT_EXPIRED)))
	return jwt.NewNumericDate(time.Now().Add(jwtExpiredInHour * time.Hour))
}

func getJwtSignatureKey() []byte {
	return []byte(config.GetString(constant.JWT_KEY))
}
