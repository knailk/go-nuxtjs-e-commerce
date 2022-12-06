package presenter

import "github.com/knailk/go-nuxtjs-e-commerce/app/entity"

//User data present
type User struct {
	ID        entity.ID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	IsDeleted  bool      `json:"isDeleted"`
}
