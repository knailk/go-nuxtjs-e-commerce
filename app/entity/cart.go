package entity

type Cart struct {
	UserId ID `json:"userId"`
	ProductId ID `json:"productId"`
	Quantity int64 `json:"quantity" validate:"required"`
}

type ProductCart struct{
	Name string `json:"name"`
	Price int64 `json:"price"`
	Quantity int64 `json:"quantity"`
}
