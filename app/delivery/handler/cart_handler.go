package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/usecase"
)



func getCart(service usecase.CartUsecase) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func MakeCartHandlers(r *mux.Router, service usecase.CartUsecase) {
	r.Handle("/cart", getCart(service)).Methods(http.MethodGet)

}

