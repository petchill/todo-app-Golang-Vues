package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-app-backend/global"
	"todo-app-backend/models"
)

func NewUserController(app *global.App) {
	app.Router.HandleFunc("/users", handlerGetUsers).Methods("GET")
}

func handlerGetUsers(res http.ResponseWriter, req *http.Request) {
	userList, err := models.GetUserList()
	if err != nil {
		log.Fatal(err)
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(userList)
}
