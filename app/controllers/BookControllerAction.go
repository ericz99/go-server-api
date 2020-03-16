package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"go-server-api/app/models"
	"go-server-api/app/service"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// # Book Struct Array
var books []models.Book

// GetBookTestRoute Controller Route
func GetBookTestRoute(w http.ResponseWriter, r *http.Request) {
	// # set header as content type json format
	w.Header().Set("Content-Type", "application/json")

	var jsonData []byte
	t := models.Test{Msg: "Hello world"}
	jsonData, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(jsonData))
}

// SaveBookRoute Controller Route
func SaveBookRoute(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var bookService service.BookService
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	// # save book
	err = bookService.SaveBook(db, &book)
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: "BAD", Message: "Failed added book to db!"})
		return
	}

	// # set header as content type json format
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer r.Body.Close()

	json.NewEncoder(w).Encode(book)
}

// GetAllBookRoute Route
func GetAllBookRoute(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var bookService service.BookService
	books := []models.Book{}
	// # set header as content type json format
	w.Header().Set("Content-Type", "application/json")
	err := bookService.FindAll(db, &books)
	if err != nil {
		fmt.Println(err)
		return
	}

	// # return json body message
	json.NewEncoder(w).Encode(books)
}

// GetBookRoute Route
func GetBookRoute(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var bookService service.BookService
	// # set header as content type json format
	w.Header().Set("Content-Type", "application/json")
	// # get the request parameters; route variable
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var book models.Book
	// # get book by id
	err := bookService.GetBookByID(db, id, &book)

	if err != nil {
		json.NewEncoder(w).Encode(models.Response{Status: "BAD", Message: "ID does not exist!"})
		return
	}

	// # return json body message
	json.NewEncoder(w).Encode(book)
}

// DeleteBookRoute Route
func DeleteBookRoute(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var bookService service.BookService
	// # set header as content type json format
	w.Header().Set("Content-Type", "application/json")
	// # get the request parameters; route variable
	vars := mux.Vars(r)
	idParams := vars["id"]

	var book models.Book
	id, _ := strconv.Atoi(idParams)
	err := bookService.GetBookByID(db, id, &book)

	if err != nil {
		json.NewEncoder(w).Encode(models.Response{Status: "BAD", Message: "ID does not exist!"})
		return
	}

	// # delete book
	_ = bookService.DeleteBook(db, &book)
	// # return json body message
	json.NewEncoder(w).Encode(&models.Response{Status: "GOOD", Message: "Successfully deleted book!"})
}
