package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/delivery/presenter"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase/user"

	"github.com/go-playground/validator/v10"
)

// listUsers return http handler
func listUsers(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var data []*entity.User
		var err error
		query := r.URL.Query().Get("query")
		switch {
		case query == "":
			data, err = service.ListUsers()
		default:
			data, err = service.SearchUsers(query)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
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
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

// createUser create new user
func createUser(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"
		var input struct {
			Email    string      `json:"email" validate:"required,email"`
			Password string      `json:"password" validate:"omitempty,min=8,passwd"`
			Name     string      `json:"name" validate:"required,min=2,max=50"`
			Gender   string      `json:"gender" validate:"omitempty,oneof=Male Female"`
			Phone    string      `json:"phone" validate:"omitempty,min=9,max=11"`
			Role     entity.Role `json:"role" validate:"omitempty,oneof=Admin Customer Seller"`
		}
		validate := validator.New()
		_ = validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
			return len(fl.Field().String()) > 8 && len(fl.Field().String()) < 20
		})

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if err := validate.Struct(input); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		id, err := service.CreateUser(input.Email, input.Password, input.Name, input.Gender, input.Phone, input.Role)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(id); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

// getUser get user by id
func getUser(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error get user by id"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
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
			w.Write([]byte(errorMessage))
		}
	})
}

// deleteUser delete a user
func deleteUser(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error delete user"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode("Delete user successful"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func updateUser(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error update user"
		var input struct {
			UserID   string `json:"id"`
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,min=8,passwd"`
			Name     string `json:"name" validate:"omitempty,min=2,max=50"`
			Gender   string `json:"gender" validate:"omitempty,oneof=Male Female"`
			Phone    string `json:"phone" validate:"required,min=9,max=11"`
		}
		validate := validator.New()
		_ = validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
			return len(fl.Field().String()) > 8 && len(fl.Field().String()) < 20
		})
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if err := validate.Struct(input); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := entity.StringToID(input.UserID)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		u := entity.User{
			UserId:   id,
			Email:    input.Email,
			Password: input.Password,
			Name:     input.Name,
			Gender:   input.Gender,
			Phone:    input.Phone,
		}
		err = service.UpdateUser(&u)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if err := json.NewEncoder(w).Encode("Update user successful"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

// func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service user.Service) {
// 	r.Handle("/user", n.With(
// 		negroni.Wrap(listUsers(service)),
// 	)).Methods("GET", "OPTIONS").Name("listUsers")

// 	r.Handle("/user", n.With(
// 		negroni.Wrap(createUser(service)),
// 	)).Methods("POST", "OPTIONS").Name("createUser")

// 	r.Handle("/user/{id}", n.With(
// 		negroni.Wrap(getUser(service)),
// 	)).Methods("GET", "OPTIONS").Name("getUser")

// 	r.Handle("/user/{id}", n.With(
// 		negroni.Wrap(deleteUser(service)),
// 	)).Methods("DELETE", "OPTIONS").Name("deleteUser")
// }

func MakeUserHandlers(r *mux.Router, service user.Service) {
	r.Handle("/user", listUsers(service)).Methods(http.MethodGet)

	r.Handle("/user", createUser(service)).Methods(http.MethodPost)

	r.Handle("/user/{id}", getUser(service)).Methods(http.MethodGet)

	r.Handle("/user/{id}", deleteUser(service)).Methods(http.MethodDelete)

	r.Handle("/user", updateUser(service)).Methods(http.MethodPut)
}
