package usecase

import (
	"errors"

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

//GetCart return a cart presentation
func (s *CartService) GetCart(userId entity.ID) ([]entity.ProductCart, int64, error) {
	//check user
	_, err := s.dao.NewUserRepo().Get(userId)
	if err != nil {
		return nil, 0, err
	}
	cart, err := s.dao.NewCartRepo().GetAll(userId)
	if err != nil {
		return nil, 0, err
	}

	var cartProduct []entity.ProductCart
	var totalPrice int64 = 0
	for _, v := range cart {
		p, err := s.dao.NewProductRepo().Get(v.ProductId)
		if err != nil {
			return nil, 0, err
		}
		cartProduct = append(cartProduct, entity.ProductCart{Name: p.Name, Price: p.Price, Quantity: v.Quantity})
		totalPrice += p.Price * v.Quantity
	}
	return cartProduct, totalPrice, nil
}

//AddToCart add a product to cart
func (s *CartService) AddToCart(cart *entity.Cart) error {
	_, err := s.dao.NewUserRepo().Get(cart.UserId)
	if err != nil {
		return err
	}
	p, err := s.dao.NewProductRepo().Get(cart.ProductId)
	if err != nil {
		return err
	}
	item, err := s.dao.NewCartRepo().GetOne(cart.UserId, cart.ProductId)
	if err != nil {
		return err
	}
	if item.Quantity == 0 {
		if cart.Quantity > p.AvailableUnits{
			return errors.New("not enough product available")
		}
		return s.dao.NewCartRepo().Add(cart)
	} else {
		item.Quantity += cart.Quantity
		if item.Quantity > p.AvailableUnits{
			return errors.New("not enough product available")
		}
		return s.dao.NewCartRepo().Update(item)
	}
}

func (s *CartService) UpdateCart(cart *entity.Cart) error {
	return s.dao.NewCartRepo().Update(cart)
}

//RemoveProduct remove a product in cart
func (s *CartService) RemoveProduct(userId,productId entity.ID) error {
	_, err := s.dao.NewUserRepo().Get(userId)
	if err != nil {
		return err
	}
	_, err = s.dao.NewProductRepo().Get(productId)
	if err != nil {
		return err
	}
	s.dao.NewCartRepo().Remove(userId,productId)
	return nil
}
