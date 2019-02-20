package main

import (
	"data"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

//loadConfig reconfigures it self.
func loadConfig() {

}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func parseTemplateFiles(filesnames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filesnames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(template.ParseFiles(files...))
	return
}

// For Logging purposes

//danger logs and returns Error Prefix typed log.
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}
