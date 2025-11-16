package cmd

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func create_db() {
	db, err := openSqlConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, _, _, _, name := loadEnvVariables()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
	if err != nil {
		log.Fatalf("database creation error: %v", err)
	}

	fmt.Println("Database successfully created")
}
