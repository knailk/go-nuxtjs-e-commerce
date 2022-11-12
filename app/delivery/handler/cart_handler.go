package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/delivery/presenter"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase"
)

func getCart(service usecase.CartUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := entity.StringToID(r.URL.Query().Get("user_id"))
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		cart, totalPrice, err := service.GetCart(userId)
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		if cart == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("no product in cart"))
			return
		}
		toJ := &presenter.Cart{
			ListProducts: cart,
			TotalPrice:   totalPrice,
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func addToCart(service usecase.CartUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := entity.StringToID(r.URL.Query().Get("user_id"))
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		var input struct {
			ProductId string `json:"productId"`
			Quantity  int64  `json:"quantity"`
		}
		err = json.NewDecoder(r.Body).Decode(&input)
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
		cart := entity.Cart{
			UserId:    userId,
			ProductId: productId,
			Quantity:  input.Quantity,
		}
		err = service.AddToCart(&cart)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode("add to cart successful"); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})
}

func removeItem(service usecase.CartUsecase) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := entity.StringToID(r.URL.Query().Get("user_id"))
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		var inputId string 
		err = json.NewDecoder(r.Body).Decode(&inputId)
		if err != nil {
			logInternalServerError(err,err.Error(),w)
			return
		}
		productId, err := entity.StringToID(inputId)
		if err != nil {
			logInternalServerError(err,err.Error(),w)
			return
		}
		service.RemoveProduct(userId,productId)

	})
}

func MakeCartHandlers(r *mux.Router, service usecase.CartUsecase) {
	r.Handle("/cart", getCart(service)).Methods(http.MethodGet)

	//example ulr, we will call function add item when add at home and item detail
	// url: "/cart?user_id=123712361"
	r.Handle("/cart", addToCart(service)).Methods(http.MethodPost)

	r.Handle("/cart", removeItem(service)).Methods(http.MethodDelete)

}
