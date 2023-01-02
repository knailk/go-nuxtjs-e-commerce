package usecase

import (
	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
)

// AuthUsecase interface.
type AuthUsecase interface {
	SignUp(user *entity.User) error
	SignIn(email string) (*entity.User, error)
	//Logout(userID entity.ID) error
}

//UseCase interface.
type CategoryUsecase interface {
	GetCategory(id int64) (*entity.Category, error)
	ListCategories() ([]*entity.Category, error)
}

//UseCase interface.
type ProductUsecase interface {
	GetProduct(id entity.ID) (*entity.Product, error)
	SearchProductsByQuery(query string) ([]*entity.Product, error)
	TopProduct() ([]*entity.Product, error)
	SearchProducts(query string) ([]*entity.Product, error)
	ListProducts(id int64) ([]*entity.Product, error)
	AdminGetProducts(id int64) ([]*entity.Product, error)
	CreateProduct(e *entity.Product) (entity.ID, error)
	UpdateProduct(id entity.ID, name string, price int64, description string, availableUnits int64, quantitySold int64) error
	DeleteProduct(id entity.ID) error
	AdminDeleteProduct(id entity.ID) error
}

//UseCase interface
type UserUsecase interface {
	GetUser(id entity.ID) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	ListUsers(filter string) ([]*entity.User, error)
	CreateUser(user *entity.User) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
	AdminDeleteUser(id entity.ID) error
}

//CartUsecase interface
type CartUsecase interface {
	GetCart(email string) ([]*entity.ProductCart, int64, error)
	AddToCart(productId entity.ID, email string, quantity int64) error
	RemoveProduct(productId entity.ID, email string, quantity int64) error
}

type AddressUsecase interface {
	GetAddress(email string) (*entity.Address, error)
	AddAddress(address *entity.Address) error
}
