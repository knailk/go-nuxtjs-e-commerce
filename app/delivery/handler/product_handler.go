package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase/category"
	"github.com/knailk/go-shopee/app/usecase/product"
)

// listCategories return http handler
func listCategories(categoryService category.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading categories"
		data, err := categoryService.ListCategories()
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
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func getCategory(productService product.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//transmit category service to get cate id
	})

}

func MakeProductHandlers(r *mux.Router, productService product.Service, categoryService category.Service) {
	r.Handle("", listCategories(categoryService)).Methods(http.MethodGet)
	
	r.Handle("/{cate_id}", getCategory(productService)).Methods(http.MethodGet)

	

}
