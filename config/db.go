package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB creates a connection to the MySQL database
func ConnectDB() *sql.DB {
	// Update with your actual MySQL username, password, and database name
	dsn := "root:Sql@#$123@tcp(127.0.0.1:3306)/userdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Connected to MySQL Database Successfully!")
	return db
}
