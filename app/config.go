package app

import (
	"payroll/constant"
	"payroll/helper"

	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()
	config.SetConfigFile(constant.CONFIG_FILE)
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	helper.PanicIfError(err)

	return config
}
