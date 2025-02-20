package response

import "payroll/model/domain"

type AllowanceResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	BaseDomainResponse
}

func ToAllowanceResponse(allowance *domain.Allowance) AllowanceResponse {
	return AllowanceResponse{
		Id:                 allowance.Id,
		Name:               allowance.Name,
		BaseDomainResponse: ToBaseDomainResponse(&allowance.BaseDomain),
	}
}

func ToAllowanceResponses(allowances []domain.Allowance) []AllowanceResponse {
	if len(allowances) == 0 {
		return make([]AllowanceResponse, 0)
	}

	var responses []AllowanceResponse
	for _, allowance := range allowances {
		responses = append(responses, ToAllowanceResponse(&allowance))
	}
	return responses
}
