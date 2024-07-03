package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	writeString, err := io.WriteString(w, "Hello World")
	if err != nil {
		return
	}
	log.Println(writeString)
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	username := p.ByName("username")
	writeString, err := io.WriteString(w, username)
	if err != nil {
		return
	}
	log.Println(writeString)
}
