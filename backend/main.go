package main

import (
	"net/http"
	"todo-app-backend/global"
)

// YourHandler is example
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	app := global.App{}
	InitializeApp(&app)
	// fmt.Println("Hello")
	// models.InitDb()
	// r := mux.NewRouter()
	// r.HandleFunc("/", controllers.TodoController).Methods("GET")
	// log.Fatal(http.ListenAndServe(":4100", r))
}
