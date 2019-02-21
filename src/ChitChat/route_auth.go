package main

import (
	"data"
	"net/http"
	"fmt"
)

//login redirects to the login form only
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("login.layout", "login")
	t.ExecuteTemplate(writer, "layout", nil)
}

func signup(writer http.ResponseWriter, request *http.Request){
	generateHTML(writer, nil, "layout", "public.navbar", "signup")
}


func signupAccount(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()
	if err != nil{
		danger(err, "Can not parse form")
	}
	user := data.User{
		Name: request.PostFormValue("name"),
		Email: request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err = user.Create(); err!=nil{
		danger(err, "Can not create user")
	}

	http.Redirect(writer, request, "/login", 302)
}


func authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil{
			danger(err, "Can not create a session!")
		}
		fmt.Println(session)
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		fmt.Println(cookie)
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		generateHTML(w, true, "login.layout", "login")
		// http.Redirect(w, r, "/login", 302)
	}
}


//logout logs out and destroy session
func logout(writer http.ResponseWriter, request *http.Request){
	cookie, err := request.Cookie("_cookie")
	if err == nil{
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
		http.Redirect(writer, request, "/", 302)
	}else{
		warning(err, "Failed to get cookie")
	}
}