package repositoryimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AllowanceRepositoryImpl struct {
	DB *gorm.DB
}

func NewAllowanceRepositoryImpl(db *gorm.DB) repository.AllowanceRepository {
	return &AllowanceRepositoryImpl{DB: db}
}

func (allowanceRepository *AllowanceRepositoryImpl) Create(allowance *domain.Allowance) *domain.Allowance {
	allowanceRepository.DB.Create(allowance)
	return allowance
}

func (allowanceRepository *AllowanceRepositoryImpl) Update(allowance *domain.Allowance) *domain.Allowance {
	allowanceRepository.DB.Save(allowance)
	return allowance
}

func (allowanceRepository *AllowanceRepositoryImpl) Delete(id int64) {
	result := allowanceRepository.DB.Where("id = ?", id).Delete(&domain.Allowance{})
	if result.Error != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.CANNOT_DELETE_THIS_DATA))
	}
}

func (allowanceRepository *AllowanceRepositoryImpl) FindById(id int64) (*domain.Allowance, error) {
	allowance := &domain.Allowance{}
	result := allowanceRepository.DB.First(allowance, "id = ?", id)

	if result.Error != nil {
		return allowance, errors.New(constant.ALLOWANCE_NOT_FOUND)
	}
	return allowance, nil
}

func (allowanceRepository *AllowanceRepositoryImpl) FindAll(search *dto.Search) []domain.Allowance {
	var allowances []domain.Allowance
	allowanceRepository.DB.
		Where(allowanceRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Find(&allowances)
	return allowances
}

func (allowanceRepository *AllowanceRepositoryImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.Allowance {
	var allowances []domain.Allowance
	allowanceRepository.DB.
		Where(allowanceRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&allowances)
	return allowances
}

func (allowanceRepository *AllowanceRepositoryImpl) FindTotalItem() int64 {
	var totalItem int64
	allowanceRepository.DB.
		Model(&domain.Allowance{}).
		Count(&totalItem)
	return totalItem
}

func (allowanceRepository *AllowanceRepositoryImpl) searchLike() string {
	return "LOWER(name) LIKE ?"
}
