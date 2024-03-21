package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func mainServe() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
