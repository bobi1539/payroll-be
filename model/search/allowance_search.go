package search

import "payroll/helper"

type AllowanceSearch struct {
	PositionId int64
}

func BuildAllowanceSearch(positionId string) AllowanceSearch {
	return AllowanceSearch{PositionId: helper.StringToInt64(positionId)}
}
