package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "strconv"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books []Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for _, book := range books {
        if book.ID == id {
            json.NewEncoder(w).Encode(book)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = len(books) + 1
    books = append(books, book)
    json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var updatedBook Book
    _ = json.NewDecoder(r.Body).Decode(&updatedBook)

    for i, book := range books {
        if book.ID == id {
            updatedBook.ID = id
            books[i] = updatedBook
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}
