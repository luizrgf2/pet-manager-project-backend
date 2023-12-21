package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	MYSQL_HOST                  = "localhost"
	MYSQL_PORT                  = 3306
	MYSQL_NAME_DATABASE         = "petManager"
	MYSQL_USER                  = "root"
	MYSQL_PASSWORD              = ""
	JWT_KEY                     = ""
	SMTP_EMAIL_SENDER           = ""
	SMTP_SERVER                 = ""
	SMTP_PASSWORD               = ""
	SMTP_PORT                   = 0
	SMTP_EMAIL_RECEIVER_TO_TEST = ""
)

func init() {
	base_path, _ := os.Getwd()
	path_complete := ""

	if strings.Contains(base_path, "/test") {
		path_complete = strings.Split(base_path, "/test")[0]
	} else if strings.Contains(base_path, "/internal") {
		path_complete = strings.Split(base_path, "/internal")[0]
	}

	enviroment := os.Getenv("ENVIROMENT")

	file_env_to_load := ""

	if enviroment == "true" {
		file_env_to_load = ".env.test"
	} else {
		file_env_to_load = ".env"
	}

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
	JWT_KEY = os.Getenv("JWT_KEY")
	SMTP_SERVER = os.Getenv("SMTP_SERVER")
	SMTP_EMAIL_SENDER = os.Getenv("SMTP_EMAIL_SENDER")
	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	SMTP_PORT, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
	SMTP_EMAIL_RECEIVER_TO_TEST = os.Getenv("SMTP_EMAIL_RECEIVER_TO_TEST")

	if err != nil {
		panic(err.Error())
	}
}
