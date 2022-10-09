package entity

import (
	"time"
)

//user data
type User struct {
	UserId	ID `gorm:"primary_key"`
	Email     string
	Password  string
	Name string
	Gender string
	Role Role
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role string

const (
	ADMIN Role = "admin"
	CUSTOMER  Role = "customer"
	SELLER Role = "seller"
)
