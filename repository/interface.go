package repository

import "github.com/knailk/go-shopee/app/entity"

//Repository interface
type UserQuery interface {
	Get(id entity.ID) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	List() ([]*entity.User, error)
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

//ProductQuery interface.
type ProductQuery interface {
	Get(id entity.ID) (*entity.Product, error)
	Search(query string) ([]*entity.Product, error)
	List(int64) ([]*entity.Product, error)
	Create(e *entity.Product) (entity.ID, error)
	Update(e *entity.Product) error
	Delete(id entity.ID) error
}

//CategoryQuery interface.
type CategoryQuery interface {
	Get(id int64) (*entity.Category, error)
	List() ([]*entity.Category, error)
}

//AuthQuery interface
type AuthQuery interface {
	SignUp(user *entity.User) (int64, error)
	SignIn(email string) (*entity.User, error)
	Logout(userID entity.ID) error
}



