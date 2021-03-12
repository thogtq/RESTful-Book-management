package models

import (
	"database/sql"
)

//Book type
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

//GetBooks returns array of book model
//return an error `resErr` as json format if any error occurs
func (*Book) GetBooks(db *sql.DB) (books []Book, resErr map[string]interface{}) {
	books = []Book{}
	rows, err := db.Query("SELECT id, title, author, year FROM book")
	if err != nil {
		resErr = *NewErrorResponse(500, "cannot exec query: "+err.Error())
		return nil, resErr
	}
	defer rows.Close()
	for rows.Next() {
		row := Book{}
		if err := rows.Scan(&row.ID, &row.Title, &row.Author, &row.Year); err != nil {
			resErr = *NewErrorResponse(500, "cannot matching books: "+err.Error())
			return nil, resErr
		}
		books = append(books, row)
	}
	return books, nil
}

//GetBook return book by id
//return an error `resErr` as json format if any error occurs
func (*Book) GetBook(db *sql.DB, id int) (book Book, resErr map[string]interface{}) {
	book = Book{}
	row, err := db.Query("SELECT id, title, author, year FROM book WHERE id=?", id)
	if err != nil {
		resErr = *NewErrorResponse(500, "cannot exec query: "+err.Error())
		return book, resErr
	}
	count := 0
	for row.Next() {
		count++
		if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			resErr = *NewErrorResponse(500, "cannot matching books: "+err.Error())
			return book, resErr
		}
	}
	if count == 0 {
		resErr = *NewErrorResponse(404, "book not found!")
		return book, resErr
	}
	return book, nil
}

//AddBook add new book and return resSucc json format
//return an error `resErr` as json format if any error occurs
func (*Book) AddBook(db *sql.DB, book *Book) (resSucc, resErr map[string]interface{}) {
	result, err := db.Exec("INSERT INTO book(title,author,year) VALUES(?,?,?)", book.Title, book.Author, book.Year)
	if err != nil {
		resErr = *NewErrorResponse(500, "cannot exec query: "+err.Error())
		return nil, resErr
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		resErr = *NewErrorResponse(500, "cannot get inserted id: "+err.Error())
		return nil, resErr
	}
	return *NewSuccessResponse(insertedID), nil
}

//UpdateBook update book and return resSucc json format
//Note: If all the New book's field are same as the old one, errors will occurs
//return an error `resErr` as json format if any error occurs
func (*Book) UpdateBook(db *sql.DB, book *Book) (resSucc, resErr map[string]interface{}) {
	result, err := db.Exec(
		"UPDATE book set title=?, author=?, year=? WHERE id =?",
		book.Title, book.Author, book.Year, book.ID,
	)
	if err != nil {
		resErr = *NewErrorResponse(500, "cannot exec query: "+err.Error())
		return nil, resErr
	}
	count, err := result.RowsAffected()
	if err != nil {
		resErr = *NewErrorResponse(500, "update book fail: "+err.Error())
		return nil, resErr
	}
	if count == 0 {
		resErr = *NewErrorResponse(404, "invalid book id or no new updated data : ")
		return nil, resErr
	}
	return *NewSuccessResponse("updated"), nil
}

//RemoveBook delete book with input `id` and return resSucc json format as successful remove
//return an error `resErr` as json format if any error occurs
func (*Book) RemoveBook(db *sql.DB, id int) (resSucc, resErr map[string]interface{}) {
	result, err := db.Exec(
		"DELETE FROM book WHERE id =?",
		id,
	)
	if err != nil {
		resErr = *NewErrorResponse(500, "cannot exec query: "+err.Error())
		return nil, resErr
	}
	count, err := result.RowsAffected()
	if err != nil {
		resErr = *NewErrorResponse(500, "delete fail: "+err.Error())
		return nil, resErr
	}
	if count == 0 {
		resErr = *NewErrorResponse(404, "invalid book id")
		return nil, resErr
	}
	return *NewSuccessResponse("Book removed"), nil
}

//NewErrorResponse return a json format of error response
func NewErrorResponse(code int, message string) *map[string]interface{} {
	errRes := &map[string]interface{}{
		"success": false,
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}
	return errRes
}

//NewSuccessResponse return a json format of success response
func NewSuccessResponse(data interface{}) *map[string]interface{} {
	res := &map[string]interface{}{
		"success": true,
		"data":    data,
	}
	return res
}
