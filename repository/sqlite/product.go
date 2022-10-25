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

//Querying, or Creating/ Inserting into any database will be stored here.

// Get product by id.
func (r *ProductRepo) Get(id entity.ID) (*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select * from product where id=? and id_deleted = false`)
	if err != nil {
		return nil, err
	}
	var p entity.Product
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID, &p.IsDeleted)
	}
	return &p, nil
}

// Search list product by query.
func (r *ProductRepo) Search(query string) ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select * from product where name like ? and is_delete = false`)
	if err != nil {
		return nil, err
	}
	var result []*entity.Product
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p entity.Product
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID, &p.IsDeleted)
		result = append(result, &p)
	}
	return result, nil
}

// List product by category id.
func (r *ProductRepo) List(id entity.ID) ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select * from product where category_id= ? and is_delete = false`)
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
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID, &p.IsDeleted)
		result = append(result, &p)
	}
	return result, nil
}

// Create a product.
func (r *ProductRepo) Create(e *entity.Product) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
	insert into product(id,name,price,description,quantity_sold,available_units,create_at,category_id,id_deleted) 
	value (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return e.ProductID, err
	}
	_, err = stmt.Exec(
		e.ProductID,
		e.Name,
		e.Price,
		e.Description,
		e.QuantitySold,
		e.AvailableUnits,
		time.Now().Format("2022-09-26"),
		time.Now().Format("2022-09-26"),
		e.CategoryID,
		e.IsDeleted,
	)
	return entity.NewID(), nil
}

// Update a product.
func (r *ProductRepo) Update(e *entity.Product) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update product set name = ?, price = ?, description = ?, quantity_sold = ?,available_units =?, updated_at = ? where id = ?", e.Name, e.Price, e.Description, e.QuantitySold, e.AvailableUnits, e.UpdatedAt.Format("2022-09-26"), e.ProductID)
	if err != nil {
		return err
	}
	return nil
}

// Delete a product.
func (r *ProductRepo) Delete(id entity.ID) error {
	_, err := r.db.Exec("update product set id_deleted = true where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
