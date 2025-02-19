package response

import "payroll/model/domain"

type PositionResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	BaseDomainResponse
}

func ToPositionResponse(position *domain.Position) PositionResponse {
	return PositionResponse{
		Id:                 position.Id,
		Name:               position.Name,
		BaseDomainResponse: ToBaseDomainResponse(&position.BaseDomain),
	}
}

func ToPositionResponses(positions []domain.Position) []PositionResponse {
	if len(positions) == 0 {
		return make([]PositionResponse, 0)
	}

	var responses []PositionResponse
	for _, position := range positions {
		responses = append(responses, ToPositionResponse(&position))
	}
	return responses
}
