package service

import "gorm.io/gorm"

type AuthentificationService struct {
	db *gorm.DB
}

func InitAuthentificationService(db *gorm.DB) *AuthentificationService {
	return &AuthentificationService{
		db: db,
	}
}
