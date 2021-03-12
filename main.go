package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"github.com/thogtq/restful-book-management/m/v1/controllers"
	"github.com/thogtq/restful-book-management/m/v1/driver"
)

var db *sql.DB

func init() {
	//Load all env varibles from .env file at root folder
	gotenv.Load()
}
func main() {

	db = driver.ConnectDB()
	defer db.Close()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
}
