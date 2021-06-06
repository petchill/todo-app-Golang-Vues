package models

import (
	"database/sql"
	"fmt"
	"log"

	"todo-app-backend/configs"

	_ "github.com/go-sql-driver/mysql"
)

// Db is
var Db *sql.DB

func InitDb() {
	dbConfig := configs.DB
	dbConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	log.Println("db connect : ", dbConnect)
	db, err := sql.Open("mysql", dbConnect)
	if err != nil {
		log.Fatal("Open connection failed: ", err)
	}
	log.Println("connected")
	Db = db
	// rows, err := db.Query("SELECT * FROM Todos;")
	// if err != nil {
	// 	log.Fatal("database error: ", err)
	// 	// res.Write([]byte("error"))
	// }
	// log.Println("TODO DB => ", rows)
	// defer rows.Close()
	// db.Close()
	// return db
}
