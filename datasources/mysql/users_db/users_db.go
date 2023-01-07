package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

/*
const (
	mysql_users_username = "USER"
	mysql_users_password = "PASSWORD"
	mysql_users_host     = "HOST"
	mysql_users_schema   = "SCHEMA"
)
*/

var (
	Client *sql.DB
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	schema := os.Getenv("SCHEMA")

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	err = Client.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Database successfully configured")

}
