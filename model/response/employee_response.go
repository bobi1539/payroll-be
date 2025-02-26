package response

import (
	"payroll/constant"
	"payroll/model/domain"
)

type EmployeeResponse struct {
	Id                int64            `json:"id"`
	Name              string           `json:"name"`
	PhoneNumber       string           `json:"phoneNumber"`
	Email             string           `json:"email"`
	Address           string           `json:"address"`
	WorkStatus        string           `json:"workStatus"`
	BankAccountNumber string           `json:"bankAccountNumber"`
	BankAccountName   string           `json:"bankAccountName"`
	Npwp              string           `json:"npwp"`
	DateOfBirth       string           `json:"dateOfBirth"`
	JoinDate          string           `json:"joinDate"`
	IsMarried         bool             `json:"isMarried"`
	TotalChild        int32            `json:"totalChild"`
	Position          PositionResponse `json:"position"`
	BaseDomainResponse
}

func ToEmployeeResponse(employee *domain.Employee) EmployeeResponse {
	return EmployeeResponse{
		Id:                 employee.Id,
		Name:               employee.Name,
		PhoneNumber:        employee.PhoneNumber,
		Email:              employee.Email,
		Address:            employee.Address,
		WorkStatus:         employee.WorkStatus,
		BankAccountNumber:  employee.BankAccountNumber,
		BankAccountName:    employee.BankAccountName,
		Npwp:               employee.Npwp,
		DateOfBirth:        employee.DateOfBirth.Format(constant.DATE_LAYOUT),
		JoinDate:           employee.JoinDate.Format(constant.DATE_LAYOUT),
		IsMarried:          employee.IsMarried,
		TotalChild:         employee.TotalChild,
		Position:           ToPositionResponse(employee.Position),
		BaseDomainResponse: ToBaseDomainResponse(&employee.BaseDomain),
	}
}

func ToEmployeeResponses(employees []domain.Employee) []EmployeeResponse {
	if len(employees) == 0 {
		return make([]EmployeeResponse, 0)
	}

	var responses []EmployeeResponse
	for _, employee := range employees {
		responses = append(responses, ToEmployeeResponse(&employee))
	}
	return responses
}
