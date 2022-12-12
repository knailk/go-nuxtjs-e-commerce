package main

import (
	cx "context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/knailk/go-nuxtjs-e-commerce/app/config"
	"github.com/knailk/go-nuxtjs-e-commerce/app/delivery/handler"
	"github.com/knailk/go-nuxtjs-e-commerce/app/usecase"
	"github.com/rs/cors"

	"github.com/knailk/go-nuxtjs-e-commerce/repository/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var dsn string
	var setLimits bool
	flag.StringVar(&dsn, "dsn", "/ecommerceDB.db", "SQLite DSN")
	flag.BoolVar(&setLimits, "limits", false, "Sets DB limits")
	flag.Parse()

	db, err := openDB(dsn, setLimits)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	dao := sqlite.NewDAO(db)
	// userRepo := sqlite.NewUserRepo(db)
	// userService := user.NewService(userRepo)

	// productRepo := sqlite.NewProductRepo(db)
	// productService := product.NewService(productRepo)

	// categoryRepo := sqlite.NewCategoryRepo(db)
	// categoryService := category.NewService(categoryRepo)
	userService := usecase.NewUserService(dao)
	productService := usecase.NewProductService(dao)
	categoryService := usecase.NewCategoryService(dao)
	authService := usecase.NewAuthService(dao)
	cartService := usecase.NewCartService(dao)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	})

	if err != nil {
		log.Panic(err.Error())
	}
	r := mux.NewRouter()
	h := c.Handler(r)
	//handler
	handler.MakeUserHandlers(r, userService)
	handler.MakeProductHandlers(r, productService, categoryService)
	handler.MakeAuthHandlers(r, authService)
	handler.MakeCartHandlers(r, cartService)

	http.Handle("/", h)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
	}
}

func openDB(dsn string, setLitmits bool) (*sql.DB, error) {
	//connect to database
	db, err := sql.Open("sqlite3", "ecommerceDB.db")
	if err != nil {
		return nil, err
	}

	if setLitmits {
		fmt.Println("Setting limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx, cancel := cx.WithTimeout(cx.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
