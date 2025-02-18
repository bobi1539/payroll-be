package helper

import (
	"payroll/model/domain"
	"time"
)

func SetCreated(base *domain.BaseDomain, user *domain.User) {
	base.CreatedAt = time.Now()
	base.CreatedBy = user.ID
	base.CreatedByName = user.Name
}

func SetUpdated(base *domain.BaseDomain, user *domain.User) {
	base.UpdatedAt = time.Now()
	base.UpdatedBy = user.ID
	base.UpdatedByName = user.Name
}
