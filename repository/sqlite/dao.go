package sqlite

import (
	"database/sql"

	"github.com/knailk/go-shopee/repository"
)

type DAO interface {
	NewAuthRepo() repository.AuthQuery
	NewCategoryRepo() repository.CategoryQuery
	NewUserRepo() repository.UserQuery
	NewProductRepo() repository.ProductQuery
}
type dao struct {
	db *sql.DB
}

func NewDAO(db *sql.DB) DAO {
	return &dao{
		db: db,
	}
}

// NewProductRepo create an implementation of product repository.
func (d *dao) NewCategoryRepo() repository.CategoryQuery {
	return &CategoryRepo{
		db: d.db,
	}
}

// NewProductRepo create an implementation of product repository.
func (d *dao) NewProductRepo() repository.ProductQuery {
	return &ProductRepo{
		db: d.db,
	}
}

// NewUserRepo create an implementation of product repository.
func (d *dao) NewUserRepo() repository.UserQuery {
	return &UserRepo{
		db: d.db,
	}
}

// NewProductRepo create an implementation of product repository.
func (d *dao) NewAuthRepo() repository.AuthQuery {
	return &AuthRepo{
		db: d.db,
	}
}
