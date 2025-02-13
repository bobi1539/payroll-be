package dto

import "payroll/helper"

type Pagination struct {
	PageNumber int
	PageSize   int
}

func BuildPagination(pageNumber string, pageSize string) Pagination {
	number := helper.StringToInt(pageNumber) - 1
	size := helper.StringToInt(pageSize)
	return Pagination{
		PageNumber: number * size,
		PageSize:   size,
	}
}
