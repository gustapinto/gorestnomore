package main

import (
	"net/http"
	"norest/src/controller"
)

func main() {
	http.HandleFunc("/api/books", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetAllBooks(rw, r)
		case http.MethodPost:
			controller.CreateBook(rw, r)
		}
	})

	http.ListenAndServe(":80", nil)
}
