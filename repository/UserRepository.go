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
    if err := r.db.Where("id = ?", user.ID).First(&createdUser).Error; err != nil {
        return entity.User{}, err
    }

    return createdUser, nil
}

func (r *UserRepository) GetUserByPseudoOrEmail(pseudoOrEmail string) (entity.User, error) {
    var user entity.User
    err := r.db.Where("email = ? OR pseudo = ?", pseudoOrEmail, pseudoOrEmail).First(&user).Error
    if err != nil {
        return entity.User{}, err
    }

    return user, nil
}

func (r *UserRepository) UpdateUser(user entity.User) (entity.User, error) {
    var existingUser entity.User
    err := r.db.First(&existingUser, user.ID).Error
    if err != nil {
        return entity.User{}, err
    }

    if err := r.db.Model(&existingUser).Updates(user).Error; err != nil {
        return entity.User{}, err
    }

    return existingUser, nil
}