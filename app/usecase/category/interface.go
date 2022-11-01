package category

import (
	"github.com/knailk/go-shopee/app/entity"
)

//Repository interface.
type Repository interface {
	Get(id int64) (*entity.Category, error)
	List() ([]*entity.Category, error)
}

//UseCase interface.
type Usecase interface {
	GetCategory(id int64) (*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
}
