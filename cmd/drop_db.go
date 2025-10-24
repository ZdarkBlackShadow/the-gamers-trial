package cmd

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func dropDb() {
	db, err := openSqlConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, _, _, _, name := loadEnvVariables()
	_, err = db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		log.Fatalf("Database deletion error: %v", err)
	}

	fmt.Println("Database successfully deleted")
}
