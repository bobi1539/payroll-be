package repository

import (
	"payroll/model/domain"
	"payroll/model/dto"
)

type UserRepository interface {
	Create(user *domain.User) *domain.User
	Update(user *domain.User) *domain.User
	Delete(id int64)
	FindById(id int64) (*domain.User, error)
	FindAll(search *dto.Search) []domain.User
	FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.User
	FindTotalItem(search *dto.Search) int64
	FindByUsername(username string) (*domain.User, error)
}
