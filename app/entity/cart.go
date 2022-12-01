package entity

type Cart struct {
	UserId    ID    `json:"userId"`
	ProductId ID    `json:"productId"`
	Quantity  int64 `json:"quantity" validate:"required"`
}

type ProductCart struct {
	ProductId      ID     `json:"productId"`
	UserId         ID     `json:"userId"`
	Name           string `json:"name"`
	Price          int64  `json:"price"`
	Image          string `json:"image"`
	Quantity       int64  `json:"quantity"`
	AvailableUnits int64  `json:"availableUnits"`
}
