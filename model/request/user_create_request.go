package request

type UserCreateRequest struct {
	Name     string `validate:"required" json:"name"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	RoleId   int64  `validate:"required" json:"roleId"`
}
