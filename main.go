package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Read the connection string from an environment variable
	connectionString := os.Getenv("MYSQL_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("MYSQL_CONNECTION_STRING was not set")
	}

	// Create a connection to the database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verify we can connect
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to the database! 👍")

	// Query the gophers table
	rows, err := db.Query("select name from gophers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the results
	log.Println("   Name    ")
	log.Println("-----------")
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(name)
	}

	// Check for errors when reading the row data
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
