package usecase

import (
	"strings"

	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
	"github.com/knailk/go-nuxtjs-e-commerce/repository"
)

// Service product usecase.
type ProductService struct {
	dao repository.DAO
}

// NewService create new service.
func NewProductService(dao repository.DAO) ProductUsecase {
	return &ProductService{
		dao: dao,
	}
}
func (s *ProductService) GetProduct(id entity.ID) (*entity.Product, error) {
	return s.dao.NewProductRepo().Get(id)
}
func (s *ProductService) SearchProductsByQuery(query string) ([]*entity.Product, error){
	return s.dao.NewProductRepo().SearchByQuery(query)
}
func (s *ProductService) TopProduct() ([]*entity.Product, error) {
	return s.dao.NewProductRepo().Top()
}
func (s *ProductService) SearchProducts(query string) ([]*entity.Product, error) {
	return s.dao.NewProductRepo().Search(strings.ToLower(query))
}
func (s *ProductService) ListProducts(id int64) ([]*entity.Product, error) {
	return s.dao.NewProductRepo().List(id)
}
func (s *ProductService) CreateProduct(p *entity.Product) (entity.ID, error) {
	return s.dao.NewProductRepo().Create(p)
}
func (s *ProductService) UpdateProduct(e *entity.Product) error {
	return s.dao.NewProductRepo().Update(e)
}
func (s *ProductService) DeleteProduct(id entity.ID) error {
	u, err := s.GetProduct(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.dao.NewProductRepo().Delete(id)
}
