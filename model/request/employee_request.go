package request

type EmployeeRequest struct {
	Name              string `validate:"required" json:"name"`
	PhoneNumber       string `validate:"required" json:"phoneNumber"`
	Email             string `validate:"required" json:"email"`
	Address           string `validate:"required" json:"address"`
	WorkStatus        string `validate:"required" json:"workStatus"`
	BankAccountNumber string `validate:"required" json:"bankAccountNumber"`
	BankAccountName   string `validate:"required" json:"bankAccountName"`
	Npwp              string `validate:"required" json:"npwp"`
	DateOfBirth       string `validate:"required" json:"dateOfBirth"`
	JoinDate          string `validate:"required" json:"joinDate"`
	IsMarried         *bool  `validate:"required" json:"isMarried"`
	TotalChild        *int32 `validate:"required" json:"totalChild"`
	PositionId        int64  `validate:"required" json:"positionId"`
}
