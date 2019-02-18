package main

import (
	"data"
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/gorilla/mux"
)

func main() {
	muxOld := http.NewServeMux()
	mux2 := mux.NewRouter()
	files := http.FileServer(http.Dir("public"))
	muxOld.Handle("/static/", http.StripPrefix("/static/", files))
	mux2.Handle("/static/", http.StripPrefix("/static", files))

	// mux2.Handle("/{mmaa}/", log(index))
	mux2.HandleFunc("/", index)
	muxOld.HandleFunc("/", index)
	// mux.HandleFunc("/err", err)
	// mux.HandleFunc("/login", login)
	// mux.HandleFunc("/logout", logout)
	// mux.HandleFunc("/signup", signup)
	// mux.HandleFunc("/signup_account", signupAccount)
	// mux.HandleFunc("/authenticate", authenticate)
	// mux.HandleFunc("/thread/new", newThread)
	// mux.HandleFunc("/thread/create", createThread)
	// mux.HandleFunc("/thread/post", postThread)
	// mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux2,
	}

	fmt.Println("Server Started")
	server.ListenAndServe()
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// values := mux.Vars(r)
	// fmt.Println("The value of mmaa: ", values["mmaa"], r.Method)

	h := r.Header["User-Agent"]
	fmt.Println(h)
	threads, err := data.Threads()
	w.WriteHeader(404) //Status Code
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	} else {
		fmt.Println("Error Found: ", err.Error())
	}
}
