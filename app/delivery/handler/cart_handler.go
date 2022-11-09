package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/knailk/go-shopee/app/config"
	"github.com/knailk/go-shopee/app/delivery/middleware"
	"github.com/knailk/go-shopee/app/usecase"
)

var store = sessions.NewCookieStore([]byte(config.SESSION_KEY))

func getCart(service usecase.CartService) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}

func MakeCartHandlers(r *mux.Router, service usecase.UserUsecase) {
	r.Handle("/cart", getCart(service)).Methods(http.MethodGet)

}

