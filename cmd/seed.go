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
		{Text: "Borderlands 2", QuestionID: question7.ID, Correct: true},
		{Text: "Tiny Tina's Wonderlands", QuestionID: question7.ID, Correct: false},
		{Text: "Mad Max", QuestionID: question7.ID, Correct: false},
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

	//Question10 : Minimap Valorant
	image10Search := entity.Image{Path: "/image10.png"}
	image10 := image10Search
	err = db.FirstOrCreate(&image10, image10Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 10: %v", err)
	}

	question10Search := entity.Question{ImageID: image10.ID}
	question10 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image10.ID,
	}
	err = db.FirstOrCreate(&question10, question10Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 10: %v", err)
	}

	options10 := []entity.Option{
		{Text: "Overwatch 2", QuestionID: question10.ID, Correct: false},
		{Text: "Valorant", QuestionID: question10.ID, Correct: true},
		{Text: "Paladins", QuestionID: question10.ID, Correct: false},
		{Text: "Splitgate", QuestionID: question10.ID, Correct: false},
	}
	for _, opt := range options10 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question10.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question11 : Minimap Warframe
	image11Search := entity.Image{Path: "/image11.png"}
	image11 := image11Search
	err = db.FirstOrCreate(&image11, image11Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 11: %v", err)
	}

	question11Search := entity.Question{ImageID: image11.ID}
	question11 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image11.ID,
	}
	err = db.FirstOrCreate(&question11, question11Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 11: %v", err)
	}

	options11 := []entity.Option{
		{Text: "Destiny 2", QuestionID: question11.ID, Correct: false},
		{Text: "Warframe", QuestionID: question11.ID, Correct: true},
		{Text: "Anthem", QuestionID: question11.ID, Correct: false},
		{Text: "Halo Infinite", QuestionID: question11.ID, Correct: false},
	}
	for _, opt := range options11 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question11.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question12 : Minimap Heroes of the Storm
	image12Search := entity.Image{Path: "/image12.png"}
	image12 := image12Search
	err = db.FirstOrCreate(&image12, image12Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 12: %v", err)
	}

	question12Search := entity.Question{ImageID: image12.ID}
	question12 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image12.ID,
	}
	err = db.FirstOrCreate(&question12, question12Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 12: %v", err)
	}

	options12 := []entity.Option{
		{Text: "League of Legends", QuestionID: question12.ID, Correct: false},
		{Text: "Dota 2", QuestionID: question12.ID, Correct: false},
		{Text: "Smite", QuestionID: question12.ID, Correct: false},
		{Text: "Heroes of the Storm", QuestionID: question12.ID, Correct: true},
	}
	for _, opt := range options12 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question12.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question13 : Minimap Genshin Impact
	image13Search := entity.Image{Path: "/image13.png"}
	image13 := image13Search
	err = db.FirstOrCreate(&image13, image13Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 13: %v", err)
	}

	question13Search := entity.Question{ImageID: image13.ID}
	question13 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image13.ID,
	}
	err = db.FirstOrCreate(&question13, question13Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 13: %v", err)
	}

	options13 := []entity.Option{
		{Text: "Genshin Impact", QuestionID: question13.ID, Correct: true},
		{Text: "Tower of Fantasy", QuestionID: question13.ID, Correct: false},
		{Text: "Immortals Fenyx Rising", QuestionID: question13.ID, Correct: false},
		{Text: "Blue Protocol", QuestionID: question13.ID, Correct: false},
	}
	for _, opt := range options13 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question13.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question14 : Minimap Gears 5
	image14Search := entity.Image{Path: "/image14.png"}
	image14 := image14Search
	err = db.FirstOrCreate(&image14, image14Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 14: %v", err)
	}

	question14Search := entity.Question{ImageID: image14.ID}
	question14 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image14.ID,
	}
	err = db.FirstOrCreate(&question14, question14Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 14: %v", err)
	}

	options14 := []entity.Option{
		{Text: "No Man's Sky", QuestionID: question14.ID, Correct: false},
		{Text: "Gears 5", QuestionID: question14.ID, Correct: true},
		{Text: "Elite Dangerous", QuestionID: question14.ID, Correct: false},
		{Text: "Everspace 2", QuestionID: question14.ID, Correct: false},
	}
	for _, opt := range options14 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question14.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question15 : Minimap Fortnite
	image15Search := entity.Image{Path: "/image15.png"}
	image15 := image15Search
	err = db.FirstOrCreate(&image15, image15Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 15: %v", err)
	}

	question15Search := entity.Question{ImageID: image15.ID}
	question15 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image15.ID,
	}
	err = db.FirstOrCreate(&question15, question15Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 15: %v", err)
	}

	options15 := []entity.Option{
		{Text: "H1Z1", QuestionID: question15.ID, Correct: false},
		{Text: "Call of Duty : Warzone", QuestionID: question15.ID, Correct: false},
		{Text: "PUBG", QuestionID: question15.ID, Correct: false},
		{Text: "Fortnite", QuestionID: question15.ID, Correct: true},
	}
	for _, opt := range options15 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question15.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question16 : Minimap Monster Hunter World
	image16Search := entity.Image{Path: "/image16.png"}
	image16 := image16Search
	err = db.FirstOrCreate(&image16, image16Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 16: %v", err)
	}

	question16Search := entity.Question{ImageID: image16.ID}
	question16 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image16.ID,
	}
	err = db.FirstOrCreate(&question16, question16Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 16: %v", err)
	}

	options16 := []entity.Option{
		{Text: "ARPG Frontier", QuestionID: question16.ID, Correct: false},
		{Text: "Dauntless", QuestionID: question16.ID, Correct: false},
		{Text: "Monster Hunter World", QuestionID: question16.ID, Correct: true},
		{Text: "Dragon's Dogma 2", QuestionID: question16.ID, Correct: false},
	}
	for _, opt := range options16 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question16.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question17 : Minimap RDR2
	image17Search := entity.Image{Path: "/image17.png"}
	image17 := image17Search
	err = db.FirstOrCreate(&image17, image17Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 17: %v", err)
	}

	question17Search := entity.Question{ImageID: image17.ID}
	question17 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image17.ID,
	}
	err = db.FirstOrCreate(&question17, question17Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 17: %v", err)
	}

	options17 := []entity.Option{
		{Text: "Horizon Forbidden West", QuestionID: question17.ID, Correct: false},
		{Text: "Days Gone", QuestionID: question17.ID, Correct: false},
		{Text: "Ghost of Tsushima", QuestionID: question17.ID, Correct: false},
		{Text: "Red Dead Redemption 2", QuestionID: question17.ID, Correct: true},
	}
	for _, opt := range options17 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question17.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question18 : Minimap Zelda Breath of the Wild
	image18Search := entity.Image{Path: "/image18.png"}
	image18 := image18Search
	err = db.FirstOrCreate(&image18, image18Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 18: %v", err)
	}

	question18Search := entity.Question{ImageID: image18.ID}
	question18 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image18.ID,
	}
	err = db.FirstOrCreate(&question18, question18Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 18: %v", err)
	}

	options18 := []entity.Option{
		{Text: "The Elder Scrolls V: Skyrim", QuestionID: question18.ID, Correct: true},
		{Text: "Kingdom Come: Deliverance", QuestionID: question18.ID, Correct: false},
		{Text: "Dragon Age: Inquisition", QuestionID: question18.ID, Correct: false},
		{Text: "Zelda Breath of the Wild", QuestionID: question18.ID, Correct: true},
	}
	for _, opt := range options18 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question18.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question19 : Minimap For Honor
	image19Search := entity.Image{Path: "/image19.png"}
	image19 := image19Search
	err = db.FirstOrCreate(&image19, image19Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 19: %v", err)
	}

	question19Search := entity.Question{ImageID: image19.ID}
	question19 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image19.ID,
	}
	err = db.FirstOrCreate(&question19, question19Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 19: %v", err)
	}

	options19 := []entity.Option{
		{Text: "For Honor", QuestionID: question19.ID, Correct: true},
		{Text: "Chivalry 2", QuestionID: question19.ID, Correct: false},
		{Text: "Mordhau", QuestionID: question19.ID, Correct: false},
		{Text: "Warlord Arena", QuestionID: question19.ID, Correct: false},
	}
	for _, opt := range options19 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question19.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question20 : Minimap Ghost Recon Wildlands
	image20Search := entity.Image{Path: "/image20.png"}
	image20 := image20Search
	err = db.FirstOrCreate(&image20, image20Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 20: %v", err)
	}

	question20Search := entity.Question{ImageID: image20.ID}
	question20 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image20.ID,
	}
	err = db.FirstOrCreate(&question20, question20Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 20: %v", err)
	}

	options20 := []entity.Option{
		{Text: "Metal Gear Solid V", QuestionID: question20.ID, Correct: false},
		{Text: "The Division 2", QuestionID: question20.ID, Correct: false},
		{Text: "Ghost Recon Wildlands", QuestionID: question20.ID, Correct: true},
		{Text: "Sniper Elite 5", QuestionID: question20.ID, Correct: false},
	}
	for _, opt := range options20 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question20.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question21 : Minimap Hitman 2
	image21Search := entity.Image{Path: "/image21.png"}
	image21 := image21Search
	err = db.FirstOrCreate(&image21, image21Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 21: %v", err)
	}

	question21Search := entity.Question{ImageID: image21.ID}
	question21 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image21.ID,
	}
	err = db.FirstOrCreate(&question21, question21Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 21: %v", err)
	}

	options21 := []entity.Option{
		{Text: "Subnautica", QuestionID: question21.ID, Correct: false},
		{Text: "Hitman 2", QuestionID: question21.ID, Correct: true},
		{Text: "Prince of Persia : The Sands of Time", QuestionID: question21.ID, Correct: false},
		{Text: "The Long Dark", QuestionID: question21.ID, Correct: false},
	}
	for _, opt := range options21 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question21.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question22 : Minimap World of Warcraft
	image22Search := entity.Image{Path: "/image22.png"}
	image22 := image22Search
	err = db.FirstOrCreate(&image22, image22Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 22: %v", err)
	}

	question22Search := entity.Question{ImageID: image22.ID}
	question22 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image22.ID,
	}
	err = db.FirstOrCreate(&question22, question22Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 22: %v", err)
	}

	options22 := []entity.Option{
		{Text: "Guild Wars 2", QuestionID: question22.ID, Correct: false},
		{Text: "Black Desert Online", QuestionID: question22.ID, Correct: false},
		{Text: "World of Warcraft", QuestionID: question22.ID, Correct: true},
		{Text: "New World", QuestionID: question22.ID, Correct: false},
	}
	for _, opt := range options22 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question22.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question23 : Minimap Forza Horizon 5
	image23Search := entity.Image{Path: "/image23.png"}
	image23 := image23Search
	err = db.FirstOrCreate(&image23, image23Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 23: %v", err)
	}

	question23Search := entity.Question{ImageID: image23.ID}
	question23 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image23.ID,
	}
	err = db.FirstOrCreate(&question23, question23Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 23: %v", err)
	}

	options23 := []entity.Option{
		{Text: "Forza Horizon 5", QuestionID: question23.ID, Correct: true},
		{Text: "Need for Speed : Heat", QuestionID: question23.ID, Correct: false},
		{Text: "The Crew 2", QuestionID: question23.ID, Correct: false},
		{Text: "Project Cars 3", QuestionID: question23.ID, Correct: false},
	}
	for _, opt := range options23 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question23.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question24 : Minimap Diablo IV
	image24Search := entity.Image{Path: "/image24.png"}
	image24 := image24Search
	err = db.FirstOrCreate(&image24, image24Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 24: %v", err)
	}

	question24Search := entity.Question{ImageID: image24.ID}
	question24 := entity.Question{
		Content: "Dans quel jeu est utilisée cette minimap ?",
		ImageID: image24.ID,
	}
	err = db.FirstOrCreate(&question24, question24Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 24: %v", err)
	}

	options24 := []entity.Option{
		{Text: "Diablo IV", QuestionID: question24.ID, Correct: true},
		{Text: "Path of Exile", QuestionID: question24.ID, Correct: false},
		{Text: "Lost Ark", QuestionID: question24.ID, Correct: false},
		{Text: "Torchlight Infinite", QuestionID: question24.ID, Correct: false},
	}
	for _, opt := range options24 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question24.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question25 : Interface The Witcher 3
	image25Search := entity.Image{Path: "/image25.png"}
	image25 := image25Search
	err = db.FirstOrCreate(&image25, image25Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 25: %v", err)
	}

	question25Search := entity.Question{ImageID: image25.ID}
	question25 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image25.ID,
	}
	err = db.FirstOrCreate(&question25, question25Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 25: %v", err)
	}

	options25 := []entity.Option{
		{Text: "The Witcher 3", QuestionID: question25.ID, Correct: true},
		{Text: "Dragon Age II", QuestionID: question25.ID, Correct: false},
		{Text: "Greedfall", QuestionID: question25.ID, Correct: false},
		{Text: "Fable Anniversary", QuestionID: question25.ID, Correct: false},
	}
	for _, opt := range options25 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question25.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question26 : Interface Outer Wilds
	image26Search := entity.Image{Path: "/image26.png"}
	image26 := image26Search
	err = db.FirstOrCreate(&image26, image26Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 26: %v", err)
	}

	question26Search := entity.Question{ImageID: image26.ID}
	question26 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image26.ID,
	}
	err = db.FirstOrCreate(&question26, question26Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 26: %v", err)
	}

	options26 := []entity.Option{
		{Text: "Mass Effect 2", QuestionID: question26.ID, Correct: false},
		{Text: "Prey", QuestionID: question26.ID, Correct: false},
		{Text: "Outer Wilds", QuestionID: question26.ID, Correct: true},
		{Text: "Deus Ex: Mankind Divided", QuestionID: question26.ID, Correct: false},
	}
	for _, opt := range options26 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question26.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question27 : Interface Fallout 4
	image27Search := entity.Image{Path: "/image27.png"}
	image27 := image27Search
	err = db.FirstOrCreate(&image27, image27Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 27: %v", err)
	}

	question27Search := entity.Question{ImageID: image27.ID}
	question27 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image27.ID,
	}
	err = db.FirstOrCreate(&question27, question27Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 27: %v", err)
	}

	options27 := []entity.Option{
		{Text: "Fallout 4", QuestionID: question27.ID, Correct: true},
		{Text: "Metro Exodus", QuestionID: question27.ID, Correct: false},
		{Text: "Stalker 2", QuestionID: question27.ID, Correct: false},
		{Text: "The Outer Worlds", QuestionID: question27.ID, Correct: false},
	}
	for _, opt := range options27 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question27.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question28 : Interface Apex Legends
	image28Search := entity.Image{Path: "/image28.png"}
	image28 := image28Search
	err = db.FirstOrCreate(&image28, image28Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 28: %v", err)
	}

	question28Search := entity.Question{ImageID: image28.ID}
	question28 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image28.ID,
	}
	err = db.FirstOrCreate(&question28, question28Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 28: %v", err)
	}

	options28 := []entity.Option{
		{Text: "Apex Legends", QuestionID: question28.ID, Correct: true},
		{Text: "Valorant", QuestionID: question28.ID, Correct: false},
		{Text: "Overwatch 2", QuestionID: question28.ID, Correct: false},
		{Text: "PUBG", QuestionID: question28.ID, Correct: false},
	}
	for _, opt := range options28 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question28.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question29 : Interface Back 4 Blood
	image29Search := entity.Image{Path: "/image29.png"}
	image29 := image29Search
	err = db.FirstOrCreate(&image29, image29Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 29: %v", err)
	}

	question29Search := entity.Question{ImageID: image29.ID}
	question29 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image29.ID,
	}
	err = db.FirstOrCreate(&question29, question29Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 29: %v", err)
	}

	options29 := []entity.Option{
		{Text: "World War Z", QuestionID: question29.ID, Correct: false},
		{Text: "Left 4 Dead 2", QuestionID: question29.ID, Correct: false},
		{Text: "Back 4 Blood", QuestionID: question29.ID, Correct: true},
		{Text: "Killing Floor 2", QuestionID: question29.ID, Correct: false},
	}
	for _, opt := range options29 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question29.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question30 : Interface Bioshock Infinite
	image30Search := entity.Image{Path: "/image30.png"}
	image30 := image30Search
	err = db.FirstOrCreate(&image30, image30Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 30: %v", err)
	}

	question30Search := entity.Question{ImageID: image30.ID}
	question30 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image30.ID,
	}
	err = db.FirstOrCreate(&question30, question30Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 30: %v", err)
	}

	options30 := []entity.Option{
		{Text: "Singularity", QuestionID: question30.ID, Correct: false},
		{Text: "Dishonored", QuestionID: question30.ID, Correct: false},
		{Text: "Prey", QuestionID: question30.ID, Correct: false},
		{Text: "Bioshock Infinite", QuestionID: question30.ID, Correct: true},
	}
	for _, opt := range options30 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question30.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question31 : Interface Chivalry 2
	image31Search := entity.Image{Path: "/image31.png"}
	image31 := image31Search
	err = db.FirstOrCreate(&image31, image31Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 31: %v", err)
	}

	question31Search := entity.Question{ImageID: image31.ID}
	question31 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image31.ID,
	}
	err = db.FirstOrCreate(&question31, question31Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 31: %v", err)
	}

	options31 := []entity.Option{
		{Text: "For Honor", QuestionID: question31.ID, Correct: false},
		{Text: "Chivalry 2", QuestionID: question31.ID, Correct: true},
		{Text: "Mordhau", QuestionID: question31.ID, Correct: false},
		{Text: "Kingdoms of Amalur", QuestionID: question31.ID, Correct: false},
	}
	for _, opt := range options31 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question31.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question32 : Interface Clair Obscur
	image32Search := entity.Image{Path: "/image32.png"}
	image32 := image32Search
	err = db.FirstOrCreate(&image32, image32Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 32: %v", err)
	}

	question32Search := entity.Question{ImageID: image32.ID}
	question32 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image32.ID,
	}
	err = db.FirstOrCreate(&question32, question32Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 32: %v", err)
	}

	options32 := []entity.Option{
		{Text: "Clair Obscur", QuestionID: question32.ID, Correct: true},
		{Text: "Sea of Stars", QuestionID: question32.ID, Correct: false},
		{Text: "Octopath Traveler 2", QuestionID: question32.ID, Correct: false},
		{Text: "Child of Light", QuestionID: question32.ID, Correct: false},
	}
	for _, opt := range options32 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question32.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question33 : Interface Conan Exiles
	image33Search := entity.Image{Path: "/image33.png"}
	image33 := image33Search
	err = db.FirstOrCreate(&image33, image33Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 33: %v", err)
	}

	question33Search := entity.Question{ImageID: image33.ID}
	question33 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image33.ID,
	}
	err = db.FirstOrCreate(&question33, question33Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 33: %v", err)
	}

	options33 := []entity.Option{
		{Text: "Ark: Survival Evolved", QuestionID: question33.ID, Correct: false},
		{Text: "Conan Exiles", QuestionID: question33.ID, Correct: true},
		{Text: "Rust", QuestionID: question33.ID, Correct: false},
		{Text: "Valheim", QuestionID: question33.ID, Correct: false},
	}
	for _, opt := range options33 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question33.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question34 : Interface Dead By Daylight
	image34Search := entity.Image{Path: "/image34.png"}
	image34 := image34Search
	err = db.FirstOrCreate(&image34, image34Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 34: %v", err)
	}

	question34Search := entity.Question{ImageID: image34.ID}
	question34 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image34.ID,
	}
	err = db.FirstOrCreate(&question34, question34Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 34: %v", err)
	}

	options34 := []entity.Option{
		{Text: "Dead By Daylight", QuestionID: question34.ID, Correct: true},
		{Text: "Friday the 13th", QuestionID: question34.ID, Correct: false},
		{Text: "The Texas Chain Saw Massacre", QuestionID: question34.ID, Correct: false},
		{Text: "Evolve", QuestionID: question34.ID, Correct: false},
	}
	for _, opt := range options34 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question34.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question35 : Interface Dead Space
	image35Search := entity.Image{Path: "/image35.png"}
	image35 := image35Search
	err = db.FirstOrCreate(&image35, image35Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 35: %v", err)
	}

	question35Search := entity.Question{ImageID: image35.ID}
	question35 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image35.ID,
	}
	err = db.FirstOrCreate(&question35, question35Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 35: %v", err)
	}

	options35 := []entity.Option{
		{Text: "Prey", QuestionID: question35.ID, Correct: false},
		{Text: "The Callisto Protocol", QuestionID: question35.ID, Correct: false},
		{Text: "Alien Isolation", QuestionID: question35.ID, Correct: false},
		{Text: "Dead Space", QuestionID: question35.ID, Correct: true},
	}
	for _, opt := range options35 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question35.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question36 : Interface Final Fantasy XV
	image36Search := entity.Image{Path: "/image36.png"}
	image36 := image36Search
	err = db.FirstOrCreate(&image36, image36Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 36: %v", err)
	}

	question36Search := entity.Question{ImageID: image36.ID}
	question36 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image36.ID,
	}
	err = db.FirstOrCreate(&question36, question36Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 36: %v", err)
	}

	options36 := []entity.Option{
		{Text: "Final Fantasy XV", QuestionID: question36.ID, Correct: true},
		{Text: "Final Fantasy VII Remake", QuestionID: question36.ID, Correct: false},
		{Text: "Tales of Arise", QuestionID: question36.ID, Correct: false},
		{Text: "Xenoblade Chronicles 3", QuestionID: question36.ID, Correct: false},
	}
	for _, opt := range options36 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question36.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question37 : Interface Kingdom Come Deliverance
	image37Search := entity.Image{Path: "/image37.png"}
	image37 := image37Search
	err = db.FirstOrCreate(&image37, image37Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 37: %v", err)
	}

	question37Search := entity.Question{ImageID: image37.ID}
	question37 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image37.ID,
	}
	err = db.FirstOrCreate(&question37, question37Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 37: %v", err)
	}

	options37 := []entity.Option{
		{Text: "Greedfall", QuestionID: question37.ID, Correct: false},
		{Text: "Skyrim", QuestionID: question37.ID, Correct: false},
		{Text: "Kingdom Come Deliverance", QuestionID: question37.ID, Correct: true},
		{Text: "Mount & Blade II", QuestionID: question37.ID, Correct: false},
	}
	for _, opt := range options37 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question37.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question38 : Interface Prey
	image38Search := entity.Image{Path: "/image38.png"}
	image38 := image38Search
	err = db.FirstOrCreate(&image38, image38Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 38: %v", err)
	}

	question38Search := entity.Question{ImageID: image38.ID}
	question38 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image38.ID,
	}
	err = db.FirstOrCreate(&question38, question38Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 38: %v", err)
	}

	options38 := []entity.Option{
		{Text: "The Outer Worlds", QuestionID: question38.ID, Correct: false},
		{Text: "System Shock", QuestionID: question38.ID, Correct: false},
		{Text: "Deus Ex: Mankind Divided", QuestionID: question38.ID, Correct: false},
		{Text: "Prey", QuestionID: question38.ID, Correct: true},
	}
	for _, opt := range options38 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question38.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question39 : Interface No Man's Sky
	image39Search := entity.Image{Path: "/image39.png"}
	image39 := image39Search
	err = db.FirstOrCreate(&image39, image39Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 39: %v", err)
	}

	question39Search := entity.Question{ImageID: image39.ID}
	question39 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image39.ID,
	}
	err = db.FirstOrCreate(&question39, question39Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 39: %v", err)
	}

	options39 := []entity.Option{
		{Text: "No Man's Sky", QuestionID: question39.ID, Correct: true},
		{Text: "Starfield", QuestionID: question39.ID, Correct: false},
		{Text: "Elite Dangerous", QuestionID: question39.ID, Correct: false},
		{Text: "Outer Wilds", QuestionID: question39.ID, Correct: false},
	}
	for _, opt := range options39 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question39.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question40 : Interface Subnautica Below Zero
	image40Search := entity.Image{Path: "/image40.png"}
	image40 := image40Search
	err = db.FirstOrCreate(&image40, image40Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 40: %v", err)
	}

	question40Search := entity.Question{ImageID: image40.ID}
	question40 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image40.ID,
	}
	err = db.FirstOrCreate(&question40, question40Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 40: %v", err)
	}

	options40 := []entity.Option{
		{Text: "ABZU", QuestionID: question40.ID, Correct: false},
		{Text: "Subnautica Below Zero", QuestionID: question40.ID, Correct: true},
		{Text: "Raft", QuestionID: question40.ID, Correct: false},
		{Text: "The Long Dark", QuestionID: question40.ID, Correct: false},
	}
	for _, opt := range options40 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question40.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question41 : Interface The Binding of Isaac
	image41Search := entity.Image{Path: "/image41.png"}
	image41 := image41Search
	err = db.FirstOrCreate(&image41, image41Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 41: %v", err)
	}

	question41Search := entity.Question{ImageID: image41.ID}
	question41 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image41.ID,
	}
	err = db.FirstOrCreate(&question41, question41Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 41: %v", err)
	}

	options41 := []entity.Option{
		{Text: "The Binding of Isaac", QuestionID: question41.ID, Correct: true},
		{Text: "Hades", QuestionID: question41.ID, Correct: false},
		{Text: "Cult of the Lamb", QuestionID: question41.ID, Correct: false},
		{Text: "Enter the Gungeon", QuestionID: question41.ID, Correct: false},
	}
	for _, opt := range options41 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question41.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	//Question42 : Interface Dark Souls 3
	image42Search := entity.Image{Path: "/image42.png"}
	image42 := image42Search
	err = db.FirstOrCreate(&image42, image42Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Image 42: %v", err)
	}

	question42Search := entity.Question{ImageID: image42.ID}
	question42 := entity.Question{
		Content: "De quel jeu viens cette interface de personnage ?",
		ImageID: image42.ID,
	}
	err = db.FirstOrCreate(&question42, question42Search).Error
	if err != nil {
		log.Fatalf("Erreur seed Question 42: %v", err)
	}

	options42 := []entity.Option{
		{Text: "Bloodborne", QuestionID: question42.ID, Correct: false},
		{Text: "Demon's Souls", QuestionID: question42.ID, Correct: false},
		{Text: "Dark Souls III", QuestionID: question42.ID, Correct: true},
		{Text: "Elden Ring", QuestionID: question42.ID, Correct: false},
	}
	for _, opt := range options42 {
		searchOpt := entity.Option{Text: opt.Text, QuestionID: question42.ID}
		err := db.FirstOrCreate(&opt, searchOpt).Error
		if err != nil {
			log.Fatalf("Erreur seed Option '%s': %v", opt.Text, err)
		}
	}

	fmt.Println("Default data successfully inserted")
}
