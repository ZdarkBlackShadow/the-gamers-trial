package entity

import "time"

type ScoreHistory struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    string    `gorm:"type:char(36);not null;index"`
	Pseudo    string    `gorm:"size:20;not null"`
	Score     int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}




