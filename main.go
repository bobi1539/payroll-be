package main

import (
	"fmt"
	"payroll/app"
	"payroll/repository/repositoryimpl"
	"payroll/service/serviceimpl"

	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()
	db := app.NewDB()
	roleRepository := repositoryimpl.NewRoleRepository(db)

	roleService := serviceimpl.NewRoleServiceImpl(roleRepository, validate)
	response := roleService.FindById(2)

	fmt.Println("response : ", response)
}
