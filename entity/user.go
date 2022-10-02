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
	Role int
	CreatedAt time.Time
	UpdatedAt time.Time
}
