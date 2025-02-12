package helper

import (
	"strings"

	"github.com/sirupsen/logrus"
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
