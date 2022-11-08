package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/config"
	"github.com/knailk/go-shopee/app/delivery/middleware"
	"github.com/knailk/go-shopee/app/delivery/presenter"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase/user"
	"golang.org/x/crypto/bcrypt"

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

func signIn(service user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error login"
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
		if err := validate.Struct(input); err != nil {
			logError(err, errorMessage, w)
			return
		}
		authUser, err := service.AuthUser(input.UserName)
		if err != nil {
			logError(err, errorMessage, w)
			return
		}
		if err != nil {
			logError(err, errorMessage, w)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(authUser.Password), []byte(input.Password)); err != nil {
			logError(err, "Incorrect password", w)
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

func isAuthorized(handler http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get token
		if r.Header["Token"] == nil {
			logError(nil, "No token found", w)
			return
		}
		var mySigningKey = []byte(config.SECRET_KEY)
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logError(nil, "There was an error in parsing", w)
			}
			return mySigningKey, nil
		})
		if err != nil {
			logError(err, "Your Token has been expired", w)
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

func logError(err error, errorMessage string, w http.ResponseWriter) {
	log.Println(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errorMessage))
}


func MakeUserHandlers(r *mux.Router, service user.Service) {
	r.Handle("/admin/user", listUsers(service)).Methods(http.MethodGet)

	r.Handle("/admin/user", createUser(service)).Methods(http.MethodPost)

	r.Handle("/admin/user/{id}", getUser(service)).Methods(http.MethodGet)

	r.Handle("/admin/user/{id}", deleteUser(service)).Methods(http.MethodDelete)

	r.Handle("/admin/user", updateUser(service)).Methods(http.MethodPut)

	r.Handle("/admin", isAuthorized(middleware.AdminIndex)).Methods(http.MethodGet)

	r.Handle("/login", signIn(service)).Methods(http.MethodPost)
}
