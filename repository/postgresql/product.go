package postgresql

import (
	"database/sql"

	"github.com/knailk/go-shopee/entity"
	"github.com/knailk/go-shopee/usecase/product"
)

type ProductRepo struct {
	db *sql.DB
}

//create an implementation of product repository
func NewProductRepo(db *sql.DB) product.Repository{
	return &ProductRepo{
		db: db,
	}
}

//TODO implement method of product repository interface
//Querying, or Creating/ Inserting into any database will be stored here.

//get product by id
func (r *ProductRepo) Get(id entity.ID) (*entity.Product, error){
	return nil,nil
}
//search list product by query
func (r *ProductRepo) Search(query string) ([]*entity.Product, error){
	return nil,nil
}
//get list product by category id
func (r *ProductRepo) List(entity.ID) ([]*entity.Product, error){
	return nil,nil
}
//create a product
func (r *ProductRepo) Create(e *entity.Product) (entity.ID, error){
	return entity.NewID(),nil
}
//update a product
func (r *ProductRepo) Update(e *entity.Product) error {
	return nil
}
//delete a product
func (r *ProductRepo) Delete(id entity.ID) error {
	return nil
}

