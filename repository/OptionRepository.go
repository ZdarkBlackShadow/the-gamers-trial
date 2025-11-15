package repository

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"gorm.io/gorm"
)

type OptionRepository struct {
	db *gorm.DB
}

func InitOptionRepository(db *gorm.DB) *OptionRepository {
	return &OptionRepository{
		db: db,
	}
}

func (r *OptionRepository) GetAllOptionsByQuestionId(questionID uint) ([]entity.Option, error) {
	var options []entity.Option
	err := r.db.Where("question_id = ?", questionID).Find(&options).Error
	if err != nil {
		return nil, err
	}
	return options, nil
}

func (r *OptionRepository) GetOptionByID(id uint) (entity.Option, error) {
	var option entity.Option
	err := r.db.Where("id = ?", id).First(&option).Error
	if err != nil {
		return entity.Option{}, err
	}
	return option, nil
}