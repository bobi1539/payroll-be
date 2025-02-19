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

type PositionRepositoryImpl struct {
	DB *gorm.DB
}

func NewPositionRepositoryImpl(db *gorm.DB) repository.PositionRepository {
	return &PositionRepositoryImpl{DB: db}
}

func (positionRepository *PositionRepositoryImpl) Create(position *domain.Position) *domain.Position {
	positionRepository.DB.Create(position)
	return position
}

func (positionRepository *PositionRepositoryImpl) Update(position *domain.Position) *domain.Position {
	positionRepository.DB.Save(position)
	return position
}

func (positionRepository *PositionRepositoryImpl) Delete(id int64) {
	result := positionRepository.DB.Where("id = ?", id).Delete(&domain.Position{})
	if result.Error != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.CANNOT_DELETE_THIS_DATA))
	}
}

func (positionRepository *PositionRepositoryImpl) FindById(id int64) (*domain.Position, error) {
	position := &domain.Position{}
	result := positionRepository.DB.First(position, "id = ?", id)

	if result.Error != nil {
		return position, errors.New(constant.POSITION_NOT_FOUND)
	}
	return position, nil
}

func (positionRepository *PositionRepositoryImpl) FindAll(search *dto.Search) []domain.Position {
	var positions []domain.Position
	positionRepository.DB.
		Where(positionRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Find(&positions)
	return positions
}

func (positionRepository *PositionRepositoryImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.Position {
	var positions []domain.Position
	positionRepository.DB.
		Where(positionRepository.searchLike(), helper.StringQueryLike(search.Value)).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&positions)
	return positions
}

func (positionRepository *PositionRepositoryImpl) FindTotalItem() int64 {
	var totalItem int64
	positionRepository.DB.
		Model(&domain.Position{}).
		Count(&totalItem)
	return totalItem
}

func (positionRepository *PositionRepositoryImpl) searchLike() string {
	return "LOWER(name) LIKE ?"
}
