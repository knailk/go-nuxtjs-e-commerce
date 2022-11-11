package usecase

import (
	"strings"
	"time"

	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)


type UserService struct {
	dao repository.DAO
}

// create new service
func NewUserService(dao repository.DAO) UserUsecase {
	return &UserService{
		dao: dao,
	}
}


func (s *UserService) GetUser(id entity.ID) (*entity.User, error){
	user, err := s.dao.NewUserRepo().Get(id)
	if user.UserId == 0{
		return nil, entity.ErrNotFound
	}
	return user, err
}

func (s *UserService) SearchUsers(query string) ([]*entity.User, error){
	return s.dao.NewUserRepo().Search(strings.ToLower(query))
}

func (s *UserService) ListUsers() ([]*entity.User, error){
	return s.dao.NewUserRepo().List()
}

func (s *UserService) CreateUser(e *entity.User) (entity.ID, error){
	
	return s.dao.NewUserRepo().Create(e)
}

func (s *UserService) UpdateUser(e *entity.User) error{
	e.UpdatedAt = time.Now().Format(time.RFC3339)
	return s.dao.NewUserRepo().Update(e)
}

func (s *UserService) DeleteUser(id entity.ID) error{
	u, err := s.GetUser(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.dao.NewUserRepo().Delete(id)
}
