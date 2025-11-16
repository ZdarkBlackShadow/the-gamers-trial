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

	//Question3 : Minimap AOE4
	image3Search := entity.Image{Path: "/image3.png"}
	image3 := image3Search
	err = db.FirstOrCreate(&image3, image3Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 3: %v", err)
	}

	question3Search := entity.Question{ImageID: image3.ID}
	question3 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image3.ID,
	}
	err = db.FirstOrCreate(&question3, question3Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 3: %v", err)
	}

	options3 := []entity.Option{
		{Text: "Age Of Empires 4", QuestionID: question3.ID, Correct: true},
		{Text: "Starcraft 2", QuestionID: question3.ID, Correct: false},
		{Text: "Starfield", QuestionID: question3.ID, Correct: false},
		{Text: "The Witcher 3", QuestionID: question3.ID, Correct: false},
	}
	for _, opt := range options3 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question3.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question4 : Minimap Assassin's creed black flag
	image4Search := entity.Image{Path: "/image4.png"}
	image4 := image4Search
	err = db.FirstOrCreate(&image4, image4Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 4: %v", err)
	}

	question4Search := entity.Question{ImageID: image4.ID}
	question4 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image4.ID,
	}
	err = db.FirstOrCreate(&question4, question4Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 4: %v", err)
	}

	options4 := []entity.Option{
		{Text: "Saints Row IV", QuestionID: question4.ID, Correct: false},
		{Text: "The Witcher 3", QuestionID: question4.ID, Correct: false},
		{Text: "Assassin's creed black flag", QuestionID: question4.ID, Correct: true},
		{Text: "Far Cry 6", QuestionID: question4.ID, Correct: false},
	}
	for _, opt := range options4 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question4.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question5 : Minimap Warzone
	image5Search := entity.Image{Path: "/image5.png"}
	image5 := image5Search
	err = db.FirstOrCreate(&image5, image5Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 5: %v", err)
	}

	question5Search := entity.Question{ImageID: image5.ID}
	question5 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image5.ID,
	}
	err = db.FirstOrCreate(&question5, question5Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 5: %v", err)
	}

	options5 := []entity.Option{
		{Text: "Battlefield V", QuestionID: question5.ID, Correct: false},
		{Text: "Call of Duty Warzone", QuestionID: question5.ID, Correct: true},
		{Text: "Saints Row IV", QuestionID: question5.ID, Correct: false},
		{Text: "Counter-Strike 2", QuestionID: question5.ID, Correct: false},
	}
	for _, opt := range options5 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question5.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question6 : Minimap Cyberpunk 2077
	image6Search := entity.Image{Path: "/image6.png"}
	image6 := image6Search
	err = db.FirstOrCreate(&image6, image6Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 6: %v", err)
	}

	question6Search := entity.Question{ImageID: image6.ID}
	question6 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image6.ID,
	}
	err = db.FirstOrCreate(&question6, question6Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 6: %v", err)
	}

	options6 := []entity.Option{
		{Text: "Grand Theft Auto V", QuestionID: question6.ID, Correct: false},
		{Text: "The Witcher 3", QuestionID: question6.ID, Correct: false},
		{Text: "Ghostrunner", QuestionID: question6.ID, Correct: false},
		{Text: "Cyberpunk 2077", QuestionID: question6.ID, Correct: true},
	}
	for _, opt := range options6 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question6.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question7 : Borderlands 2
	image7Search := entity.Image{Path: "/image7.png"}
	image7 := image7Search
	err = db.FirstOrCreate(&image7, image7Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 7: %v", err)
	}

	question7Search := entity.Question{ImageID: image7.ID}
	question7 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image7.ID,
	}
	err = db.FirstOrCreate(&question7, question7Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 7: %v", err)
	}

	options7 := []entity.Option{
		{Text: "Borderlands 2", QuestionID: question7.ID, Correct: false},
		{Text: "Tiny Tina's Wonderlands", QuestionID: question7.ID, Correct: false},
		{Text: "Mad Max", QuestionID: question7.ID, Correct: true},
		{Text: "Rage 2", QuestionID: question7.ID, Correct: false},
	}
	for _, opt := range options7 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question7.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question8 : Minimap Burnout Paradise
	image8Search := entity.Image{Path: "/image8.png"}
	image8 := image8Search
	err = db.FirstOrCreate(&image8, image8Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 8: %v", err)
	}

	question8Search := entity.Question{ImageID: image8.ID}
	question8 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image8.ID,
	}
	err = db.FirstOrCreate(&question8, question8Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 8: %v", err)
	}

	options8 := []entity.Option{
		{Text: "Forza Horizon 5", QuestionID: question8.ID, Correct: false},
		{Text: "Burnout Paradise", QuestionID: question8.ID, Correct: true},
		{Text: "Need for Speed: Most Wanted", QuestionID: question8.ID, Correct: false},
		{Text: "The Crew Motorfest", QuestionID: question8.ID, Correct: false},
	}
	for _, opt := range options8 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question8.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question9 : Minimap Darksiders 2
	image9Search := entity.Image{Path: "/image9.png"}
	image9 := image9Search
	err = db.FirstOrCreate(&image9, image9Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 9: %v", err)
	}

	question9Search := entity.Question{ImageID: image9.ID}
	question9 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image9.ID,
	}
	err = db.FirstOrCreate(&question9, question9Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 9: %v", err)
	}

	options9 := []entity.Option{
		{Text: "Darksiders 2", QuestionID: question9.ID, Correct: true},
		{Text: "Elden Ring", QuestionID: question9.ID, Correct: false},
		{Text: "Dark Souls 2", QuestionID: question9.ID, Correct: false},
		{Text: "World of Warcraft", QuestionID: question9.ID, Correct: false},
	}
	for _, opt := range options9 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question9.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	fmt.Println("Default data successfully inserted")
}
