package usecase

import "github.com/knailk/go-shopee/app/entity"

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
	SearchProducts(query string) ([]*entity.Product, error)
	ListProducts(id int64) ([]*entity.Product, error)
	CreateProduct(e *entity.Product) (entity.ID, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id entity.ID) error
}

//UseCase interface
type UserUsecase interface{
	GetUser(id entity.ID) (*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(user *entity.User) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
}

//CartUsecase interface
type CartUsecase interface{
	GetCart(userId entity.ID) ([]entity.ProductCart,int64, error)
	AddToCart(cart *entity.Cart) error
	UpdateCart(cart *entity.Cart) error
	RemoveProduct(userId entity.ID,productId entity.ID) error
}