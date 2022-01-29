package routes

import (
	"net/http"
	"norest/src/controller"
	"strings"
)

func Api() {
	http.HandleFunc("/api/books", func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetAllBooks(rw, r)
		case http.MethodPost:
			controller.CreateBook(rw, r)
		}
	})
	http.HandleFunc("/api/books/", func(rw http.ResponseWriter, r *http.Request) {
		// Any pattern ended with "/" will match all subppaterns, so when usign the net/http as a
		// router we can extract the subpattern as a value
		id := strings.TrimPrefix(r.URL.Path, "/api/books/")

		switch r.Method {
		case http.MethodGet:
			controller.GetBookById(rw, r, id)
		case http.MethodDelete:
			controller.DeleteBook(rw, r, id)
		case http.MethodPut:
			controller.UpdateBook(rw, r, id)
		}
	})
}
