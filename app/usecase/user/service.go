package user

import (
	"strings"
	"time"

	"github.com/knailk/go-shopee/app/entity"
)

type Service struct {
	repo Repository
}

// create new service
func NewService(r Repository) *Service {
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
