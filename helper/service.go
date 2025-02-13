package helper

import (
	"payroll/constant"
	"payroll/model/domain"
	"time"
)

func SetCreated(base *domain.BaseDomain) {
	base.CreatedAt = time.Now()
	base.CreatedBy = 1
	base.CreatedByName = constant.SYSTEM
}

func SetUpdated(base *domain.BaseDomain) {
	base.UpdatedAt = time.Now()
	base.UpdatedBy = 1
	base.UpdatedByName = constant.SYSTEM
}
