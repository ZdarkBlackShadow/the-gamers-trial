package entity

import "time"

type Image struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Path      string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
