package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"stream_media/api/handlers"
	_ "stream_media/api/handlers"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handlers.CreateUserHandler)
	router.POST("/user/:username", handlers.UserLoginHandler)
	return router
}

func main() {
	router := RegisterHandlers()
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
