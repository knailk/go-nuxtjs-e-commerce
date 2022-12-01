package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/knailk/go-nuxtjs-e-commerce/app/delivery/middleware"
	"github.com/knailk/go-nuxtjs-e-commerce/app/entity"
	"github.com/knailk/go-nuxtjs-e-commerce/app/usecase"
)

func getCart(service usecase.CartUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		clams, _ := middleware.ExtractClaims(w, r)
		strValue := fmt.Sprintf("%v", clams["email"])
		cart, totalPrice, err := service.GetCart(strValue)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		type Cart struct {
			ListProductsInCart []*entity.ProductCart `json:"listProductsInCart"`
			TotalPrice         int64                 `json:"totalPrice"`
		}
		toJ := &Cart{
			ListProductsInCart: cart,
			TotalPrice:         totalPrice,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func addToCart(service usecase.CartUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		clams, _ := middleware.ExtractClaims(w, r)
		strValue := fmt.Sprintf("%v", clams["email"])
		var input struct {
			ProductId string `json:"productId"`
			Quantity  int64  `json:"quantity"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		productId, err := entity.StringToID(input.ProductId)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		err = service.AddToCart(productId, strValue, input.Quantity)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

func removeItem(service usecase.CartUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		clams, _ := middleware.ExtractClaims(w, r)
		strValue := fmt.Sprintf("%v", clams["email"])
		var input struct {
			ProductId string `json:"productId"`
			Quantity  int64  `json:"quantity"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		productId, err := entity.StringToID(input.ProductId)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		err = service.RemoveProduct(productId, strValue, input.Quantity)
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

func MakeCartHandlers(r *mux.Router, service usecase.CartUsecase) {
	r.Handle("/cart", getCart(service)).Methods(http.MethodGet)

	r.Handle("/cart/add", addToCart(service)).Methods(http.MethodPost)

	r.Handle("/cart/remove", removeItem(service)).Methods(http.MethodPost)
}
