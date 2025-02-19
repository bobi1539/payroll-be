package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
)

type PositionRepository interface {
	Create(position *domain.Position) *domain.Position
	Update(position *domain.Position) *domain.Position
	Delete(id int64)
	FindById(id int64) (*domain.Position, error)
	FindAll(search *dto.Search) []domain.Position
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.Position
	FindTotalItem() int64
}
