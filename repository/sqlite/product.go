package sqlite

import (
	"database/sql"
	"time"

	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
)

type ProductRepo struct {
	db *sql.DB
}

// Get product by id.
func (r *ProductRepo) Get(id entity.ID) (*entity.Product, error) {
	stmt, err := r.db.Prepare(`select id,name,price,description,quantitySold,availableUnits,image,createdAt,updatedAt,categoryId from product where id=? and isDeleted = 0`)
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
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
	}
	if err != nil {
		return nil, err
	}
	if p.ProductID == 0 {
		return nil, entity.ErrNotFound
	}
	return &p, nil
}

// SearchByQuery get list product by query
func (r *ProductRepo) SearchByQuery(query string) ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`
	select id,name,price,description,quantitySold,availableUnits,image,createdAt,updatedAt,categoryId 
	from product 
	where name like ? and isDeleted = 0
	order by name ASC
	LIMIT 10`)
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
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
		result = append(result, &p)
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Top return limit top product.
func (r *ProductRepo) Top() ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(
		`Select id,name,price,description,quantitySold,availableUnits,image,createdAt,updatedAt,categoryId 
		from product 
		where isDeleted = 0
		order by quantitySold DESC
		LIMIT 8`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var result []*entity.Product
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p entity.Product
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
		result = append(result, &p)
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Search list product by query.
func (r *ProductRepo) Search(query string) ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select id,name,price,description,quantitySold,availableUnits,image,createdAt,updatedAt,categoryId from product where name like ? and isDeleted = 0`)
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
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		result = append(result, &p)
	}
	return result, nil
}

// List product by category id.
func (r *ProductRepo) List(id int64) ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select id,name,price,description,quantitySold,availableUnits,image,createdAt,updatedAt,categoryId from product where categoryID = ? and isDeleted = 0`)
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
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.Image, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID)
		if err != nil {
			return nil, err
		}
		result = append(result, &p)
	}
	return result, nil
}
// ListAll product by category id include isdelete = true
func (r *ProductRepo) ListAll(id int64) ([]*entity.Product, error) {
	stmt, err := r.db.Prepare(`Select * from product where categoryID = ?`)
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
		err = rows.Scan(&p.ProductID, &p.Name, &p.Price, &p.Description, &p.QuantitySold, &p.AvailableUnits, &p.CreatedAt, &p.UpdatedAt, &p.CategoryID, &p.IsDeleted, &p.Image,)
		if err != nil {
			return nil, err
		}
		result = append(result, &p)
	}
	return result, nil
}
// Create a product.
func (r *ProductRepo) Create(e *entity.Product) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
	insert into product(id, name, price, description, quantitySold, availableUnits,image, createdAt, updatedAt, categoryId, isDeleted) 
	values (?,?,?,?,?,?,?,?,?,?,?)`)
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
		e.Image,
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
func (r *ProductRepo) Update(id entity.ID, name string, price int64, description string, availableUnits, quantitySold int64) error {
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
		name,
		price,
		description,
		quantitySold,
		availableUnits,
		time.Now().Format(time.RFC3339),
		id,
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
	_, err := r.db.Exec("update product set isDeleted = NOT isDeleted where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
