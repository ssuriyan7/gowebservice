package main

import (
	"firstWebService/src/main/endpoint"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/anime/", endpoint.AddAnime).Methods("POST")

	fmt.Println("Starting server at 8080.")
	log.Fatal(http.ListenAndServe(":8080", router))

}
