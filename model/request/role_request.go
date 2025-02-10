package request

type RoleRequest struct {
	Name string `validate:"required" json:"name"`
}
