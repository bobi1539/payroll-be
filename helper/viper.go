package helper

import (
	"payroll/constant"

	"github.com/spf13/viper"
)

var config *viper.Viper

func NewViper() *viper.Viper {
	if config == nil {
		config = viper.New()
	}

	config.SetConfigFile(constant.CONFIG_FILE)
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	PanicIfError(err)

	return config
}
