package data

import (
	"chitChat/constants"
	"database/sql"
	"fmt"
	"log"
)

func init() {
	connStr := fmt.Sprintf("host=localhost port=5432 "+
		"dbname=chat-app user=postgres password=%s sslmode=prefer "+
		"connect_timeout=10", constants.PASSWORD)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
		}
	}(db)

	if err == nil {
		checkDb(db)
	}

	return
}

func checkDb(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	fmt.Println("Successfully connected to the PostgresSQL database!")
}
