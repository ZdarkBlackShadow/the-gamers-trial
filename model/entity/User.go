package entity

import "time"

type User struct {
	ID        string    `gorm:"primaryKey;type:char(36);not null"`
	Pseudo    string    `gorm:"size:20;not null"`
	Email     string    `gorm:"size:50;not null"`
	Password  string    `gorm:"size:150;not null"`
	Score     int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
