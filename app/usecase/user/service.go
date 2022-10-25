package user

import "github.com/knailk/go-shopee/app/entity"

type Service struct {
	repo Repository
}

// create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// TODO implement method of usecase interface
func GetUser(id entity.ID) (*entity.User, error)
func SearchUsers(query string) ([]*entity.User, error)
func ListUsers() ([]*entity.User, error)
func CreateUser(user entity.User) (entity.ID, error)
func UpdateUser(e *entity.User) error
func DeleteUser(id entity.ID) error