package main

import (
	"data"
	"net/http"
	"fmt"
)

func err(writer http.ResponseWriter, request *http.Request){
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err == nil{
		generateHTML(writer,vals.Get("msg"),"layout","private.navbar","error")
	}else{
		generateHTML(writer,vals.Get("msg"),"layout","public.navbar","error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// values := mux.Vars(r)
	// fmt.Println("The value of mmaa: ", values["mmaa"], r.Method)
	threads, err := data.Threads()
	w.WriteHeader(404) //Status Code
	if err == nil {
		_, err := session(w, r)
		if err == nil {
			generateHTML(w, &threads, "layout", "private.navbar", "index")
		} else {
			generateHTML(w, &threads, "layout", "public.navbar", "index")
		}
	} else {
		fmt.Println("Error Found: ", err.Error())
	}
}