package models

import (
	"log"
)

type IUser struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Password_salt string `json:"password_salt"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
}

func GetUserList() ([]IUser, error) {
	const qString = "SELECT * FROM Users"
	rows, err := Db.Query(qString)
	if err != nil {
		log.Fatal("database error: ", err)
		return nil, err
	}
	var userList []IUser
	for rows.Next() {
		var data IUser
		err = rows.Scan(&data.Id, &data.Email, &data.Password, &data.Password_salt, &data.Created_at, &data.Updated_at)
		if err != nil {
			log.Fatal("database error", err)
			return nil, err
		}
		userList = append(userList, data)
	}
	return userList, nil
}
