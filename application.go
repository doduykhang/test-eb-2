package main

import (
	"net/http"
)

func main() {
    	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":5000", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
