package main

import (
	"fmt"
	_"github.com/lib/pq"
	"net/http"
)

func main(){
	fmt.Println("Hallo")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080",nil)
}

func handler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}