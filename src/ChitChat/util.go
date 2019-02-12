package main


import (
	"encoding/json"
	"errors"
	"fmt"
	"data"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func session(w http.ResponseWriter, r *http.Request)(sess data.Session, err error){
	
}