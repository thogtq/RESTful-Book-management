package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thogtq/restful-book-management/m/v1/models"
)

//Controller method for book
type Controller struct{}

//bookModel is global varible for functions to access methods of book models
var bookModel models.Book

//GetBooks get all book from db
func (*Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookArr, resErr := bookModel.GetBooks(db)
		if resErr != nil {
			json.NewEncoder(w).Encode(resErr)
			return
		}
		json.NewEncoder(w).Encode(bookArr)
	}
}

//GetBook get book with `id` from GET params
func (*Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bookID, _ := strconv.Atoi(params["id"])
		bookData, resErr := bookModel.GetBook(db, bookID)
		if resErr != nil {
			json.NewEncoder(w).Encode(resErr)
			return
		}
		json.NewEncoder(w).Encode(bookData)
		return
	}
}

//AddBook add new book with POST method
func (*Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookData := &models.Book{}
		if err := json.NewDecoder(r.Body).Decode(&bookData); err != nil {
			errRes := *models.NewErrorResponse(500, "cannot matching data: "+err.Error())
			json.NewEncoder(w).Encode(errRes)
			return
		}
		resSucc, resErr := bookModel.AddBook(db, bookData)
		if resErr != nil {
			json.NewEncoder(w).Encode(resErr)
			return
		}
		json.NewEncoder(w).Encode(resSucc)
	}
}

//UpdateBook update date book with PUT method
func (*Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bookData := &models.Book{}
		if err := json.NewDecoder(r.Body).Decode(&bookData); err != nil {
			errRes := *models.NewErrorResponse(500, "cannot matching data: "+err.Error())
			json.NewEncoder(w).Encode(errRes)
			return
		}
		resSucc, resErr := bookModel.UpdateBook(db, bookData)
		if resErr != nil {
			json.NewEncoder(w).Encode(resErr)
			return
		}
		json.NewEncoder(w).Encode(resSucc)
	}
}

//RemoveBook remove book with GET `id request
//No authorization needed
func (*Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bookID, _ := strconv.Atoi(params["id"])
		resSucc, resErr := bookModel.RemoveBook(db, bookID)
		if resErr != nil {
			json.NewEncoder(w).Encode(resErr)
			return
		}
		json.NewEncoder(w).Encode(resSucc)
	}
}
