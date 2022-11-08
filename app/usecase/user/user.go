package user

import (
	"strings"
	"time"

	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)

//UseCase interface
type Usecase interface{
	GetUser(id entity.ID) (*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(user entity.User) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
}
type Service struct {
	repo repository.UserQuery
}

// create new service
func NewService(r repository.UserQuery) *Service {
	return &Service{
		repo: r,
	}
}


func (s *Service) GetUser(id entity.ID) (*entity.User, error){
	return s.repo.Get(id)
}
func (s *Service) SearchUsers(query string) ([]*entity.User, error){
	return s.repo.Search(strings.ToLower(query))
}
func (s *Service) ListUsers() ([]*entity.User, error){
	return s.repo.List()
}
func (s *Service) CreateUser(email string, password string, name string, gender string, phone string, role entity.Role) (entity.ID, error){
	e, err := entity.NewUser(email, password, name, gender,phone, role)
	if err != nil {
		return e.UserId, err
	}
	return s.repo.Create(e)
}
func (s *Service) UpdateUser(e *entity.User) error{
	e.UpdatedAt = time.Now().Format(time.RFC3339)
	return s.repo.Update(e)
}
func (s *Service) DeleteUser(id entity.ID) error{
	u, err := s.GetUser(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
