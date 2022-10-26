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
func (s *Service) GetUser(id entity.ID) (*entity.User, error)
func (s *Service) SearchUsers(query string) ([]*entity.User, error)
func (s *Service) ListUsers() ([]*entity.User, error)
func (s *Service) CreateUser(email string, password string, name string, gender string, phone string, role entity.Role) (entity.ID, error)
func (s *Service) UpdateUser(e *entity.User) error
func (s *Service) DeleteUser(id entity.ID) error
