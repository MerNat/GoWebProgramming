package main

import (
	_ "data"
	"net/http"
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