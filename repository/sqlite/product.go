package sqlite

import (
	"database/sql"
	"time"

	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase/product"
)

type ProductRepo struct {
	db *sql.DB
}

// NewProductRepo create an implementation of product repository.
func NewProductRepo(db *sql.DB) product.Repository {
	return &ProductRepo{
		db: db,
	}
}

// Get product by id.
func (r *ProductRepo) Get(id entity.ID) (*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select id,name,price,description,quantitySold,availableUnits,createdAt,updatedAt,categoryId from product where id=? and isDeleted = 0`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p entity.Product
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Search list product by query.
func (r *ProductRepo) Search(query string) ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select id,name,price,description,quantitySold,availableUnits,createdAt,updatedAt,categoryId from product where name like ? and isDeleted = 0`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	var result []*entity.Product
	for rows.Next() {
		var p entity.Product
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		result = append(result, &p)
	}
	return result, nil
}

// List product by category id.
func (r *ProductRepo) List(id int64) ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select id,name,price,description,quantitySold,availableUnits,createdAt,updatedAt,categoryId from product where categoryID = ? and isDeleted = 0`)
	if err != nil {
		return nil, err
	}
	var result []*entity.Product
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p entity.Product
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
		if err != nil{
			return nil, err
		}
		result = append(result, &p)
	}
	return result, nil
}

// Create a product.
func (r *ProductRepo) Create(e *entity.Product) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
	insert into product(id, name, price, description, quantitySold, availableUnits, createdAt, updatedAt, categoryId, isDeleted) 
	values (?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return e.ProductID, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		e.ProductID,
		e.Name,
		e.Price,
		e.Description,
		e.QuantitySold,
		e.AvailableUnits,
		e.CreatedAt,
		e.UpdatedAt,
		e.CategoryID,
		false,
	)
	if err != nil {
		return e.ProductID, err
	}
	return e.ProductID, nil
}

// Update a product.
func (r *ProductRepo) Update(e *entity.Product) error {
	stmt, err := r.db.Prepare(`update product set 
	name = ?, 
	price = ?, 
	description = ?, 
	quantitySold = ?,
	availableUnits =?, 
	updatedAt = ? 
	where id = ? and isDeleted = 0`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.Name, 
		e.Price, 
		e.Description, 
		e.QuantitySold, 
		e.AvailableUnits, 
		time.Now().Format(time.RFC3339), 
		e.ProductID,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

// Delete a product.
func (r *ProductRepo) Delete(id entity.ID) error {
	_, err := r.db.Exec("update product set isDeleted = true where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
