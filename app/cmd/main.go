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

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/knailk/go-shopee/app/config"
	"github.com/knailk/go-shopee/app/delivery/handler"
	"github.com/knailk/go-shopee/app/delivery/middleware"
	"github.com/knailk/go-shopee/app/usecase/user"
	"github.com/knailk/go-shopee/repository/sqlite"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type app struct {
	UserService user.Service
}

func main() {
	var dsn string
	var port string
	var setLimits bool
	flag.StringVar(&dsn, "dsn", "/shopeeDB.db", "SQLite DSN")
	flag.StringVar(&port, "port", "8080", "Service Port")
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
	
	userRepo := sqlite.NewUserRepo(db)
	//application := app{UserService: *user.NewService(userRepo)}
	userService := user.NewService(userRepo)

	r := mux.NewRouter()
	
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.NewLogger(),
	)
	//handler user
	handler.MakeUserHandlers(r,*n,*userService)
	
	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

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

func openDB(dsn string, setLitmits bool) (*sql.DB, error){
	//connect to database
	db,err := sql.Open("sqlite3",dsn)
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

func checkErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}