package category

import (
	"github.com/knailk/go-shopee/app/entity"
)

// Service product usecase.
type Service struct {
	repo Repository
}

// NewService create new service.
func NewService(r Repository) *Service {
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
