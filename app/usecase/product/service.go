package product

import (
	"strings"

	"github.com/knailk/go-shopee/app/entity"
)

//Service product usecase.
type Service struct {
	repo Repository
}
//NewService create new service.
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}
func (s *Service) GetProduct(id entity.ID) (*entity.Product, error){
	return s.repo.Get(id)
}
func (s *Service) SearchProducts(query string) ([]*entity.Product, error){
	return s.repo.Search(strings.ToLower(query))
}
func (s *Service) ListProducts(id int64) ([]*entity.Product, error){
	return s.repo.List(id)
}
func (s *Service) CreateProduct(name string, price int64, description string, quantitySold int64, availableInits int64, category int64) (entity.ID, error){
	p := entity.NewProduct(name,price,description,quantitySold,availableInits,category)
	return s.repo.Create(p)
}
func (s *Service) UpdateProduct(e *entity.Product) error{
	return s.repo.Update(e)
}
func (s *Service) DeleteProduct(id entity.ID) error{
	u, err := s.GetProduct(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}