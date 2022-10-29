package user

import (
	"github.com/knailk/go-shopee/app/entity"
)

//Reader interface
type Reader interface {
	//get user by id
	Get(id entity.ID) (*entity.User, error)
	//search list User by query
	Search(query string) ([]*entity.User, error)
	//get list User
	List() ([]*entity.User, error)
}

//Writer interface
type Writer interface {
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type Usecase interface{
	GetUser(id entity.ID) (*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(user entity.User) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
}