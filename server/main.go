package main

import (
	"capiwara-boilerplate/gql"
	"comunicode-hachathon/server/db"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	router := mux.NewRouter()

	if err := db.NewSession("", ""); err != nil {
		return
	}
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &gql.Schema,
	})
	router.HandleFunc("/login", loginAuth)
	router.Handle("/graphql", requireAuth(handler))
	log.Println("Server started at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
