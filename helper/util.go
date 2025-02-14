package helper

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var log *logrus.Logger

func GetLogger() *logrus.Logger {
	if log == nil {
		log = logrus.New()
	}
	return log
}

func StringQueryLike(value string) string {
	return "%" + strings.ToLower(value) + "%"
}

func StringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	PanicIfError(err)
	return intValue
}

func ContainsDigitUpperLower(s string) bool {
	hasDigit, hasUpper, hasLower := false, false, false

	for _, char := range s {
		if unicode.IsDigit(char) {
			hasDigit = true
		} else if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		}

		if hasDigit && hasUpper && hasLower {
			return true
		}
	}

	return false
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIfError(err)
	return string(hashedPassword)
}
