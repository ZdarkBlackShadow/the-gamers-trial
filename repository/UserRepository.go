package repository

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (r *UserRepository) AddUser(user entity.User) (entity.User, error) {
    if err := r.db.Create(&user).Error; err != nil {
        return entity.User{}, err
    }

    var createdUser entity.User
    if err := r.db.First(&createdUser, user.ID).Error; err != nil {
        return entity.User{}, err
    }

    return createdUser, nil
}
