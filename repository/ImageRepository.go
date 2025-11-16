package repository

import (
    "github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
    "gorm.io/gorm"
)

type ImageRepository struct {
    db *gorm.DB
}

func InitImageRepository(db *gorm.DB) *ImageRepository {
    return &ImageRepository{
        db: db,
    }
}

func (r *ImageRepository) GetImageByPath(path string) (entity.Image, error) {
    var image entity.Image
    err := r.db.Where("path = ?", path).First(&image).Error 
    if err != nil {
        return entity.Image{}, err
    }
    return image, nil
}