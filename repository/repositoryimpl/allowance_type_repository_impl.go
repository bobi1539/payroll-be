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

type AllowanceTypeRepositoryImpl struct {
	DB *gorm.DB
}

func NewAllowanceTypeRepositoryImpl(db *gorm.DB) repository.AllowanceTypeRepository {
	return &AllowanceTypeRepositoryImpl{DB: db}
}

func (atRepository *AllowanceTypeRepositoryImpl) Create(allowanceType *domain.AllowanceType) *domain.AllowanceType {
	atRepository.DB.Create(allowanceType)
	return allowanceType
}

func (atRepository *AllowanceTypeRepositoryImpl) Update(allowanceType *domain.AllowanceType) *domain.AllowanceType {
	atRepository.DB.Save(allowanceType)
	return allowanceType
}

func (atRepository *AllowanceTypeRepositoryImpl) Delete(id int64) {
	result := atRepository.DB.Where("id = ?", id).Delete(&domain.AllowanceType{})
	if result.Error != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.CANNOT_DELETE_THIS_DATA))
	}
}

func (atRepository *AllowanceTypeRepositoryImpl) FindById(id int64) (*domain.AllowanceType, error) {
	allowanceType := &domain.AllowanceType{}
	result := atRepository.DB.First(allowanceType, "id = ?", id)

	if result.Error != nil {
		return allowanceType, errors.New(constant.ALLOWANCE_TYPE_NOT_FOUND)
	}
	return allowanceType, nil
}

func (atRepository *AllowanceTypeRepositoryImpl) FindAll(search *dto.Search) []domain.AllowanceType {
	var allowanceTypes []domain.AllowanceType
	atRepository.DB.
		Where(atRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Find(&allowanceTypes)
	return allowanceTypes
}

func (atRepository *AllowanceTypeRepositoryImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.AllowanceType {
	var allowanceTypes []domain.AllowanceType
	atRepository.DB.
		Where(atRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&allowanceTypes)
	return allowanceTypes
}

func (atRepository *AllowanceTypeRepositoryImpl) FindTotalItem() int64 {
	var totalItem int64
	atRepository.DB.
		Model(&domain.AllowanceType{}).
		Count(&totalItem)
	return totalItem
}

func (atRepository *AllowanceTypeRepositoryImpl) searchLike() string {
	return "LOWER(name) LIKE ?"
}
