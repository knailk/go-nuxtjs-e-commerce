package sqlite

import (
	"database/sql"

	"github.com/knailk/go-nuxtjs-e-commerce/repository"
)

type dao struct {
	db *sql.DB
}

func NewDAO(db *sql.DB) repository.DAO {
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

func (d *dao) NewCartRepo() repository.CartQuery {
	return &CartRepo{
		db: d.db,
	}
}

func(d *dao) NewAddressRepo() repository.AddressQuery{
	return &AddressRepo{
		db: d.db,
	}
}

// // NewProductRepo create an implementation of product repository.
// func (d *dao) NewAuthRepo() repository.AuthQuery {
// 	return &AuthRepo{
// 		db: d.db,
// 	}
// }
