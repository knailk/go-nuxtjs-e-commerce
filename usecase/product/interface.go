package product

import (
	"github.com/knailk/go-shopee/entity"
)

//Reader interface
type Reader interface {
	//get product by id
	Get(id entity.ID) (*entity.Product, error)
	//search list product by query
	Search(query string) ([]*entity.Product, error)
	//get list product by category id
	List(entity.ID) ([]*entity.Product, error)
}

//Writer interface
type Writer interface {
	Create(e *entity.Product) (entity.ID, error)
	Update(e *entity.Product) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetProduct(id entity.ID) (*entity.Product, error)
	SearchProducts(query string) ([]*entity.Product, error)
	ListProducts(entity.ID) ([]*entity.Product, error)
	CreateProduct(e *entity.Product) (entity.ID, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id entity.ID) error
}
