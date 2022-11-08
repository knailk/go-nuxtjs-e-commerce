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
	"github.com/knailk/go-shopee/app/delivery/handler"
	"github.com/knailk/go-shopee/app/usecase"

	"github.com/knailk/go-shopee/repository/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func main() {
	var dsn string
	var setLimits bool
	flag.StringVar(&dsn, "dsn", "/shopeeDB.db", "SQLite DSN")
	flag.BoolVar(&setLimits, "limits", false, "Sets DB limits")
	flag.Parse()

	db, err := openDB(dsn, setLimits)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sql.DB){
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

	
	if err != nil {
		log.Panic(err.Error())
	}
	r := mux.NewRouter()
	
	//handler
	handler.MakeUserHandlers(r, userService)
	handler.MakeProductHandlers(r, productService, categoryService)
	handler.MakeAuthHandlers(r,authService)
	
	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(8081),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err.Error())
	}
}

func openDB(dsn string, setLitmits bool) (*sql.DB, error){
	//connect to database
	db,err := sql.Open("sqlite3", "shopeeDB.db")
	if err != nil {
		return nil, err
	}

	if setLitmits {
		fmt.Println("Setting limits")
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(5)
	}

	ctx,cancel := cx.WithTimeout(cx.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
