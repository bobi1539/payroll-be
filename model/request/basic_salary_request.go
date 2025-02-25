package request

type BasicSalaryRequest struct {
	SalaryAmount int64  `validate:"required" json:"salaryAmount"`
	TotalYear    *int32 `validate:"required" json:"totalYear"`
	PositionId   int64  `validate:"required" json:"positionId"`
}
