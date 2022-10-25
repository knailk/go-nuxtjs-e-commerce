package product

import "github.com/knailk/go-shopee/app/entity"

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
//TODO implement method of usecase interface
func GetProduct(id entity.ID) (*entity.Product, error)
func SearchProducts(query string) ([]*entity.Product, error)
func ListProducts(entity.ID) ([]*entity.Product, error)
func CreateProduct(e *entity.Product) (entity.ID, error)
func UpdateProduct(e *entity.Product) error
func DeleteProduct(id entity.ID) error