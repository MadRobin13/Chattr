package main

import (
	"fmt"
	"net/http"
)

func setUpRoutes() {
	http.handleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}

func main() {
	setUpRoutes()

}