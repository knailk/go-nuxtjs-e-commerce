package entity

import "time"

type Product struct {
	ProductID      ID
	Name           string
	Price          int64
	Description    string
	QuantitySold   int64
	AvailableUnits int64
	CreatedAt      time.Time
	UpdatedAt       time.Time
	CategoryID     ID
}
