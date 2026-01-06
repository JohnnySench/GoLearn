package main

import (
	"restapi/http"
	"restapi/todo"
)

func main() {
	todos := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todos)
	httpServer := http.NewHTTPServer(httpHandlers)
	httpServer.StartServer()
}
