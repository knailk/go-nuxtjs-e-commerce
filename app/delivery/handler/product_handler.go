package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/delivery/middleware"
	"github.com/knailk/go-shopee/app/delivery/presenter"
	"github.com/knailk/go-shopee/app/entity"
	"github.com/knailk/go-shopee/app/usecase"
)

// listCategories return http handler
func listCategories(productService usecase.ProductUsecase, categoryService usecase.CategoryUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		data, err := categoryService.ListCategories()
		if err != nil {
			logInternalServerError(err, err.Error(), w)
			return
		}
		var toJson []*presenter.Category
		for _, v := range data {
			d, err := productService.ListProducts(int64(v.CategoryId))
			if err != nil {
				logInternalServerError(err, err.Error(), w)
				return
			}
			toJson = append(toJson, &presenter.Category{
				CategoryId:    entity.ID(v.CategoryId),
				Name:          v.CategoryName,
				NumberProduct: len(d),
			})
		}
		middleware.Cors(w, r)
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

func topProducts(productService usecase.ProductUsecase) http.Handler{
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		data,err := productService.TopProduct()
		if err != nil && err != entity.ErrNotFound {
			logInternalServerError(err,err.Error(),w)
			return
		}
		middleware.Cors(w, r)
		var toJson []*presenter.Product
		for _, d := range data {
			toJson = append(toJson, &presenter.Product{
				ProductId:      d.ProductID,
				Name:           d.Name,
				Price:          d.Price,
				Description:    d.Description,
				QuantitySold:   d.QuantitySold,
				AvailableUnits: d.AvailableUnits,
				Image:          d.Image,
				CreatedAt:      d.CreatedAt,
				UpdatedAt:      d.UpdatedAt,
				Category: int(d.CategoryID),
			})
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

// getProducts get list product by category id
func getProducts(productService usecase.ProductUsecase, categoryService usecase.CategoryUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error get products by category id"
		middleware.Cors(w, r)
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["cate_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		category, err := categoryService.GetCategory(int64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		data, err := productService.ListProducts(int64(id))
		fmt.Println(data)
		w.Header().Set("Content-type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			logInternalServerError(err,err.Error(),w)
			return
		}
		var toJson []*presenter.Product
		for _, d := range data {
			toJson = append(toJson, &presenter.Product{
				ProductId:      d.ProductID,
				Name:           d.Name,
				Price:          d.Price,
				Description:    d.Description,
				QuantitySold:   d.QuantitySold,
				AvailableUnits: d.AvailableUnits,
				Image:          d.Image,
				CreatedAt:      d.CreatedAt,
				UpdatedAt:      d.UpdatedAt,
				Category: 		category.CategoryId,
			})
		}
		var outPut struct {
			Products []*presenter.Product `json:"products"`
			Category string `json:"category"`
		}
		outPut.Products= toJson
		outPut.Category = category.CategoryName
		if err := json.NewEncoder(w).Encode(outPut); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		//data, err:= categoryService.GetCategory(int64(id))
	})
}

func getProduct(productService usecase.ProductUsecase, categoryService usecase.CategoryUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error get product by id"
		middleware.Cors(w, r)
		vars := mux.Vars(r)
		//get category data
		categoryId, err := strconv.Atoi(vars["cate_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		category, err := categoryService.GetCategory(int64(categoryId))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		//get product data
		productId, err := entity.StringToID(vars["product_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		p, err := productService.GetProduct(productId)

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("data not found"))
			return
		}

		toJson := &presenter.Product{
			ProductId:      p.ProductID,
			Name:           p.Name,
			Price:          p.Price,
			Description:    p.Description,
			QuantitySold:   p.QuantitySold,
			AvailableUnits: p.AvailableUnits,
			Image: p.Image,
			CreatedAt:      p.CreatedAt,
			UpdatedAt:      p.UpdatedAt,
			Category:       category.CategoryId,
		}
		if err := json.NewEncoder(w).Encode(toJson); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
}

// createProduct create new product by admin
func createProduct(productService usecase.ProductUsecase) http.Handler {
	return middleware.ValidateJWT(func(w http.ResponseWriter, r *http.Request) {
		//errorMessage := "error adding product"
		middleware.Cors(w, r)
		var input struct {
			Name           string `json:"name" validate:"required,min=2,max=50"`
			Price          int64  `json:"price" validate:"omitempty"`
			Description    string `json:"description" validate:"omitempty"`
			QuantitySold   int64  `json:"quantitySold"`
			AvailableUnits int64  `json:"availableUnits"`
			Image          string `json:"image"`
			Category       int64  `json:"categoryId"`
		}
		validate := validator.New()
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
		p := entity.NewProduct(input.Name, input.Price, input.Description, input.QuantitySold, input.AvailableUnits, input.Image, input.Category)
		id, err := productService.CreateProduct(p)
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


func MakeProductHandlers(r *mux.Router, productService usecase.ProductUsecase, categoryService usecase.CategoryUsecase) {

	r.Handle("/product", listCategories(productService, categoryService)).Methods(http.MethodGet)

	r.Handle("/product/top", topProducts(productService)).Methods(http.MethodGet)

	r.Handle("/product/{cate_id}", getProducts(productService, categoryService)).Methods(http.MethodGet)

	r.Handle("/product/{cate_id}/{product_id}", getProduct(productService, categoryService)).Methods(http.MethodGet)

	r.Handle("/product", createProduct(productService)).Methods(http.MethodPost)
}
