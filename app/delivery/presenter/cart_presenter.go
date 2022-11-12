package presenter

import "github.com/knailk/go-shopee/app/entity"

//Product data present
type Cart struct {
	ListProducts []entity.ProductCart `json:"listProducs"`
	TotalPrice   int64                `json:"totalPrice"`
}
