package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/knailk/go-shopee/app/config"
	"github.com/knailk/go-shopee/app/entity"
)

// GenerateJWT return a token string
func GenerateJWT(email string, role entity.Role) (string, error) {
	var mySigningKey = []byte(config.SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}

func CustomerIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "Customer" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, Customer."))
}

func SellerIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "Seller" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, Seller."))
}
