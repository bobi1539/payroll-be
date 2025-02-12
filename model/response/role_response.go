package response

import "payroll/model/domain"

type RoleResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	BaseDomainResponse
}

func ToRoleResponse(role *domain.Role) RoleResponse {
	return RoleResponse{
		Id:                 role.ID,
		Name:               role.Name,
		BaseDomainResponse: ToBaseDomainResponse(&role.BaseDomain),
	}
}

func ToRoleResponses(roles []domain.Role) []RoleResponse {
	if roles == nil {
		return make([]RoleResponse, 0)
	}
	var responses []RoleResponse
	for _, role := range roles {
		responses = append(responses, ToRoleResponse(&role))
	}
	return responses
}
