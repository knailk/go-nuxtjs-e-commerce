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

	return tokenString, err
}
func verifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					writer.WriteHeader(http.StatusUnauthorized)
					_, err := writer.Write([]byte("You're Unauthorized!"))
					if err != nil {
						return nil, err

					}
				}
				return "", nil

			})
			if err != nil {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err2 := writer.Write([]byte("You're Unauthorized due to error parsing the JWT"))
				if err2 != nil {
					return
				}
			}
			if token.Valid {
				endpointHandler(writer, request)
				  } else {
						  writer.WriteHeader(http.StatusUnauthorized)
						  _, err := writer.Write([]byte("You're Unauthorized due to invalid token"))
						  if err != nil {
								  return
						  }
}
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
