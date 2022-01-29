package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"norest/src/repository"
	"strconv"
)

func GetAllBooks(rw http.ResponseWriter, r *http.Request) {
	books := repository.GetAllBooks()
	booksJson, err := json.Marshal(books)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Failed to marshal %v as JSON, got error %+v", books, err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, string(booksJson))
}

func CreateBook(rw http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "Failed to parse request body, go error %+v", err)
		return
	}

	var newBook repository.Book
	if err := json.Unmarshal(requestBody, &newBook); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Failed to unmarshal json, got error %+v", err)
		return
	}

	repository.AddBook(&newBook)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(requestBody))
}

func GetBookById(rw http.ResponseWriter, r *http.Request, bookId string) {
	book, err := repository.GetBookById(bookId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(rw, "Failed to find book with id: %s", bookId)
			return
		}

		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Failed to get book, error %+v", err)
		return
	}

	bookJson, err := json.Marshal(book)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Failed to marshal %v as JSON, got error %+v", book, err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, string(bookJson))
}

func DeleteBook(rw http.ResponseWriter, r *http.Request, bookId string) {
	err := repository.DeleteBook(bookId)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Failed to delete book, error %+v", err)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
	fmt.Fprint(rw, "")
}

func UpdateBook(rw http.ResponseWriter, r *http.Request, bookId string) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "Failed to parse request body, go error %+v", err)
		return
	}

	var newBook repository.Book
	err = json.Unmarshal(requestBody, &newBook)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "Failed to unmarshal json, got error %+v", err)
		return
	}

	err = repository.UpdateBook(bookId, &newBook)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(rw, "Failed to find book with id: %s", bookId)
			return
		}

		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Failed to update book, error %+v", err)
		return
	}

	bookIdInt, _ := strconv.Atoi(bookId)
	newBook.Id = bookIdInt
	bookJson, err := json.Marshal(newBook)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Failed to marshal %v as JSON, got error %+v", newBook, err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, string(bookJson))
}
