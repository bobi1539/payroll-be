package request

type AllowanceRequest struct {
	PositionId      int64 `validate:"required" json:"positionId"`
	AllowanceTypeId int64 `validate:"required" json:"allowanceTypeId"`
	AllowanceAmount int64 `validate:"required" json:"allowanceAmount"`
}
