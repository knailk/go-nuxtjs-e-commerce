package entity

import "time"

// Order data.
type Order struct {
	OrderID         ID
	DeliveryAddress ID
	OrderDate       time.Time
	TotalPrice      int64
	Status          string
	DeliverDate     time.Time
	ShipFee         int64
	Product []ID
}
