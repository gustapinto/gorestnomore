package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"norest/src/repository"
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

	rw.WriteHeader(http.StatusCreated)
	fmt.Fprint(rw, string(requestBody))
}
