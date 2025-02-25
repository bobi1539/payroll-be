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

type BasicSalaryRepositoryImpl struct {
	DB *gorm.DB
}

func NewBasicSalaryRepositoryImpl(db *gorm.DB) repository.BasicSalaryRepository {
	return &BasicSalaryRepositoryImpl{DB: db}
}

func (bsRepository *BasicSalaryRepositoryImpl) Create(basicSalary *domain.BasicSalary) *domain.BasicSalary {
	bsRepository.DB.Create(basicSalary)
	return basicSalary
}

func (bsRepository *BasicSalaryRepositoryImpl) Update(basicSalary *domain.BasicSalary) *domain.BasicSalary {
	bsRepository.DB.Save(basicSalary)
	return basicSalary
}

func (bsRepository *BasicSalaryRepositoryImpl) Delete(id int64) {
	result := bsRepository.DB.Where("id = ?", id).Delete(&domain.BasicSalary{})
	if result.Error != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.CANNOT_DELETE_THIS_DATA))
	}
}

func (bsRepository *BasicSalaryRepositoryImpl) FindById(id int64) (*domain.BasicSalary, error) {
	basicSalary := &domain.BasicSalary{}
	result := bsRepository.DB.
		Preload(domain.POSITION).
		First(basicSalary, "id = ?", id)

	if result.Error != nil {
		return basicSalary, errors.New(constant.BASIC_SALARY_NOT_FOUND)
	}
	return basicSalary, nil
}

func (bsRepository *BasicSalaryRepositoryImpl) FindByPositionIdAndTotalYear(positionId int64, totalyear int32) (*domain.BasicSalary, error) {
	basicSalary := &domain.BasicSalary{}
	result := bsRepository.DB.
		Preload(domain.POSITION).
		First(basicSalary, "position_id = ? AND total_year = ?", positionId, totalyear)

	if result.Error != nil {
		return nil, errors.New(constant.BASIC_SALARY_NOT_FOUND)
	}
	return basicSalary, nil
}

func (bsRepository *BasicSalaryRepositoryImpl) FindAll(search *search.BasicSalarySearch) []domain.BasicSalary {
	var basicSalaries []domain.BasicSalary
	bsRepository.DB.
		Preload(domain.POSITION).
		Where(bsRepository.searchEqual(), search.PositionId).
		Order("id ASC").
		Find(&basicSalaries)
	return basicSalaries
}

func (bsRepository *BasicSalaryRepositoryImpl) FindAllPagination(search *search.BasicSalarySearch, pagination *dto.Pagination) []domain.BasicSalary {
	var basicSalaries []domain.BasicSalary
	bsRepository.DB.
		Preload(domain.POSITION).
		Where(bsRepository.searchEqual(), search.PositionId).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&basicSalaries)
	return basicSalaries
}

func (bsRepository *BasicSalaryRepositoryImpl) FindTotalItem(search *search.BasicSalarySearch) int64 {
	var totalItem int64
	bsRepository.DB.
		Model(&domain.BasicSalary{}).
		Where(bsRepository.searchEqual(), search.PositionId).
		Count(&totalItem)
	return totalItem
}

func (bsRepository *BasicSalaryRepositoryImpl) searchEqual() string {
	return "position_id = ?"
}
