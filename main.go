package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"url-shortener/cache"
	"url-shortener/handler"
	"url-shortener/service"
	"url-shortener/store"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	connStr := "postgres://user1:password1@localhost:5432/user1?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Database Connection failed....")
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Database Connection failed at db.Ping:", err)
		return
	}
	redisCache := cache.NewRedisCache()
	str := store.NewStore(db, redisCache)
	srv := service.NewService(str)
	h := handler.NewHandler(srv)

	router.HandleFunc("/shorten", h.ShortenUrl).Methods("POST")
	router.HandleFunc("/redirect", h.Redirect).Methods("GET")

	err = http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println("Server Connetion Failed..")
	}

}
