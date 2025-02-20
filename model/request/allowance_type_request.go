package request

type AllowanceTypeRequest struct {
	Name string `validate:"required" json:"name"`
}
