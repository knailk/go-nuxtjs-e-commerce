package entity

import "time"

const ProductTableName = "product"

// Product data.
type Product struct {
	ProductID      ID     `json:"id"`
	Name           string `json:"name"`
	Price          int64  `json:"price"`
	Description    string `json:"description"`
	QuantitySold   int64  `json:"quantitySold"`
	AvailableUnits int64  `json:"availableUnits"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	CategoryID     ID     `json:"categoryId"`
	IsDeleted      bool   `json:"isDeleted"`
}

//NewProduct create a new product with new ID
func NewProduct(name string, price int64, description string, quantitySold int64, availableInits int64, category ID) *Product {
	return &Product{
		ProductID:      NewID(),
		Name:           name,
		Price:          price,
		Description:    description,
		QuantitySold:   quantitySold,
		AvailableUnits: availableInits,
		CreatedAt:      time.Now().Format(time.RFC3339),
		UpdatedAt:      time.Now().Format(time.RFC3339),
		CategoryID:     category,
		IsDeleted:      false,
	}
	
}
