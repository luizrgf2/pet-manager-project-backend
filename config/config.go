package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MYSQL_HOST          = "localhost"
	MYSQL_PORT          = 3306
	MYSQL_NAME_DATABASE = "petManager"
	MYSQL_USER          = "root"
	MYSQL_PASSWORD      = ""
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error to load enviroment variables!")
	}

	MYSQL_HOST = os.Getenv("MYSQL_HOST")
	MYSQL_PORT, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	MYSQL_NAME_DATABASE = os.Getenv("MYSQL_NAME_DATABASE")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
}
