package main

import (
	"fmt"
	"net/http"
	// "reflect"
	// "runtime"

	"github.com/gorilla/mux"
)

func main() {
	muxOld := http.NewServeMux()
	mux2 := mux.NewRouter()
	files := http.FileServer(http.Dir("public"))
	muxOld.Handle("/static/", http.StripPrefix("/static/", files)) // builtIn
	mux2.PathPrefix("/static/").Handler(http.StripPrefix("/static/", files)) // for gorilla mux

	// mux2.Handle("/{mmaa}/", log(index))
	mux2.HandleFunc("/", index)
	muxOld.HandleFunc("/", index)
	// mux.HandleFunc("/err", err)
	mux2.HandleFunc("/login", login)
	mux2.HandleFunc("/logout", logout)
	mux2.HandleFunc("/signup", signup)
	mux2.HandleFunc("/signup_account", signupAccount)
	mux2.HandleFunc("/authenticate", authenticate)
	mux2.HandleFunc("/thread/new", newThread)
	mux2.HandleFunc("/thread/create", createThread)
	mux2.HandleFunc("/thread/post", postThread)
	mux2.HandleFunc("/thread/read/{id}", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux2,
	}

	fmt.Println("Server Started")
	server.ListenAndServe()
}

// func log(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
// 		fmt.Println("Handler function called - " + name)
// 		h(w, r)
// 	}
// }