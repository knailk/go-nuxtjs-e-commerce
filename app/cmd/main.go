package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"
)

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

	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
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