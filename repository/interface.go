package repository

import "github.com/knailk/go-nuxtjs-e-commerce/app/entity"

type DAO interface {
	//NewAuthRepo() AuthQuery
	NewCategoryRepo() CategoryQuery
	NewUserRepo() UserQuery
	NewProductRepo() ProductQuery
	NewCartRepo() CartQuery
}

//Repository interface
type UserQuery interface {
	Get(id entity.ID) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	List() ([]*entity.User, error)
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
	Create(e *entity.Product) (entity.ID, error)
	Update(e *entity.Product) error
	Delete(id entity.ID) error
}

//CategoryQuery interface.
type CategoryQuery interface {
	Get(id int64) (*entity.Category, error)
	List() ([]*entity.Category, error)
}

//CartQuery interface
type CartQuery interface {
	GetAll(userId entity.ID) ([]*entity.Cart, error)
	GetOne(userId entity.ID, productId entity.ID) (*entity.Cart, error)
	Add(cart *entity.Cart) error
	Update(cart *entity.Cart) error
	Remove(userId entity.ID, productId entity.ID) error
}

// //AuthQuery interface
// type AuthQuery interface {
// 	//SignUp(user *entity.User) error
// 	SignIn(email string) (*entity.User, error)
// 	Logout(userID entity.ID) error
// }
