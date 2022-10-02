package entity

import "time"

//credit card data
type Card struct {
	CardId     ID
	Name       string
	CardNumber string
	ExpireTime time.Time
}