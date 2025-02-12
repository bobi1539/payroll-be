package helper

import "github.com/sirupsen/logrus"

var log *logrus.Logger

func GetLogger() *logrus.Logger {
	if log == nil {
		log = logrus.New()
	}
	return log
}
