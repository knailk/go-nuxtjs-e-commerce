package entity

import (
	"time"
)

// User data.
type User struct {
	UserId    ID        `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"id_deleted"`
}

type Role string

const (
	ADMIN    Role = "admin"
	CUSTOMER Role = "customer"
	SELLER   Role = "seller"
)
