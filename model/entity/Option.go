package entity

import (
	"time"
)

type Option struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	Text       string    `gorm:"size:100;not null"`
	Correct    bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	QuestionID int       `gorm:"not null"`
	Question   Question  `gorm:"foreignKey:QuestionID"`
}
