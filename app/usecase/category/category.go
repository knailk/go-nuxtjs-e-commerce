package category

import (
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)

//UseCase interface.
type Usecase interface {
	GetCategory(id int64) (*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
}

// Service product usecase.
type Service struct {
	repo repository.CategoryQuery
}

// NewService create new service.
func NewService(r repository.CategoryQuery) *Service {
	return &Service{
		repo: r,
	}
}
func (s *Service) GetCategory(id int64) (*entity.Category, error) {
	return s.repo.Get(id)
}
func (s *Service) ListCategories() ([]*entity.Category, error) {
	return s.repo.List()
}
