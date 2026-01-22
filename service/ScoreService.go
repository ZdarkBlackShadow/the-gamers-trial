package service

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/views"
	"github.com/ZdarkBlackShadow/the-gamers-trial/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ScoreService struct {
	userRepo          *repository.UserRepository
	scoreHistoryRepo  *repository.ScoreHistoryRepository
}

func InitScoreService(db *gorm.DB) *ScoreService {
	return &ScoreService{
		userRepo:         repository.InitUserRepository(db),
		scoreHistoryRepo: repository.InitScoreHistoryRepository(db),
	}
}

func (s *ScoreService) GetRanking() (string, []views.UserScore, error) {
	scoreEntries, err := s.scoreHistoryRepo.GetAllOrdered()
	if err != nil {
		config.Log.Error("error when trying to get score history: " + err.Error())
		return fiber.ErrInternalServerError.Message, []views.UserScore{}, err
	}
	var ranking []views.UserScore
	for _, entry := range scoreEntries {
		ranking = append(ranking, views.UserScore{Pseudo: entry.Pseudo, Score: entry.Score})
	}
	return "", ranking, nil
}

func (s *ScoreService) SaveScoreHistory(user entity.User) error {
	_, err := s.scoreHistoryRepo.Create(entity.ScoreHistory{
		UserID: user.ID,
		Pseudo: user.Pseudo,
		Score:  user.Score,
	})
	return err
}