package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/luizrgf2/pet-manager-project-backend/config"
)

var (
	DB *sql.DB
)

func create_uri_conn_db() string {
	if len(config.MYSQL_PASSWORD) == 0 {
		uri := fmt.Sprintf("%s@tcp(%s:%d)/%s", config.MYSQL_USER, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_NAME_DATABASE)
		return uri
	} else {
		uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_NAME_DATABASE)
		return uri
	}
}

func init() {
	uri := create_uri_conn_db()
	db, err := sql.Open("mysql", uri)
	if err != nil {
		panic("Error to connect database with uri " + uri)
	} else {
		fmt.Println("Success to connect db")
	}
	DB = db
}
