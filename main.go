package main

import (
	"log"
	"net/http"
	"server/handler"
	"server/middleware"
)

func main() {
	http.Handle("/hello", middleware.Handle(http.HandlerFunc(handler.HelloHandle)))
	log.Println("server running...")
	http.ListenAndServe(":8080", nil)
}
