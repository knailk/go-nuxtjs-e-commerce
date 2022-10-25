package entity

import "time"

//Card data.
type Card struct {
	CardId     ID
	Name       string
	CardNumber string
	ExpireTime time.Time
}