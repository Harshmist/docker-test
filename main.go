package main

import (
	"fmt"
	"net/http"
)

func display(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello there!")
}

func main() {
	http.HandleFunc("/home", display)
	http.ListenAndServe(":8080", nil)
}
