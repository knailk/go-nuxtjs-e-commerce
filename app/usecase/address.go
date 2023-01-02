package usecase

import (
	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
	"github.com/knailk/go-nuxtjs-e-commerce/repository"
)

type AddressService struct {
	dao repository.DAO
}

func NewAddressService (dao repository.DAO) AddressUsecase{
	return &AddressService{
		dao:dao,
	}
}

func (s *AddressService) GetAddress(email string) (*entity.Address, error){
	return s.dao.NewAddressRepo().Get(email)
}

func (s *AddressService) AddAddress(address *entity.Address) error {
	_, err := s.dao.NewAddressRepo().Get(address.Email)
	if err != nil && err != entity.ErrNotFound{
		return err
	} else if err == entity.ErrNotFound{
		return s.dao.NewAddressRepo().Add(address)
	} else {
		return s.dao.NewAddressRepo().Update(address)
	}
}