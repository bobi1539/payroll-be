package repositoryimpl

import (
	"errors"
	"payroll/constant"
	"payroll/exception"
	"payroll/model/domain"
	"payroll/model/dto"
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

func (basicSalaryRepository *BasicSalaryRepositoryImpl) Create(basicSalary *domain.BasicSalary) *domain.BasicSalary {
	basicSalaryRepository.DB.Create(basicSalary)
	return basicSalary
}

func (basicSalaryRepository *BasicSalaryRepositoryImpl) Update(basicSalary *domain.BasicSalary) *domain.BasicSalary {
	basicSalaryRepository.DB.Save(basicSalary)
	return basicSalary
}

func (basicSalaryRepository *BasicSalaryRepositoryImpl) Delete(id int64) {
	result := basicSalaryRepository.DB.Where("id = ?", id).Delete(&domain.BasicSalary{})
	if result.Error != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.CANNOT_DELETE_THIS_DATA))
	}
}

func (basicSalaryRepository *BasicSalaryRepositoryImpl) FindById(id int64) (*domain.BasicSalary, error) {
	basicSalary := &domain.BasicSalary{}
	result := basicSalaryRepository.DB.
		Preload(domain.POSITION).
		First(basicSalary, "id = ?", id)

	if result.Error != nil {
		return basicSalary, errors.New(constant.DATA_NOT_FOUND)
	}
	return basicSalary, nil
}

func (basicSalaryRepository *BasicSalaryRepositoryImpl) FindAll() []domain.BasicSalary {
	var basicSalaries []domain.BasicSalary
	basicSalaryRepository.DB.
		Preload(domain.POSITION).
		Order("id ASC").
		Find(&basicSalaries)
	return basicSalaries
}

func (basicSalaryRepository *BasicSalaryRepositoryImpl) FindAllPagination(pagination *dto.Pagination) []domain.BasicSalary {
	var basicSalaries []domain.BasicSalary
	basicSalaryRepository.DB.
		Preload(domain.POSITION).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&basicSalaries)
	return basicSalaries
}

func (basicSalaryRepository *BasicSalaryRepositoryImpl) FindTotalItem() int64 {
	var totalItem int64
	basicSalaryRepository.DB.
		Model(&domain.BasicSalary{}).
		Count(&totalItem)
	return totalItem
}
