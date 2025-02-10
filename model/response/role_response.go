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
