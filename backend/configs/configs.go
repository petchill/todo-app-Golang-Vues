package configs

import "os"

type IConfig struct {
	Port string
	DB   DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

var DB = DBConfig{
	Host:     os.Getenv("DB_HOST"),
	Port:     os.Getenv("DB_PORT"),
	Username: os.Getenv("DB_USERNAME"),
	Password: os.Getenv("DB_PASSWORD"),
	Name:     os.Getenv("DB_NAME"),
}

var Port string = os.Getenv("PORT")

// var Configs = &IConfig{
// 	Port: os.Getenv("PORT"),
// 	DB: DBConfig{
// 		Host:     os.Getenv("DB_HOST"),
// 		Port:     os.Getenv("DB_PORT"),
// 		Username: os.Getenv("DB_USERNAME"),
// 		Password: os.Getenv("PASSWORD"),
// 		Name:     os.Getenv("DB_NAME"),
// 	},
// }

// func readConfig() config {

// }
