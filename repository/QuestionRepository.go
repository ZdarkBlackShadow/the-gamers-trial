package repository

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func InitQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{
		db: db,
	}
}

func (r *QuestionRepository) GetQuestionById(id uint) (entity.Question, error) {
	var question entity.Question
	err := r.db.Preload("Image").Where("id = ?", id).First(&question).Error
	if err != nil {
		return entity.Question{}, err
	}
	return question, nil
}

func (r *QuestionRepository) GetNumberOfQuestion() (int, error) {
	var count int64
	err := r.db.Model(&entity.Question{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
