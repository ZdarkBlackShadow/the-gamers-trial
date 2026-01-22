package repository

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"gorm.io/gorm"
)

type ScoreHistoryRepository struct {
	db *gorm.DB
}

func InitScoreHistoryRepository(db *gorm.DB) *ScoreHistoryRepository {
	return &ScoreHistoryRepository{
		db: db,
	}
}

func (r *ScoreHistoryRepository) Create(entry entity.ScoreHistory) (entity.ScoreHistory, error) {
	if err := r.db.Create(&entry).Error; err != nil {
		return entity.ScoreHistory{}, err
	}
	return entry, nil
}

func (r *ScoreHistoryRepository) GetAllOrdered() ([]entity.ScoreHistory, error) {
	var entries []entity.ScoreHistory
	if err := r.db.Order("score DESC, created_at DESC").Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}




