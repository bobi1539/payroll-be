package repositoryimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/model/search"
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
	result := allowanceRepository.DB.
		Preload(domain.POSITION).
		Preload(domain.ALLOWANCE_TYPE).
		First(allowance, "id = ?", id)

	if result.Error != nil {
		return allowance, errors.New(constant.ALLOWANCE_NOT_FOUND)
	}
	return allowance, nil
}

func (allowanceRepository *AllowanceRepositoryImpl) FindAll(search *search.AllowanceSearch) []domain.Allowance {
	var allowances []domain.Allowance
	allowanceRepository.DB.
		Preload(domain.POSITION).
		Preload(domain.ALLOWANCE_TYPE).
		Where(allowanceRepository.searchEqual(), search.PositionId).
		Order("id ASC").
		Find(&allowances)
	return allowances
}

func (allowanceRepository *AllowanceRepositoryImpl) FindAllPagination(search *search.AllowanceSearch, pagination *dto.Pagination) []domain.Allowance {
	var allowances []domain.Allowance
	allowanceRepository.DB.
		Preload(domain.POSITION).
		Preload(domain.ALLOWANCE_TYPE).
		Where(allowanceRepository.searchEqual(), search.PositionId).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&allowances)
	return allowances
}

func (allowanceRepository *AllowanceRepositoryImpl) FindTotalItem(search *search.AllowanceSearch) int64 {
	var totalItem int64
	allowanceRepository.DB.
		Model(&domain.Allowance{}).
		Where(allowanceRepository.searchEqual(), search.PositionId).
		Count(&totalItem)
	return totalItem
}

func (allowanceRepository *AllowanceRepositoryImpl) searchEqual() string {
	return "position_id = ?"
}
