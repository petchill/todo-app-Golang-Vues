package controllers

import (
	"todo-app-backend/global"
)

var app *global.App

func InitializeControllers(initialApp *global.App) {
	app = initialApp
	NewToDoController(initialApp)
}
