package presenter

import "github.com/knailk/go-nuxtjs-e-commerce/app/entity"

//Product data present
type Product struct {
	ProductId      entity.ID `json:"id"`
	Name           string    `json:"name"`
	Price          int64     `json:"price"`
	Description    string    `json:"description"`
	QuantitySold   int64     `json:"quantitySold"`
	AvailableUnits int64     `json:"availableUnits"`
	Image          string    `json:"image"`
	CreatedAt      string    `json:"createdAt"`
	UpdatedAt      string    `json:"updatedAt"`
	Category       int       `json:"category"`
	IsDeleted      bool      `json:"isDeleted"`
}

type Category struct {
	CategoryId    entity.ID `json:"id"`
	Name          string    `json:"name"`
	NumberProduct int       `json:"numberProduct"`
}
