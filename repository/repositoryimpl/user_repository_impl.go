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

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) repository.UserRepository {
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

func (userRepository *UserRepositoryImpl) Delete(id int64) {
	result := userRepository.DB.Where("id = ?", id).Delete(&domain.User{})
	if result.Error != nil {
		exception.PanicErrorBusiness(fiber.StatusBadRequest, errors.New(constant.CANNOT_DELETE_THIS_DATA))
	}
}

func (userRepository *UserRepositoryImpl) FindById(id int64) (*domain.User, error) {
	user := &domain.User{}
	result := userRepository.DB.
		Preload(domain.ROLE).
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
		Preload(domain.ROLE).
		Where(userRepository.searchLike(), valueLike, valueLike).
		Order("id ASC").
		Find(&users)
	return users
}

func (userRepository *UserRepositoryImpl) FindAllPagination(search *dto.Search, pagination *dto.Pagination) []domain.User {
	var users []domain.User
	valueLike := helper.StringQueryLike(search.Value)
	userRepository.DB.
		Preload(domain.ROLE).
		Where(userRepository.searchLike(), valueLike, valueLike).
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
		Count(&totalItem)
	return totalItem
}

func (userRepository *UserRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
	user := &domain.User{}
	result := userRepository.DB.
		Preload(domain.ROLE).
		First(&user, "username = ?", username)

	if result.Error == nil {
		return user, nil
	}
	return nil, errors.New(constant.DATA_NOT_FOUND)
}

func (userRepository *UserRepositoryImpl) searchLike() string {
	return "LOWER(name) LIKE ? OR LOWER(username) LIKE ?"
}
