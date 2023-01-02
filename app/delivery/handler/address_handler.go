package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/knailk/go-nuxtjs-e-commerce/app/delivery/middleware"
	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
	"github.com/knailk/go-nuxtjs-e-commerce/app/usecase"
)

func getAddress(service usecase.AddressUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		clams, _ := middleware.ExtractClaims(w, r)
		email := fmt.Sprintf("%v", clams["email"])
		address, err := service.GetAddress(email)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err == entity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err := json.NewEncoder(w).Encode(address); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}
func addAddress(service usecase.AddressUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		clams, _ := middleware.ExtractClaims(w, r)
		email := fmt.Sprintf("%v", clams["email"])
		var input struct {
			Country  string `json:"country" validate:"required"`
			City     string `json:"city" validate:"required"`
			District string `json:"district" validate:"required"`
			Ward     string `json:"ward" validate:"required"`
			Address  string `json:"address" validate:"required"`
		}
		//validate
		validate := validator.New()
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		if err := validate.Struct(input); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte("Input invalid"))
			return
		}
		address := entity.NewAddress(email, input.Country, input.City, input.District, input.Ward, input.Address)
		err = service.AddAddress(address)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

func MakeAddressHandlers(r *mux.Router, service usecase.AddressUsecase) {
	r.Handle("/address", getAddress(service)).Methods(http.MethodGet)
	r.Handle("/address", addAddress(service)).Methods(http.MethodPost)
}
