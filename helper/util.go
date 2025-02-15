package helper

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var log *logrus.Logger

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
	if len(value) == 0 {
		return 0
	}

	intValue, err := strconv.Atoi(value)
	PanicIfError(err)
	return intValue
}

func StringToInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}

	intValue, err := strconv.Atoi(value)
	PanicIfError(err)
	return int64(intValue)
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

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		PanicIfError(err)
		bytes[i] = charset[randomIndex.Int64()]
	}
	return string(bytes)
}
