package repository

import (
	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
)

type DAO interface {
	//NewAuthRepo() AuthQuery
	NewCategoryRepo() CategoryQuery
	NewUserRepo() UserQuery
	NewProductRepo() ProductQuery
	NewCartRepo() CartQuery
	NewAddressRepo() AddressQuery
}

//Repository interface
type UserQuery interface {
	Get(id entity.ID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	List(filter string) ([]*entity.User, error)
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

//ProductQuery interface.
type ProductQuery interface {
	Get(id entity.ID) (*entity.Product, error)
	SearchByQuery(query string) ([]*entity.Product,error)
	Top() ([]*entity.Product, error)
	Search(query string) ([]*entity.Product, error)
	List(int64) ([]*entity.Product, error)
	ListAll(int64) ([]*entity.Product, error)
	Create(e *entity.Product) (entity.ID, error)
	Update(id entity.ID, name string, price int64, description string, availableUnits, quantitySold int64) error
	Delete(id entity.ID) error
}

//CategoryQuery interface.
type CategoryQuery interface {
	Get(id int64) (*entity.Category, error)
	List() ([]*entity.Category, error)
}

//CartQuery interface
type CartQuery interface {
	GetAll(email string) ([]*entity.ProductCart, error)
	GetOne(userId entity.ID, productId entity.ID) (*entity.Cart, error)
	Add(cart *entity.Cart) error
	Update(cart *entity.Cart) error
	Remove(cart *entity.Cart) error
	Decrease(cart *entity.Cart) error
}

//AddressQuery interface.
type AddressQuery interface {
	Get(email string) (*entity.Address, error)
	Add(address *entity.Address) error
	Update(address *entity.Address) error
}
// //AuthQuery interface
// type AuthQuery interface {
// 	//SignUp(user *entity.User) error
// 	SignIn(email string) (*entity.User, error)
// 	Logout(userID entity.ID) error
// }
