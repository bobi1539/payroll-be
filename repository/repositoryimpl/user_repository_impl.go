package repositoryimpl

import (
	"errors"
	"payroll/constant"
	"payroll/helper"
	"payroll/model/domain"
	"payroll/model/dto"
	"payroll/repository"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (userRepository *UserRepositoryImpl) Create(user *domain.User) *domain.User {
	userRepository.DB.Create(user)
	return user
}

func (userRepository *UserRepositoryImpl) Update(user *domain.User) *domain.User {
	userRepository.DB.Save(user)
	return user
}

func (userRepository *UserRepositoryImpl) FindById(id int64) (*domain.User, error) {
	user := &domain.User{}
	result := userRepository.DB.
		Where(constant.IS_DELETED_FALSE).
		First(&user, "id = ?", id)

	if result.Error == nil {
		return user, nil
	}
	return user, errors.New(constant.DATA_NOT_FOUND)
}

func (userRepository *UserRepositoryImpl) FindAll(search *dto.Search) []domain.User {
	var users []domain.User
	valueLike := helper.StringQueryLike(search.Value)
	userRepository.DB.
		Where(userRepository.searchLike(), valueLike, valueLike).
		Where(constant.IS_DELETED_FALSE).
		Order("id ASC").
		Find(&users)
	return users
}

func (userRepository *UserRepositoryImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.User {
	var users []domain.User
	valueLike := helper.StringQueryLike(search.Value)
	userRepository.DB.
		Where(userRepository.searchLike(), valueLike, valueLike).
		Where(constant.IS_DELETED_FALSE).
		Order("id ASC").
		Offset(pagination.PageNumber).
		Limit(pagination.PageSize).
		Find(&users)
	return users
}

func (userRepository *UserRepositoryImpl) FindTotalItem() int64 {
	var totalItem int64
	userRepository.DB.
		Model(&domain.User{}).
		Where(constant.IS_DELETED_FALSE).
		Count(&totalItem)
	return totalItem
}

func (userRepository *UserRepositoryImpl) searchLike() string {
	return "LOWER(name) LIKE ? OR LOWER(username) LIKE ?"
}
