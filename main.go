package main

import (
	"net/http"

	"server/handler"
	"server/logger"
	"server/middleware"
)

func main() {
	http.Handle("/hello", middleware.Handle(http.HandlerFunc(handler.HelloHandle)))

	logger.Log.Info("server running...")

	http.ListenAndServe(":8080", nil)
}
