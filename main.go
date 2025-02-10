package main

import (
	"fmt"
	"payroll/app"
	"payroll/model/domain"
	"payroll/repository/repositoryimpl"
)

func main() {
	db := app.NewDB()

	role := &domain.Role{
		Name: "admin",
	}

	roleRepository := repositoryimpl.NewRoleRepository(db)
	roleRepository.Create(role)
	fmt.Println("role : ", role)
}
