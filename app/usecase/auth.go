package usecase

import (
	"errors"

	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
	"github.com/knailk/go-nuxtjs-e-commerce/repository"
)

// Service product usecase.
type AuthService struct {
	dao repository.DAO
}

// NewService create new service.
func NewAuthService(dao repository.DAO) AuthUsecase {
	return &AuthService{
		dao: dao,
	}
}
func (s *AuthService) SignUp(user *entity.User) error {
	u,err := s.dao.NewUserRepo().GetByEmail(user.Email)
	if err != nil && err != entity.ErrNotFound{
		return err
	}
	if u != nil {
		return errors.New("email had exist")
	}
	_, err = s.dao.NewUserRepo().Create(user)
	return err
}
func (s *AuthService) SignIn(email string) (*entity.User, error) {
	return s.dao.NewUserRepo().GetByEmail(email)
}

// func (s *AuthService) Logout(userID entity.ID) error {
// 	return s.dao.NewAuthRepo().Logout(userID)
// }
