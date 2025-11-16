package cmd

import (
	"fmt"
	"log"

	// Supposons que votre connexion db utilise gorm.DB

	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
)

func seed() {
	db, err := openDBConnection()
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	imageSearch := entity.Image{Path: "/image.png"}
	defaultImage := imageSearch

	err = db.FirstOrCreate(&defaultImage, imageSearch).Error
	if err != nil {
		log.Fatalf("Erreur seed Image: %v", err)
	}

	questionSearch := entity.Question{Content: "Dans quel jeu est utilisée cette minimap ?"}
	defaultQuestion := entity.Question{
		Content: questionSearch.Content,
		ImageID: defaultImage.ID,
	}

	err = db.FirstOrCreate(&defaultQuestion, questionSearch).Error
	if err != nil {
		log.Fatalf("Erreur seed Question: %v", err)
	}

	questionID := defaultQuestion.ID

	defaultOptions := []entity.Option{
		{Text: "Call of Duty MW3", QuestionID: questionID, Correct: false},
		{Text: "Fortnite", QuestionID: questionID, Correct: false},
		{Text: "Apex Legends", QuestionID: questionID, Correct: true},
		{Text: "Halo 6", QuestionID: questionID, Correct: false},
	}

	for _, do := range defaultOptions {
		searchOption := entity.Option{Text: do.Text, QuestionID: questionID}
		err := db.FirstOrCreate(&do, searchOption).Error
		if err != nil {
			log.Fatalf("Erreur seed Option %s: %v", do.Text, err)
		}
	}

	fmt.Println("Default data successfully inserted")
}
