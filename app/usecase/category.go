package usecase

import (
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)



// Service product usecase.
type CategoryService struct {
	repo repository.CategoryQuery
}

// NewService create new service.
func NewCategoryService(r repository.CategoryQuery) CategoryUsecase {
	return &CategoryService{
		repo: r,
	}
}
func (s *CategoryService) GetCategory(id int64) (*entity.Category, error) {
	return s.repo.Get(id)
}
func (s *CategoryService) ListCategories() ([]*entity.Category, error) {
	return s.repo.List()
}
