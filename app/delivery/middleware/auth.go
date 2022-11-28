package middleware

import (
	"fmt"
	"net/http"
	"strings"
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
	claims["exp"] = time.Now().Add(time.Second * 60).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	return tokenString, err
}
func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] != nil {
			fmt.Println(r.Header["Authorization"])

			keyFunc := func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized user"))
				}
				return []byte(config.SECRET_KEY), nil
			}
			token, err := jwt.Parse(strings.Split(r.Header["Authorization"][0], "Bearer ")[1],keyFunc)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			fmt.Println("2222222")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "Admin" {
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

// func SellerIndex(w http.ResponseWriter, r *http.Request) {
// 	if r.Header.Get("Role") != "Seller" {
// 		w.Write([]byte("Not Authorized."))
// 		return
// 	}
// 	w.Write([]byte("Welcome, Seller."))
// }
