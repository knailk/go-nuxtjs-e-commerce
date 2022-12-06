package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/knailk/go-nuxtjs-e-commerce/app/delivery/middleware"
	"github.com/knailk/go-nuxtjs-e-commerce/app/delivery/presenter"
	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
	"github.com/knailk/go-nuxtjs-e-commerce/app/usecase"
	"golang.org/x/crypto/bcrypt"
)

// listUsers return http handler
func listUsers(service usecase.UserUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		var data []*entity.User
		var err error
		query := r.URL.Query().Get("query")
		filter := r.URL.Query().Get("filter")
		switch {
		case query == "":
			data, err = service.ListUsers(filter)
		default:
			data, err = service.SearchUsers(query)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("data not found"))
			return
		}
		var toJ []*presenter.User
		for _, d := range data {
			toJ = append(toJ, &presenter.User{
				ID:        d.UserId,
				Email:     d.Email,
				Name:      d.Name,
				Phone:     d.Phone,
				Gender:    d.Gender,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
				IsDeleted: d.IsDeleted,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

// createUser create new user
func createUser(service usecase.UserUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error adding user"
		var input struct {
			Email    string      `json:"email" validate:"required,email"`
			Password string      `json:"password" validate:"omitempty,min=8,passwd"`
			Name     string      `json:"name" validate:"required,min=2,max=50"`
			Gender   string      `json:"gender" validate:"omitempty,oneof=Male Female"`
			Phone    string      `json:"phone" validate:"omitempty,min=9,max=11"`
			Role     entity.Role `json:"role" validate:"omitempty,oneof=Admin Customer"`
		}
		validate := validator.New()
		_ = validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
			return len(fl.Field().String()) > 8 && len(fl.Field().String()) < 20
		})

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err := validate.Struct(input); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		e, err := entity.NewUser(input.Email, input.Password, input.Name, input.Gender, input.Phone, input.Role)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		id, err := service.CreateUser(e)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

// currUser get user data who logged in
func currUser(service usecase.UserUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		clams, _ := middleware.ExtractClaims(w, r)
		strValue := fmt.Sprintf("%v", clams["email"])
		data, err := service.GetUserByEmail(strValue)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("data not found"))
			return
		}
		user := &presenter.User{
			ID:        data.UserId,
			Email:     data.Email,
			Name:      data.Name,
			Phone:     data.Phone,
			Gender:    data.Gender,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

// getUser get user by id
func getUser(service usecase.UserUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error get user by id"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		data, err := service.GetUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("data not found"))
			return
		}
		toJ := &presenter.User{
			ID:        data.UserId,
			Email:     data.Email,
			Name:      data.Name,
			Phone:     data.Phone,
			Gender:    data.Gender,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

// deleteUser delete a user
func deleteUser(service usecase.UserUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Id       string `json:"id" validate:"required"`
			Email    string `json:"email" validate:"required"`
			Password string `json:"password" validate:"required"`
		}
		validate := validator.New()
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}
		//validate
		if err := validate.Struct(input); err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		id, err := entity.StringToID(input.Id)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		user, err := service.GetUserByEmail(input.Email)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("incorrect password"))
			return
		}
		err = service.DeleteUser(id)
		if err != nil && err != entity.ErrNotFound {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

// deleteUser delete a user
func adminDeleteUser(service usecase.UserUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
			return
		}
		err = service.AdminDeleteUser(id)
		if err != nil && err != entity.ErrNotFound {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

func updateUser(service usecase.UserUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error update user"
		var input struct {
			UserID string `json:"id" validate:"required"`
			Email  string `json:"email" validate:"required"`
			Name   string `json:"name" validate:"required,min=2,max=50"`
			Gender string `json:"gender" validate:"required,oneof=Male Female"`
			Phone  string `json:"phone" validate:"required,min=9,max=11"`
		}
		validate := validator.New()
		_ = validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
			return len(fl.Field().String()) > 8 && len(fl.Field().String()) < 20
		})
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		if err := validate.Struct(input); err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		fmt.Println(input)
		id, err := entity.StringToID(input.UserID)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		u := entity.User{
			UserId: id,
			Email:  input.Email,
			Name:   input.Name,
			Gender: input.Gender,
			Phone:  input.Phone,
		}
		err = service.UpdateUser(&u)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err := json.NewEncoder(w).Encode("update user successful"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func logInternalServerError(err error, errorMessage string, w http.ResponseWriter) {
	log.Println(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errorMessage))
}

func MakeUserHandlers(r *mux.Router, service usecase.UserUsecase) {
	r.Handle("/admin/user", listUsers(service)).Methods(http.MethodGet)

	r.Handle("/admin/user/me", currUser(service)).Methods(http.MethodGet)

	r.Handle("/admin/user", createUser(service)).Methods(http.MethodPost)

	r.Handle("/admin/user/{id}", getUser(service)).Methods(http.MethodGet)

	r.Handle("/user", deleteUser(service)).Methods(http.MethodDelete)

	r.Handle("/admin/user/{id}", adminDeleteUser(service)).Methods(http.MethodDelete)

	r.Handle("/admin/user/edit", updateUser(service)).Methods(http.MethodPost)
}
