package entity

import "time"

type Question struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Content   string    `gorm:"size:300;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	ImageID   int       `gorm:"not null"`
	Image     Image     `gorm:"foreignKey:ImageID"`
}
