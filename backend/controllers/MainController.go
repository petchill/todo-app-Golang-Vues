package controllers

import (
	"net/http"
	"todo-app-backend/global"
)

var app *global.App

func InitializeControllers(initialApp *global.App) {
	app = initialApp
	app.Router.HandleFunc("/healthz", handlerHealthZ).Methods("GET")
	NewToDoController(initialApp)
	NewUserController(initialApp)
}

func handlerHealthZ(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("OK"))
}
