package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello ZAWARUDO\n")
	})

	http.ListenAndServe(":80", nil)
}
