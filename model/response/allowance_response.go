package response

import "payroll/model/domain"

type AllowanceResponse struct {
	Id              int64                 `json:"id"`
	Position        PositionResponse      `json:"position"`
	AllowanceType   AllowanceTypeResponse `json:"allowanceType"`
	AllowanceAmount int64                 `json:"allowanceAmount"`
	BaseDomainResponse
}

func ToAllowanceResponse(allowance *domain.Allowance) AllowanceResponse {
	return AllowanceResponse{
		Id:                 allowance.Id,
		Position:           ToPositionResponse(allowance.Position),
		AllowanceType:      ToAllowanceTypeResponse(allowance.AllowanceType),
		AllowanceAmount:    allowance.AllowanceAmount,
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
