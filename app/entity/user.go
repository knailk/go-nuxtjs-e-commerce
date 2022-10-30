package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User data.
type User struct {
	UserId    ID     `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Role      Role   `json:"role"`
	CreatedAt string `json:"createAt"`
	UpdatedAt string `json:"updateAt"`
	IsDeleted bool   `json:"isDelete"`
}

type Role string

const (
	ADMIN    Role = "Admin"
	CUSTOMER Role = "Customer"
	SELLER   Role = "Seller"
)

func NewUser(email, password, name, gender, phone string, role Role) (*User, error) {
	u := &User{
		UserId:    NewID(),
		Email:     email,
		Name:      name,
		Gender:    gender,
		Phone:     phone,
		Role:      role,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
	pwd, err := GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	return u, nil
}

func GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
