package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func SetupDatabase() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "college"

	var err error
	DB, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	fmt.Println("Database connected!")
}
