package request

type UserUpdateRequest struct {
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" json:"username"`
	RoleId   int64  `validate:"required" json:"roleId"`
}
