package app

import (
	"fmt"
	"payroll/constant"
	"payroll/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := getDataSourceName()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}

func getDataSourceName() string {
	config := helper.NewViper()
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.GetString(constant.DB_HOST),
		config.GetString(constant.DB_USERNAME),
		config.GetString(constant.DB_PASSWORD),
		config.GetString(constant.DB_NAME),
		config.GetString(constant.DB_PORT),
	)
}
