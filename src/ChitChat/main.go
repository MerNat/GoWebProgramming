package main

import (
	"net/http"
	// "reflect"
	// "runtime"
	"time"
	"github.com/gorilla/mux"
)

func main() {

	p("ChitChat", version(), "started at", config.Address)

	muxOld := http.NewServeMux()
	mux2 := mux.NewRouter()
	files := http.FileServer(http.Dir(config.Static))
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
		Addr:    config.Address,
		Handler: mux2,
		ReadHeaderTimeout: time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout: time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

// func log(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
// 		fmt.Println("Handler function called - " + name)
// 		h(w, r)
// 	}
// }