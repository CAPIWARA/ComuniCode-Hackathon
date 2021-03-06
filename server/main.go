package main

import (
	"comunicode/server/db"
	"comunicode/server/gql"
	"comunicode/server/users"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		if r.Method == "OPTIONS" {
			log.Println("options method ")
			return
		}
		token := r.Header.Get("Authorization")
		log.Println("token: ", token)
		res, err := users.Decode(token)
		if err != nil {
			log.Printf("Permission denied: %v", err)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if res.Id == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(r.Context(), "email", res.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func loginAuth(w http.ResponseWriter, r *http.Request) {
	var login users.Login
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	if r.Method == "OPTIONS" {
		log.Println("options method ")
		return
	}
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
		w.WriteHeader(http.StatusOK)
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
	if err := user.Save(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
