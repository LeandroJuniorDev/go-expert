package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", findZipCode)
	http.ListenAndServe(":8080", nil)
}

func findZipCode(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte( "Hello, Word!"))
}

