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

func (s *CartService) GetCart(userId entity.ID) ([]*entity.Cart, error){
	//check user
	_, err := s.dao.NewUserRepo().Get(userId)
	if err != nil {
		return nil, err
	}
	return s.dao.NewCartRepo().GetAll(userId)
}

func (s *CartService) AddToCart(cart *entity.Cart) error {
	item, err := s.dao.NewCartRepo().GetOne(cart.UserId,cart.ProductId)
	if err != nil{
		return err
	}
	if item.Quantity == 0 {
		return s.dao.NewCartRepo().Add(cart)
	} else{
		item.Quantity += cart.Quantity
		return s.dao.NewCartRepo().Update(item)
	}
}

func (s *CartService) UpdateCart(productId entity.ID) error{
	return nil
}

func (s *CartService) RemoveProduct(productId entity.ID) error{
	return nil
}