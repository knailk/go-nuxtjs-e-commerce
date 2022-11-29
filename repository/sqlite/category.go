package sqlite

import (
	"database/sql"

	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
)

type CategoryRepo struct {
	db *sql.DB
}

func (r *CategoryRepo) Get(id int64) (*entity.Category, error) {
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
	u.CategoryId = int(id)
	for rows.Next() {
		err = rows.Scan(&u.CategoryName)
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *CategoryRepo) List() ([]*entity.Category, error) {
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
