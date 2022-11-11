package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/delivery/presenter"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase"
)

// listCategories return http handler
func listCategories(categoryService usecase.CategoryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error reading categories"
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

// getProducts get list product by category id
func getProducts(productService usecase.ProductUsecase, categoryService usecase.CategoryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error get products by category id"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["cate_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		category, err := categoryService.GetCategory(int64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := productService.ListProducts(int64(id))

		w.Header().Set("Content-type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		var toJson []*presenter.Product
		for _, d := range data {
			toJson = append(toJson, &presenter.Product{
				ProductID:      d.ProductID,
				Name:           d.Name,
				Price:          d.Price,
				Description:    d.Description,
				QuantitySold:   d.QuantitySold,
				AvailableUnits: d.AvailableUnits,
				CreatedAt:      d.CreatedAt,
				UpdatedAt:      d.UpdatedAt,
				Category:       category.CategoryName,
			})
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		//data, err:= categoryService.GetCategory(int64(id))
	})
}

func getProduct(productService usecase.ProductUsecase, categoryService usecase.CategoryUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error get product by id"
		vars := mux.Vars(r)
		//get category data
		categoryId, err := strconv.Atoi(vars["cate_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		category, err := categoryService.GetCategory(int64(categoryId))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		//get product data
		productId, err := entity.StringToID(vars["product_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		p, err := productService.GetProduct(productId)

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		toJson := &presenter.Product{
			ProductID:      p.ProductID,
			Name:           p.Name,
			Price:          p.Price,
			Description:    p.Description,
			QuantitySold:   p.QuantitySold,
			AvailableUnits: p.AvailableUnits,
			CreatedAt:      p.CreatedAt,
			UpdatedAt:      p.UpdatedAt,
			Category:       category.CategoryName,
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

// createProduct create new product by admin
func createProduct(productService usecase.ProductUsecase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error adding product"
		var input struct {
			Name           string `json:"name" validate:"required,min=2,max=50"`
			Price          int64  `json:"price" validate:"omitempty"`
			Description    string `json:"description" validate:"omitempty"`
			QuantitySold   int64  `json:"quantitySold"`
			AvailableUnits int64  `json:"availableUnits"`
			Category       int64  `json:"categoryId"`
		}
		validate := validator.New()
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
		p := entity.NewProduct(input.Name, input.Price, input.Description, input.QuantitySold, input.AvailableUnits, input.Category)
		id, err := productService.CreateProduct(p)
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

func MakeProductHandlers(r *mux.Router, productService usecase.ProductUsecase, categoryService usecase.CategoryUsecase) {
	r.Handle("/product", listCategories(categoryService)).Methods(http.MethodGet)

	r.Handle("/product/{cate_id}", getProducts(productService, categoryService)).Methods(http.MethodGet)

	r.Handle("/product/{cate_id}/{product_id}", getProduct(productService, categoryService)).Methods(http.MethodGet)

	r.Handle("/product", createProduct(productService)).Methods(http.MethodPost)

}
