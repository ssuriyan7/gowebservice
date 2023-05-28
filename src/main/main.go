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

	router.HandleFunc("/anime/add", endpoint.AddAnime).Methods("POST")
	router.HandleFunc("/anime/delete", endpoint.DeleteAnime).Methods("DELETE")

	fmt.Println("Starting server at 8080.")
	log.Print(http.ListenAndServe(":8080", router))

}
