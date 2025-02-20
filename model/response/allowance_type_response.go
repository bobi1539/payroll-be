package response

import "payroll/model/domain"

type AllowanceTypeResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	BaseDomainResponse
}

func ToAllowanceTypeResponse(allowanceType *domain.AllowanceType) AllowanceTypeResponse {
	return AllowanceTypeResponse{
		Id:                 allowanceType.Id,
		Name:               allowanceType.Name,
		BaseDomainResponse: ToBaseDomainResponse(&allowanceType.BaseDomain),
	}
}

func ToAllowanceTypeResponses(allowanceTypes []domain.AllowanceType) []AllowanceTypeResponse {
	if len(allowanceTypes) == 0 {
		return make([]AllowanceTypeResponse, 0)
	}

	var responses []AllowanceTypeResponse
	for _, allowanceType := range allowanceTypes {
		responses = append(responses, ToAllowanceTypeResponse(&allowanceType))
	}
	return responses
}
