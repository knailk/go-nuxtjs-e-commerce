package usecase

import (
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/repository"
)

type CartService struct {
	dao repository.DAO
}

func NewCartService(dao repository.DAO) CartUsecase {
	return &CartService{
		dao: dao,
	}
}

func (s *CartService) AddProduct(productId entity.ID) error {
	return nil
}

func (s *CartService) RemoveProduct(productId entity.ID) error{
	return nil
}

func (s *CartService) UpdateQuantity(productId entity.ID) error
