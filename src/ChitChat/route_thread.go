package main

import (
	"data"
	"net/http"
	"github.com/gorilla/mux"
)

func newThread(w http.ResponseWriter, request *http.Request) {
	_, err := session(w, request)
	if err != nil {
		http.Redirect(w, request, "/", 302)
	} else {
		generateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

func createThread(w http.ResponseWriter, request *http.Request) {
	session, err := session(w, request)
	if err != nil {
		http.Redirect(w, request, "/", 302)
	}

	err = request.ParseForm()

	if err != nil {
		danger(err, "Can't Parse Form")
	}

	user, err := session.User()

	_, err = user.CreateThread(request.PostFormValue("topic"))

	http.Redirect(w, request, "/", 302)
}

func readThread(w http.ResponseWriter, request *http.Request){
	requestVals := mux.Vars(request)
	id := requestVals["id"]

	thread, _ := data.ThreadByUUID(id)

	_, err := session(w, request)

	if err!= nil{
		// public one

		// whenever possible we should use a generated pointers to represent a Struct 
		generateHTML(w, &thread, "layout", "public.navbar", "public.thread")
	}else{
		// private one
		generateHTML(w, &thread, "layout", "private.navbar", "private.thread")
	}

}

func postThread(w http.ResponseWriter, request *http.Request){
	session, err := session(w, request)
	if err!=nil{
		http.Redirect(w, request, "/", 302)
	}

	err = request.ParseForm()

	if err!=nil{
		// error found
		danger(err, "Error found when parsing form")
	}

	body := request.PostFormValue("body")

	uuidThread := request.PostFormValue("uuid")

	thread, err := data.ThreadByUUID(uuidThread)

	if err!=nil {
		danger (err, "Error When getting a thread")
	}

	user, _ := session.User()

	user.CreatePost(thread, body)

	http.Redirect(w, request, "/thread/read/"+thread.Uuid, 302)
}