package request

type PositionRequest struct {
	Name string `validate:"required" json:"name"`
}
