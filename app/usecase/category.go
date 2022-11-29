package usecase

import (
	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
	"github.com/knailk/go-nuxtjs-e-commerce/repository"
)

// Service product usecase.
type CategoryService struct {
	dao repository.DAO
}

// NewService create new service.
func NewCategoryService(dao repository.DAO) CategoryUsecase {
	return &CategoryService{
		dao: dao,
	}
}
func (s *CategoryService) GetCategory(id int64) (*entity.Category, error) {
	return s.dao.NewCategoryRepo().Get(id)
}
func (s *CategoryService) ListCategories() ([]*entity.Category, error) {
	return s.dao.NewCategoryRepo().List()
}
