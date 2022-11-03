package presenter

import "github.com/knailk/go-shopee/app/entity"

//Product data present
type Product struct {
	ProductID      entity.ID `json:"id"`
	Name           string    `json:"name"`
	Price          int64     `json:"price"`
	Description    string    `json:"description"`
	QuantitySold   int64     `json:"quantitySold"`
	AvailableUnits int64     `json:"availableUnits"`
	CreatedAt      string    `json:"createdAt"`
	UpdatedAt      string    `json:"updatedAt"`
	Category       string    `json:"category"`
}