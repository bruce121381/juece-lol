package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {

	router := httprouter.New()

	router.GET("/videos/:vid-id", streamHandler)

	router.POST("/video/:vid-id", uploadHandler)

	return router

}

func main() {

	r := RegisterHandlers()
	/**
	监听9000端口
	*/
	http.ListenAndServe(":9000", r)

}
