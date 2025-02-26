package search

import "payroll/model/dto"

type EmployeeSearch struct {
	dto.Search
}

func BuildEmployeeSearch(value string) EmployeeSearch {
	return EmployeeSearch{
		Search: dto.Search{
			Value: value,
		},
	}
}
