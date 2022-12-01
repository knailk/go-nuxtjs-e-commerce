package usecase

import (
	"errors"
	"fmt"

	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
	"github.com/knailk/go-nuxtjs-e-commerce/repository"
)

type CartService struct {
	dao repository.DAO
}

func NewCartService(dao repository.DAO) CartUsecase {
	return &CartService{
		dao: dao,
	}
}

// GetCart return a cart presentation
func (s *CartService) GetCart(email string) ([]*entity.ProductCart, int64, error) {
	//check user
	_, err := s.dao.NewCartRepo().GetAll(email)
	if err != nil {
		return nil,0, err
	}
	cartProduct, err := s.dao.NewCartRepo().GetAll(email)
	if err != nil {
		return nil, 0, err
	}

	var totalPrice int64 = 0
	for _, v := range cartProduct {
		totalPrice += v.Price * v.Quantity
	}
	return cartProduct,totalPrice, nil
}

// AddToCart add a product to cart
func (s *CartService) AddToCart(productId entity.ID, email string, quantity int64) error {
	u, err := s.dao.NewUserRepo().GetByEmail(email)
	if err != nil {
		return err
	}
	p, err := s.dao.NewProductRepo().Get(productId)
	if err != nil {
		return err
	}
	item, err := s.dao.NewCartRepo().GetOne(u.UserId, productId)
	if err != nil {
		return err
	}
	fmt.Println(item)
	if item.Quantity == 0 {
		if quantity > p.AvailableUnits {
			return errors.New("not enough product available")
		}
		newCart := entity.Cart{
			UserId: u.UserId,
			ProductId: productId,
			Quantity: quantity,
		}
		return s.dao.NewCartRepo().Add(&newCart)
	} else {
		item.Quantity += quantity
		if item.Quantity > p.AvailableUnits {
			return errors.New("not enough product available")
		}
		return s.dao.NewCartRepo().Update(item)
	}
}

// RemoveProduct remove a product in cart
func (s *CartService) RemoveProduct(productId entity.ID, email string, quantity int64) error {
	u, err := s.dao.NewUserRepo().GetByEmail(email)
	if err != nil {
		return err
	}
	_, err = s.dao.NewProductRepo().Get(productId)
	if err != nil {
		return err
	}
	item, err := s.dao.NewCartRepo().GetOne(u.UserId, productId)
	if err != nil {
		return err
	}
	if quantity == item.Quantity || item.Quantity == 1 {
		s.dao.NewCartRepo().Remove(item)
	} else {
		item.Quantity--
		s.dao.NewCartRepo().Decrease(item)
	}
	return nil
}
