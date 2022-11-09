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

func signIn(service usecase.AuthUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error login"
		var input struct {
			UserName string `json:"userName" validate:"required"`
			Password string `json:"password" validate:"required"`
		}
		validate := validator.New()
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logError(err, errorMessage, w)
			return
		}
		//validate
		if err := validate.Struct(input); err != nil {
			logError(err, errorMessage, w)
			return
		}
		//get user
		authUser,err := service.SignIn(input.UserName)
		if err != nil {
			logError(err, errorMessage, w)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(input.Password)); err != nil {
			logError(err, "incorrect password", w)
			return
		}
		validToken, err := middleware.GenerateJWT(authUser.Email, authUser.Role)
		if err != nil {
			logError(err, errorMessage, w)
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
			w.Write([]byte(errorMessage))
		}
	})
}

func signUp(service usecase.AuthUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error sign up"
		var input struct {
			Email string `json:"email" validate:"required"`
			Password string `json:"password" validate:"required"`
			Name string `string:"name" validate:"required"`
		}
		//validate
		validate := validator.New()
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logError(err, errorMessage, w)
			return
		}
		if err := validate.Struct(input); err != nil {
			logError(err, errorMessage, w)
			return
		}
		//create user
		user,err := entity.NewUser(input.Email,input.Password,input.Name,"","","Customer")
		if err != nil {
			if err != nil {
				logError(err, errorMessage, w)
				return
			}
		}
		err = service.SignUp(user)
		if err != nil {
			logError(err, errorMessage, w)
			return
		}
		validToken, err := middleware.GenerateJWT(user.Email, user.Role)
		if err != nil {
			logError(err, errorMessage, w)
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
			w.Write([]byte(errorMessage))
		}
	})
}

func MakeAuthHandlers(r *mux.Router, service usecase.AuthUsecase) {

	r.Handle("/signin", signIn(service)).Methods(http.MethodPost)

	r.Handle("/signup", signUp(service)).Methods(http.MethodPost)
}
