package controllers

import (
	"log"
	"net/http"
	"todo-app-backend/global"
	"todo-app-backend/models"
	// "todo-app-backend/models"
)

type ITodo struct {
	Id          int
	Topic       string
	Describtion string
	Status      string
	Created_at  string
	Updated_at  string
	Deadline    string
	Created_by  int
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
	app.Router.HandleFunc("/todos", handleGet).Methods("GET")
}

func handleGet(res http.ResponseWriter, req *http.Request) {
	db := models.Db
	_, err := db.Query("SELECT * FROM Todos")
	if err != nil {
		log.Fatal("database error: ", err)
		res.Write([]byte("error"))
	}
	res.Write([]byte("todos"))
	// log.Println(rows)
	// var dataList []ITodo
	// for rows.Next() {
	// 	var data ITodo
	// 	err = rows.Scan(&data.Id, &data.Topic, &data.Describtion, &data.Status, &data.Created_at, &data.Updated_at, &data.Deadline, &data.Created_by)
	// 	if err != nil {
	// 		log.Fatal("database error", err)
	// 		res.Write([]byte("error"))
	// 	}
	// 	// res.Header().Set("Content-Type", "application/json")
	// 	// json.NewEncoder(res).Encode(data)
	// 	dataList = append(dataList, data)
	// 	res.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// 	res.Header().Set("Access-Control-Allow-Origin", "*")
	// 	// log.Println(resBody)
	// 	// log.Println(json.Marshal(dataList))
	// 	res.Header().Set("Content-Type", "application/json")
	// 	// json.NewEncoder(res).Encode(dataList[0])
	// }
	// result, _ := json.Marshal(dataList)
	// res.Write(result)
}
