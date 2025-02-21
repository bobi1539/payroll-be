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
	config.AutomaticEnv()

	config.BindEnv(constant.APP_NAME, constant.APP_NAME_ENV_NAME)
	config.BindEnv(constant.DB_HOST, constant.DB_HOST_ENV_NAME)
	config.BindEnv(constant.DB_PORT, constant.DB_PORT_ENV_NAME)
	config.BindEnv(constant.DB_USERNAME, constant.DB_USERNAME_ENV_NAME)
	config.BindEnv(constant.DB_PASSWORD, constant.DB_PASSWORD_ENV_NAME)
	config.BindEnv(constant.DB_NAME, constant.DB_NAME_ENV_NAME)
	config.BindEnv(constant.DB_SHOW_SQL, constant.DB_SHOW_SQL_ENV_NAME)
	config.BindEnv(constant.JWT_KEY, constant.JWT_KEY_ENV_NAME)
	config.BindEnv(constant.JWT_EXPIRED, constant.JWT_EXPIRED_ENV_NAME)

	err := config.ReadInConfig()
	PanicIfError(err)

	return config
}
