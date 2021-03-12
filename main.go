package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

var db *sql.DB

func init() {
	//Load all env varibles from .env file at root folder
	gotenv.Load()
}
func main() {
	var ( //database connection varibles
		dbConnection   = os.Getenv("DB_CONNECTION")
		mysqlUsername  = os.Getenv("DB_USERNAME")
		mysqlPassword  = os.Getenv("DB_PASSWORD")
		mysqlHost      = os.Getenv("DB_HOST")
		mysqlPort      = os.Getenv("DB_PORT")
		mysqlDb        = os.Getenv("DB_DATABASE")
		dataSourceName = mysqlUsername + ":" + mysqlPassword + "@tcp(" + mysqlHost + mysqlPort + ")/" + mysqlDb
	)
	var err error
	db, err = sql.Open(dbConnection, dataSourceName)
	if err != nil {
		//An exception must not happens
		log.Printf("cannot connect to mysql database\n")
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", router))
}
func getBooks(w http.ResponseWriter, r *http.Request) {
	books := []book{}
	rows, err := db.Query("SELECT id, title, author, year FROM book")
	if err != nil {
		errRes := *newErrorResponse(500, "cannot exec query")
		json.NewEncoder(w).Encode(errRes)
		return
	}
	defer rows.Close()
	for rows.Next() {
		row := book{}
		if err := rows.Scan(&row.ID, &row.Title, &row.Author, &row.Year); err != nil {
			errRes := *newErrorResponse(500, "cannot matching books")
			json.NewEncoder(w).Encode(errRes)
			return
		}
		books = append(books, row)
	}
	json.NewEncoder(w).Encode(books)
	return
}
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]
	bookData := book{}
	row, err := db.Query("SELECT id, title, author, year FROM book WHERE id=?", bookID)
	if err != nil {
		errRes := *newErrorResponse(500, "cannot exec query")
		json.NewEncoder(w).Encode(errRes)
		return
	}
	count := 0
	for row.Next() {
		count++
		if err := row.Scan(&bookData.ID, &bookData.Title, &bookData.Author, &bookData.Year); err != nil {
			errRes := *newErrorResponse(500, "cannot matching books: "+err.Error())
			json.NewEncoder(w).Encode(errRes)
			return
		}
	}
	if count == 0 {
		errRes := *newErrorResponse(404, "book not found!")
		json.NewEncoder(w).Encode(errRes)
		return
	}
	json.NewEncoder(w).Encode(bookData)
	return
}

//"INSERT INTO book(title,author,year) VALUES(?,?,?)", params["title"], params["author"], params["year"]
func addBook(w http.ResponseWriter, r *http.Request) {
	bookData := book{}
	json.NewDecoder(r.Body).Decode(&bookData)

	result, err := db.Exec("INSERT INTO book(title,author,year) VALUES(?,?,?)", bookData.Title, bookData.Author, bookData.Year)
	if err != nil {
		errRes := *newErrorResponse(500, "cannot insert query: "+err.Error())
		json.NewEncoder(w).Encode(errRes)
		return
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		errRes := *newErrorResponse(500, "cannot get inserted id: "+err.Error())
		json.NewEncoder(w).Encode(errRes)
		return
	}
	json.NewEncoder(w).Encode(newSuccessResponse(insertedID))
	return
}
func updateBook(w http.ResponseWriter, r *http.Request) {

}
func removeBook(w http.ResponseWriter, r *http.Request) {

}
func newErrorResponse(code int, message string) *map[string]interface{} {
	errRes := &map[string]interface{}{
		"success": false,
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}
	return errRes
}
func newSuccessResponse(data interface{}) *map[string]interface{} {
	res := &map[string]interface{}{
		"success": true,
		"data":    data,
	}
	return res
}
