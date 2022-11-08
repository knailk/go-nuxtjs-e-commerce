package usecase

import (
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
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
	_,err :=s.dao.NewUserRepo().Create(user)
	return err
}
func (s *AuthService) SignIn(email string) (*entity.User, error) {
	return s.dao.NewUserRepo().GetUserByEmail(email)
}

// func (s *AuthService) Logout(userID entity.ID) error {
// 	return s.dao.NewAuthRepo().Logout(userID)
// }
