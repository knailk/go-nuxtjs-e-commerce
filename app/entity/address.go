package entity

import "time"

// Address information for shipper data.
type Address struct {
	Email     string `json:"email"`
	Country   string `json:"country"`
	City      string `json:"city"`
	District  string `json:"district"`
	Ward      string `json:"ward"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

//NewAddress create a new product with new ID
func NewAddress(email string, country string, city string, district string, ward string, address string) *Address {
	return &Address{
		Email:     email,
		Country:   country,
		City:      city,
		District:  district,
		Ward:      ward,
		Address:   address,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}
