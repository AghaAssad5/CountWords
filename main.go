package main

import (
	"net/http"

	"github.com/aghaasad/countWords/wordcount"
	"github.com/gorilla/mux"
)

// Main function
func main() {

	router := mux.NewRouter()

	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/countWords", wordcount.CountWords).Methods("POST")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}
