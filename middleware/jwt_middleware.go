package middleware

import (
	"errors"
	"payroll/constant"
	"payroll/endpoint"
	"payroll/exception"
	"payroll/model/dto"
	"payroll/service/serviceimpl"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware(ctx *fiber.Ctx) error {
	if shouldSkipMiddleware(ctx.Path()) {
		return ctx.Next()
	}

	jwtToken := validateAuthHeader(ctx.Get(constant.AUTHORIZATION))

	jwtService := serviceimpl.NewJwtServiceImpl()
	claims := jwtService.ExtractJwtClaims(jwtToken)

	header := dto.Header{
		UserId: claims.UserId,
	}
	ctx.Locals(constant.HEADER, header)

	return ctx.Next()
}

func shouldSkipMiddleware(path string) bool {
	skipPrefixes := []string{endpoint.AUTHS, endpoint.SWAGGER}

	for _, prefix := range skipPrefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}

func validateAuthHeader(authHeader string) string {
	if authHeader == "" || !strings.HasPrefix(authHeader, constant.BEARER) {
		panicUnauthorized()
	}
	return strings.TrimPrefix(authHeader, constant.BEARER)
}

func panicUnauthorized() {
	exception.PanicErrorBusiness(fiber.StatusUnauthorized, errors.New(constant.UNAUTHORIZED))
}
