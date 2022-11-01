package sqlite

import (
	"database/sql"

	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase/category"
)

type CategoryRepo struct {
	db *sql.DB
}

// NewProductRepo create an implementation of product repository.
func NewCategoryRepo(db *sql.DB) category.Repository {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) Get(id int64) (*entity.Category, error){
	stmt, err := r.db.Prepare(`select name from category where id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	var u entity.Category
	for rows.Next() {
		err = rows.Scan(&u.CategoryName)
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *CategoryRepo) 	List() ([]*entity.Category, error){
	stmt, err := r.db.Prepare(`select* from category`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var result []*entity.Category
	for rows.Next() {
		var u entity.Category
		rows.Scan(&u.CategoryId, &u.CategoryName)
		result = append(result, &u)
	}
	return result, nil
}