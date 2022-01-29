package main

import (
	"fmt"
	"net/http"
	"norest/src/routes"
)

func main() {
	routes.Api()

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Failed to start Http server, go error %+v", err)
	}
}
