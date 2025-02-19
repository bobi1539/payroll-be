package response

import "payroll/model/domain"

type BasicSalaryResponse struct {
	Id           int64            `json:"id"`
	SalaryAmount int64            `json:"salaryAmount"`
	TotalYear    int32            `json:"totalYear"`
	Position     PositionResponse `json:"position"`
	BaseDomainResponse
}

func ToBasicSalaryResponse(basicSalary *domain.BasicSalary) BasicSalaryResponse {
	return BasicSalaryResponse{
		Id:                 basicSalary.Id,
		SalaryAmount:       basicSalary.SalaryAmount,
		TotalYear:          basicSalary.TotalYear,
		Position:           ToPositionResponse(basicSalary.Position),
		BaseDomainResponse: ToBaseDomainResponse(&basicSalary.BaseDomain),
	}
}

func ToBasicSalaryResponses(basicSalaries []domain.BasicSalary) []BasicSalaryResponse {
	if len(basicSalaries) == 0 {
		return make([]BasicSalaryResponse, 0)
	}

	var responses []BasicSalaryResponse
	for _, basicSalary := range basicSalaries {
		responses = append(responses, ToBasicSalaryResponse(&basicSalary))
	}
	return responses
}
