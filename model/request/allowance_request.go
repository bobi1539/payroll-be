package request

type AllowanceRequest struct {
	Name string `validate:"required" json:"name"`
}
