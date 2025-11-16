package service

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/views"
	"github.com/ZdarkBlackShadow/the-gamers-trial/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ScoreService struct {
	userRepo *repository.UserRepository
}

func InitScoreService(db *gorm.DB) *ScoreService {
	return &ScoreService{
		userRepo: repository.InitUserRepository(db),
	}
}

func (s *ScoreService) GetRanking() (string, []views.UserScore, error) {
	allUser, err := s.userRepo.GetAllUser()
	if err != nil {
		config.Log.Error("error when trying to get user and score: " + err.Error())
		return fiber.ErrInternalServerError.Message, []views.UserScore{}, err
	}
	var ranking []views.UserScore
	for _, data := range allUser {
		ranking = append(ranking, views.UserScore{Pseudo: data.Pseudo, Score: data.Score})
	}
	return "", ranking, nil
}