package usecase

import (
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)



// Service product usecase.
type AuthService struct {
	repo repository.AuthQuery
}

// NewService create new service.
func NewAuthService(r repository.AuthQuery) AuthUsecase {
	return &AuthService{
		repo: r,
	}
}
func (s *AuthService) SignUp(user *entity.User) (int64, error) {
	return s.repo.SignUp(user)
}
func (s *AuthService) SignIn(email string) (*entity.User, error) {
	return s.repo.SignIn(email)
}

func (s *AuthService) Logout(userID entity.ID) error {
	return s.repo.Logout(userID)
}
