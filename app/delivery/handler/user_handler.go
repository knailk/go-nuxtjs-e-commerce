package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/delivery/presenter"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase/user"
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
			fmt.Println("case query 1")
		default:
			data, err = service.SearchUsers(query)
			fmt.Println("case defalt 1")
		}
		fmt.Println("Line 1")
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		fmt.Println("Line 2")
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			fmt.Println("Data == nil")
			return
		}
		var toJ []*presenter.User
		for _, d := range data {
			toJ = append(toJ, &presenter.User{
				ID:     d.UserId,
				Email:  d.Email,
				Name:   d.Name,
				Phone:  d.Phone,
				Gender: d.Gender,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createUser(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding user"
		var input struct {
			Email    string      `json:"email"`
			Password string      `json:"password"`
			Name     string      `json:"name"`
			Gender   string      `json:"gender"`
			Phone    string      `json:"phone"`
			Role     entity.Role `json:"role"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
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
		toJ := &presenter.User{
			ID:     id,
			Email:  input.Email,
			Name:   input.Name,
			Phone:  input.Phone,
			Gender: input.Gender,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getUser(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error get user by id"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			fmt.Printf("err1: %v\n", err)
			return
		}
		data, err := service.GetUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			fmt.Printf("err2: %v\n", err)
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			fmt.Println("data==nil")
			return
		}
		toJ := &presenter.User{
			ID:     data.UserId,
			Email:  data.Email,
			Name:   data.Name,
			Phone:  data.Phone,
			Gender: data.Gender,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deleteUser(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error delete user"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			fmt.Println(err)
			return
		}
		err = service.DeleteUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
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

	r.Handle("/user/{id}",deleteUser(service)).Methods(http.MethodDelete)
}
