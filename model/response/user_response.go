package response

import "payroll/model/domain"

type UserResponse struct {
	Id       int64        `json:"id"`
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Role     RoleResponse `json:"role"`
	BaseDomainResponse
}

func ToUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		Id:                 user.Id,
		Name:               user.Name,
		Username:           user.Username,
		Role:               ToRoleResponse(user.Role),
		BaseDomainResponse: ToBaseDomainResponse(&user.BaseDomain),
	}
}

func ToUserResponses(users []domain.User) []UserResponse {
	if len(users) == 0 {
		return make([]UserResponse, 0)
	}

	var responses []UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(&user))
	}
	return responses
}
