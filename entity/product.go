package entity

import "time"

const ProductTableName = "product"
type Product struct {
	ProductID      ID        `json:"id"`
	Name           string    `json:"name"`
	Price          int64     `json:"price"`
	Description    string    `json:"description"`
	QuantitySold   int64     `json:"quantity_sold"`
	AvailableUnits int64     `json:"available_units"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CategoryID     ID        `json:"category_id"`
	IsDeleted      bool      `json:"id_deleted"`
}
