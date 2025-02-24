package search

import "payroll/helper"

type BasicSalarySearch struct {
	PositionId int64
}

func BuildBasicSalarySearch(positionId string) BasicSalarySearch {
	return BasicSalarySearch{PositionId: helper.StringToInt64(positionId)}
}
