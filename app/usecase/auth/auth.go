package auth

import (
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)

// UseCase interface.
type Usecase interface {
	SignUp(user *entity.User) (int64, error)
	SignIn(email string) (string, error)
	Logout(userID entity.ID) error
}

// Service product usecase.
type Service struct {
	repo repository.AuthQuery
}

// NewService create new service.
func NewService(r repository.AuthQuery) *Service {
	return &Service{
		repo: r,
	}
}
func (s *Service) SignUp(user *entity.User) (int64, error) {
	return s.repo.SignUp(user)
}
func (s *Service) SignIn(email string) (*entity.User, error) {
	return s.repo.SignIn(email)
}

func (s *Service) Logout(userID entity.ID) error {
	return s.repo.Logout(userID)
}
