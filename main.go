package main

import (
	"fmt"
	"payroll/app"
)

func main() {
	db := app.NewDB()

	if db != nil {
		fmt.Println("Hello")
	}
}
