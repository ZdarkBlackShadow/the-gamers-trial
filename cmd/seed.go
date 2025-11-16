package cmd

import (
	"fmt"
	"log"

	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
)

func seed() {
	db, err := openDBConnection()
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	//Question1 : Minimap Apex Legends
	imageSearch := entity.Image{Path: "/image.png"}
	defaultImage := imageSearch

	err = db.FirstOrCreate(&defaultImage, imageSearch).Error
	if err != nil {
		log.Fatalf("Erreur seed Image: %v", err)
	}

	questionSearch := entity.Question{ImageID: defaultImage.ID}
	defaultQuestion := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
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

	//Question2 : Minimap Animal Crossing
	image2Search := entity.Image{Path: "/image2.png"}
	image2 := image2Search
	err = db.FirstOrCreate(&image2, image2Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 2: %v", err)
	}

	question2Search := entity.Question{ImageID: image2.ID}
	question2 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image2.ID,
	}
	err = db.FirstOrCreate(&question2, question2Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 2: %v", err)
	}

	options2 := []entity.Option{
		{Text: "Animal Crossing", QuestionID: question2.ID, Correct: true},
		{Text: "Pokemon XY", QuestionID: question2.ID, Correct: false},
		{Text: "Zelda Wind Waker", QuestionID: question2.ID, Correct: false},
		{Text: "Palworld", QuestionID: question2.ID, Correct: false},
	}
	for _, opt := range options2 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question2.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	fmt.Println("Default data successfully inserted")
}
