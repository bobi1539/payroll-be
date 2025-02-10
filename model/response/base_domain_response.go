package response

import (
	"payroll/model/domain"
	"time"
)

type BaseDomainResponse struct {
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	CreatedBy     int64     `json:"createdBy"`
	UpdatedBy     int64     `json:"updatedBy"`
	CreatedByName string    `json:"createdByName"`
	UpdatedByName string    `json:"updatedByName"`
	IsDeleted     bool      `json:"isDeleted"`
}

func ToBaseDomainResponse(base *domain.BaseDomain) BaseDomainResponse {
	return BaseDomainResponse{
		CreatedAt:     base.CreatedAt,
		UpdatedAt:     base.UpdatedAt,
		CreatedBy:     base.CreatedBy,
		UpdatedBy:     base.UpdatedBy,
		CreatedByName: base.CreatedByName,
		UpdatedByName: base.UpdatedByName,
		IsDeleted:     base.IsDeleted,
	}
}
