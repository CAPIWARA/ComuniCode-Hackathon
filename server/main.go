package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

	"github.com/VitorLuizC/ComuniCode-Hackathon/server/db"
	"github.com/VitorLuizC/ComuniCode-Hackathon/server/gql"
	"github.com/VitorLuizC/ComuniCode-Hackathon/server/users"
	"github.com/gorilla/mux"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	router := mux.NewRouter()

	if err := db.NewSession(); err != nil {
		return
	}
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &gql.Schema,
	})
	router.HandleFunc("/signup", signUp).Methods("POST")
	router.HandleFunc("/login", loginAuth)
	router.Handle("/graphql", requireAuth(handler))
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		res, err := users.Decode(token)
		if err != nil {
			log.Printf("Permission denied: %v", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if res.Id == "" {
			//TODO return error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "id", res.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func loginAuth(w http.ResponseWriter, r *http.Request) {
	var login users.Login
	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(body, &login)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := login.Auth()
	if err == nil {
		w.Header().Set("Authorization", token)
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user users.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer r.Body.Close()
	fmt.Println("response: aaa", user.Save())
}
