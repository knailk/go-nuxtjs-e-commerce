package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/delivery/middleware"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase"
	"golang.org/x/crypto/bcrypt"
)

// signIn login user
func signIn(service usecase.AuthUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Email    string `json:"email" validate:"required"`
			Password string `json:"password" validate:"required"`
		}
		validate := validator.New()
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		//validate
		if err := validate.Struct(input); err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		//get user
		authUser, err := service.SignIn(input.Email)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(input.Password)); err != nil {
			logInternalServerError(err, "incorrect password", w)
			return
		}
		validToken, err := middleware.GenerateJWT(authUser.Email, authUser.Role)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		var token struct {
			Role        entity.Role `json:"role"`
			Email       string      `json:"email"`
			TokenString string      `json:"token"`
		}
		token.Email = authUser.Email
		token.Role = authUser.Role
		token.TokenString = validToken
		
		if err := json.NewEncoder(w).Encode(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
		}
	})
}

// signUp register new user
func signUp(service usecase.AuthUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Email    string `json:"email" validate:"required"`
			Password string `json:"password" validate:"required"`
			FullName string `string:"fullName" validate:"required"`
			Phone    string `string:"phone" validate:"required"`
			Gender   string `string:"gender" validate:"required,oneof=Male Female"`
		}
		//validate
		validate := validator.New()
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		if err := validate.Struct(input); err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		//create user
		user, err := entity.NewUser(input.Email, input.Password, input.FullName, input.Gender, input.Phone, "Customer")
		if err != nil {
			if err != nil {
				logInternalServerError(err, err.Error(), w)
				return
			}
		}
		err = service.SignUp(user)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		validToken, err := middleware.GenerateJWT(user.Email, user.Role)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		var token struct {
			Role        entity.Role `json:"role"`
			Email       string      `json:"email"`
			TokenString string      `json:"token"`
		}
		token.Email = user.Email
		token.Role = user.Role
		token.TokenString = validToken

		if err := json.NewEncoder(w).Encode(token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
		}
	})
}

func MakeAuthHandlers(r *mux.Router, service usecase.AuthUsecase) {

	r.Handle("/signin", signIn(service)).Methods(http.MethodPost)

	r.Handle("/signup", signUp(service)).Methods(http.MethodPost)
}
