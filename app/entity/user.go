package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User data.
type User struct {
	UserId    ID     `json:"id"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password"`
	Name      string `json:"name" validate:"required"`
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

//NewUser create a new user with new ID
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
		IsDeleted: false,
	}
	pwd, err := GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	return u, nil
}

//GeneratePassword return strirng hash of the password
func GeneratePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// func (u *User) Validate() error {
// 	//validate email
// 	_, err := mail.ParseAddress(u.Email)
// 	if err != nil {
// 		return err
// 	}
// 	//validate 
// 	errs.Add
// 	return nil
// }
