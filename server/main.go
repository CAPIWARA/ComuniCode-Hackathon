package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hola", auth)
	log.Fatal(http.ListenAndServe(":8080", router), nil)
}
