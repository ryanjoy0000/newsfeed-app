package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryanjoy0000/newsfeed-app/kit/customErr"
)

const (
	mysql_user    = "root"
	mysql_pwd     = "Password@1"
	mysql_addr    = "localhost"
	mysql_port    = "3306"
	mysql_db_name = "news_db"
)

func InitMySQLDB() *sql.DB {
	connInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		mysql_user,
		mysql_pwd,
		mysql_addr,
		mysql_port,
		mysql_db_name,
	)

	fmt.Println("Connecting to DB: ", mysql_db_name, " on ", mysql_addr, ":", mysql_port)
	db, err := sql.Open("mysql", connInfo)
	customErr.HandleErr(err)

	err = db.Ping()
	customErr.HandleErr(err)
	fmt.Println("DB connected...")

	return db
}

func CloseDBConn(db *sql.DB) {
	db.Close()
	fmt.Println("DB connection closed...")
}
