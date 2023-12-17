package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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
	base_path, _ := os.Getwd()
	path_complete := strings.Split(base_path, "/test")[0]
	enviroment := os.Getenv("ENVIROMENT")

	file_env_to_load := ""

	if enviroment == "true" {
		file_env_to_load = ".env.test"
	} else {
		file_env_to_load = ".env"
	}

	fmt.Println(enviroment)

	err := godotenv.Load(path_complete + "/" + file_env_to_load)
	if err != nil {
		fmt.Println(err.Error())
		panic("Error to load enviroment variables!")
	}

	MYSQL_HOST = os.Getenv("MYSQL_HOST")
	MYSQL_PORT, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	MYSQL_NAME_DATABASE = os.Getenv("MYSQL_NAME_DATABASE")
	MYSQL_USER = os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
}
