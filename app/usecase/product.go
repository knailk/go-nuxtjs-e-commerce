package usecase

import (
	"strings"

	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)



//Service product usecase.
type ProductService struct {
	repo repository.ProductQuery
}
//NewService create new service.
func NewProductService(r repository.ProductQuery) ProductUsecase {
	return &ProductService{
		repo: r,
	}
}
func (s *ProductService) GetProduct(id entity.ID) (*entity.Product, error){
	return s.repo.Get(id)
}
func (s *ProductService) SearchProducts(query string) ([]*entity.Product, error){
	return s.repo.Search(strings.ToLower(query))
}
func (s *ProductService) ListProducts(id int64) ([]*entity.Product, error){
	return s.repo.List(id)
}
func (s *ProductService) CreateProduct(p *entity.Product) (entity.ID, error){
	return s.repo.Create(p)
}
func (s *ProductService) UpdateProduct(e *entity.Product) error{
	return s.repo.Update(e)
}
func (s *ProductService) DeleteProduct(id entity.ID) error{
	u, err := s.GetProduct(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}