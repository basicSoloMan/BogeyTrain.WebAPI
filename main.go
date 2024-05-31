package main

import (
	"BogeyTrain/pkg/driver"
	"database/sql"
	"fmt"
)

func main() {
	db, err := driver.Connect()
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	fmt.Println("Connection Successful")
}
