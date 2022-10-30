package handler

import (
	"github.com/knailk/go-shopee/app/usecase/product"
	"github.com/labstack/echo"
)

// represent a response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// represent http handler for product
type ProductHandler struct {
	PUsecase product.Usecase
}

// initialize the product/ resources endpoint
func NewProductHandler(e *echo.Echo, pu product.Usecase) {
	handler := &ProductHandler{
		PUsecase: pu,
	}
	e.GET("/products:category", handler.List)
	e.GET("/products:query", handler.Search)
	e.POST("/products", handler.Create)
	e.GET("/products/:id", handler.GetByID)
	e.PUT("/products/:id", handler.Update)
	e.DELETE("/products/:id", handler.Delete)
}

// TODO: implement function product handler
func (p *ProductHandler) List(c echo.Context) error {
	return nil
}

func (p *ProductHandler) Search(c echo.Context) error {
	return nil
}

func (p *ProductHandler) Create(c echo.Context) error {
	return nil
}

func (p *ProductHandler) GetByID(c echo.Context) error {
	return nil
}

func (p *ProductHandler) Update(c echo.Context) error {
	return nil
}

func (p *ProductHandler) Delete(c echo.Context) error {
	return nil
}
