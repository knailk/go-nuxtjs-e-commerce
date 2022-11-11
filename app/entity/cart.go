package entity

type Cart struct {
	UserId ID `json:"userId"`
	ProductId ID `json:"productId"`
	Quantity int64 `json:"quantity" validate:"required"`
}
