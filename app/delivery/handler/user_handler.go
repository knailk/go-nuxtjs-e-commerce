package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/config"
	"github.com/knailk/go-shopee/app/delivery/middleware"
	"github.com/knailk/go-shopee/app/delivery/presenter"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase"
)

// listUsers return http handler
func listUsers(service usecase.UserUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error reading users"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// getUser get user by id
func getUser(service usecase.UserUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error delete user"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		err = service.DeleteUser(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err := json.NewEncoder(w).Encode("delete user successful"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func updateUser(service usecase.UserUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error update user"
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
			logInternalServerError(err,err.Error(),w)
			return
		}
		if err := validate.Struct(input); err != nil {
			logInternalServerError(err,err.Error(),w)
			return
		}
		id, err := entity.StringToID(input.UserID)
		if err != nil {
			logInternalServerError(err,err.Error(),w)
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
			w.Write([]byte(err.Error()))
			return
		}
		if err := json.NewEncoder(w).Encode("update user successful"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func isAuthorized(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get token
		if r.Header["Token"] == nil {
			logInternalServerError(nil, "no token found", w)
			return
		}
		var mySigningKey = []byte(config.SECRET_KEY)
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logInternalServerError(nil, "there was an error in parsing", w)
			}
			return mySigningKey, nil
		})
		if err != nil {
			logInternalServerError(err, "your Token has been expired", w)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "Admin" {
				r.Header.Set("Role", "Admin")
				return
			} else if claims["role"] == "Customer" {
				r.Header.Set("Role", "Customer")
				return
			}
			// } else if claims["role"] == "Seller" {
			// 	r.Header.Set("Role", "Seller")
			// 	return
			// }
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

	r.Handle("/admin/user", createUser(service)).Methods(http.MethodPost)

	r.Handle("/admin/user/{id}", getUser(service)).Methods(http.MethodGet)

	r.Handle("/admin/user/{id}", deleteUser(service)).Methods(http.MethodDelete)

	r.Handle("/admin/user", updateUser(service)).Methods(http.MethodPut)

	r.Handle("/admin", isAuthorized(middleware.AdminIndex)).Methods(http.MethodGet)
}
