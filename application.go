package main

import (
	"net/http"

	"github.com/tjarratt/babble"
)

func main() {
    	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":5000", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	babbler := babble.NewBabbler()
	w.Write([]byte("hello " + babbler.Babble()))
}
