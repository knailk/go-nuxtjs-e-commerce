package usecase

import (
	"strings"
	"time"

	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)


type UserService struct {
	repo repository.UserQuery
}

// create new service
func NewUserService(r repository.UserQuery) UserUsecase {
	return &UserService{
		repo: r,
	}
}


func (s *UserService) GetUser(id entity.ID) (*entity.User, error){
	return s.repo.Get(id)
}

func (s *UserService) SearchUsers(query string) ([]*entity.User, error){
	return s.repo.Search(strings.ToLower(query))
}

func (s *UserService) ListUsers() ([]*entity.User, error){
	return s.repo.List()
}

func (s *UserService) CreateUser(e *entity.User) (entity.ID, error){
	
	return s.repo.Create(e)
}

func (s *UserService) UpdateUser(e *entity.User) error{
	e.UpdatedAt = time.Now().Format(time.RFC3339)
	return s.repo.Update(e)
}

func (s *UserService) DeleteUser(id entity.ID) error{
	u, err := s.GetUser(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}
