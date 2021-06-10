package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"todo-app-backend/global"
	"todo-app-backend/models"
	// "todo-app-backend/models"
)

type ITodo struct {
	Id          int    `json:"id"`
	Topic       string `json:"topic"`
	Describtion string `json:"describtion"`
	Status      string `json:"status"`
	Deadline    string `json:"deadline"`
	Color       string `json:"color"`
	Created_by  int    `json:"created_by"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type IResponse struct {
	ok     bool
	status int
	data   interface{}
}

// func TodoController(res http.ResponseWriter, req *http.Request) {
// 	db := models.Db
// 	defer db.Close()
// 	rows, err := db.Query("SELECT * FROM Todos")
// 	if err != nil {
// 		log.Fatal("database error", err)
// 	}
// 	log.Println(rows)
// 	var dataList []ITodo
// 	for rows.Next() {
// 		var data ITodo
// 		err = rows.Scan(&data.id, &data.topic, &data.describtion, &data.createAt, &data.color)
// 		if err != nil {
// 			log.Fatal("database error", err)
// 		}
// 		// res.Header().Set("Content-Type", "application/json")
// 		// json.NewEncoder(res).Encode(data)
// 		dataList = append(dataList, data)
// 	}
// 	// resBody := IResponse{
// 	// 	ok:     true,
// 	// 	status: 200,
// 	// 	data:   dataList,
// 	// }
// 	res.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	res.Header().Set("Access-Control-Allow-Origin", "*")
// 	// log.Println(resBody)
// 	// log.Println(json.Marshal(dataList))
// 	res.Header().Set("Content-Type", "application/json")
// 	// json.NewEncoder(res).Encode(dataList[0])
// 	res.Write(json.Marshal(dataList))
// 	// models.Db
// 	// res.Write([]byte("Todo"))
// }

func NewToDoController(app *global.App) {
	app.Router.HandleFunc("/todos", handleGetTodos).Methods("GET")
}

func handleGetTodos(res http.ResponseWriter, req *http.Request) {
	db := models.Db
	rows, err := db.Query("SELECT * FROM Todos")
	if err != nil {
		log.Fatal("database error: ", err)
		res.Write([]byte("error"))
	}
	defer rows.Close()
	// res.Write([]byte("todos"))
	log.Println(rows)
	var dataList []ITodo
	for rows.Next() {
		var data ITodo
		err = rows.Scan(&data.Id, &data.Topic, &data.Describtion, &data.Status, &data.Deadline, &data.Color, &data.Created_by, &data.Created_at, &data.Updated_at)
		if err != nil {
			log.Fatal("database error", err)
			res.Write([]byte("error"))
		}
		log.Println("data topic => ", data)
		// res.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(res).Encode(data)
		dataList = append(dataList, data)
	}
	// res.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Content-Type", "application/json")
	// log.Println("dataList => ", dataList)
	// result, err := json.Marshal(dataList)
	// if err != nil {
	// 	log.Fatal("marshal error", err)
	// }
	// log.Println("result => ", result)
	json.NewEncoder(res).Encode(dataList)
}
