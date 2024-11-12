package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", helloWord)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func helloWord(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipCode := r.URL.Query().Get( "zipcode")
	if zipCode ==  "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := w.Write([]byte(zipCode))
	if err != nil {
		panic(err)
	}
}

