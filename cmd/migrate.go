package cmd

import (
	"fmt"
	"log"

	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
)

func migrate() {
	db, err := openDBConnection()
	if err != nil {
		log.Fatalf("Database connexion error: %v", err)
	}

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Image{},
		&entity.Question{},
		&entity.Option{},
	)

	if err != nil {
		log.Fatalf("migration errror: %v", err)
	}

	fmt.Println("Successfully implemented migrations")
}
