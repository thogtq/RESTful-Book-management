# RESTful-Book-management
The book management API created in RESTful API architecture with the help of Golang [Gorilla/Mux](https://github.com/gorilla/mux) package
### Packages
* [Gorilla/Mux](https://github.com/gorilla/mux) 
* [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
* [Gotenv](https://github.com/subosito/gotenv)
### Project structure
#### Models
* Performing SQL query to database
* Receive query data from database
* Process data and return to controller
* Handle and return errors to controller
#### Controllers
* Prepare parameters for model functions
* Call functions from models and receive data
* Process data
* Response JSON data to client
#### Driver
* Open connection to database and return `db`*sql.DB  variable
### APIs

#### Get Books
Return all books
* __URL__: /books
* __Method__: GET
* __URL params:__<br>
* __Data params:__<br>
* __Success Response:__
```
{
   [
    {
        "id": 3,
        "title": "Get Programming with Go",
        "author": "Nathan Youngman",
        "year": 2017
    },
    {
        "id": 4,
        "title": "Go Web Programming",
        "author": "Sau Sheong Chang",
        "year": 2013
    }
]
}
```
* __Error Response:__
#### Get book
Return book with book id
* __URL__: /books
* __Method__: GET
* __URL params:__<br>
__Required :__<br>{id}=[integer]
* __Data params:__<br>
* __Success Response:__
```
{
        "id": 4,
        "title": "Go Web Programming",
        "author": "Sau Sheong Chang",
        "year": 2013
}
```
* __Error Response:__
```
{
    "error": {
        "code": 404,
        "message": "book not found!"
    },
    "success": false
}
```
#### Add Book
Add new book
* __URL__: /books
* __Method__: POST
* __URL params:__<br>
* __Data params:__<br>
__Required :__<br>{title}=[string]<br>{author}=[string]<br>{year}=[integer]
* __Success Response:__```{"data":16,"success":true}```
* __Error Response:__<br>
```
{
    "error": {
        "code": 500,
        "message": "cannot matching data: json: cannot unmarshal number into Go struct field Book.title of type string"
    },
    "success": false
}
```
#### Update book
Update existing book
* __URL__: /books
* __Method__: PUT
* __URL params:__<br>
* __Data params:__<br>
__Required :__<br>{id}=[integer]<br>{title}=[string]<br>{author}=[string]<br>{year}=[integer]
* __Success Response:__```{"data":"updated","success":true}```
* __Error Response:__
```
{
    "error": {
        "code": 404,
        "message": "invalid book id or no new updated data : "
    },
    "success": false
}
```
#### Remove book
Remove book with book id
* __URL__: /books/
* __Method__: DELETE
* __URL params:__<br>
__Required :__<br>{id}=[integer]
* __Data params:__<br>
* __Success Response:__```{"data":"Book removed","success":true}```
* __Error Response:__```{
    "error": {
        "code": 404,
        "message": "invalid book id"
    },
    "success": false
}```
### Database
* Database system: MySQL
* Schema :<br>
<img src="https://i.imgur.com/YlhC3AO.png" width="240"><br>
### Docker
Docker-compose.yml file with built in MySQL service and pre-loaded book database with sample data
