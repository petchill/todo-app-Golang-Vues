package main

import (
	"log"
	"net/http"
	"todo-app-backend/configs"
	"todo-app-backend/controllers"
	"todo-app-backend/global"
	"todo-app-backend/models"

	"github.com/gorilla/mux"
)

// type App struct {
// 	Router *mux.Router
// 	DB     *sql.DB
// }

func InitializeApp(app *global.App) {
	models.InitDb()
	app.DB = models.Db
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", handleGet).Methods(("GET"))
	log.Println("appDB => ", app.DB)
	controllers.InitializeControllers(app)
	log.Println("listen on port: ", configs.Port)
	log.Fatal(http.ListenAndServe(":"+configs.Port, app.Router))
}

func handleGet(res http.ResponseWriter, req *http.Request) {
	db := models.Db
	rows, err := db.Query("SELECT * FROM Todos;")
	if err != nil {
		log.Fatal("database error: ", err)
		res.Write([]byte("error"))
	}
	defer rows.Close()
	// log.Println(rows)
	res.Write([]byte("todos"))
}
