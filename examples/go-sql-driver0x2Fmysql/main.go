package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysql_addr string
	db_name    string
	basic_auth string
)

func init() {
	flag.StringVar(&mysql_addr, "mysql-server", "172.17.0.5:3306", "MySQL host and port for TCP connection")
	flag.StringVar(&db_name, "db-name", "mysql", "Database name to access")
	flag.StringVar(&basic_auth, "credentials", "root:password", "User and password for basic auth")
}

func main() {
	flag.Parse()

	connection := fmt.Sprintf("%s@tcp(%s)/%s", basic_auth, mysql_addr, db_name)
	db, err := sql.Open("mysql", connection)
	//db, err := sql.Open("mysql", "root:password@tcp(172.17.0.5:3306)/mysql")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Use the DB normally, execute the querys etc

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT Host, User FROM mysql.user WHERE User = ?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var hostname, username string // we "scan" the result in here

	// Query the square-number of 1
	err = stmtOut.QueryRow("root").Scan(&hostname, &username) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The result is: %v\n", []interface{}{hostname, username})
}
