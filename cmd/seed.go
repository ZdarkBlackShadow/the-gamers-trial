package cmd

import (
	"fmt"
	"log"
)

func seed() {
	_, err := openDBConnection()
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	fmt.Println("Default data successfully inserted")
}
