package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/repository"
	"gorm.io/gorm"
)

type ImageService struct {
	imageRepo *repository.ImageRepository
	basePath  string
}

func InitImageService(db *gorm.DB, baseDir string) *ImageService {
	return &ImageService{
		imageRepo: repository.InitImageRepository(db),
		basePath:  baseDir,
	}
}

func (s *ImageService) GetImagePathByDBPath(pathFromURL string) (string, error) {
	_, err := s.imageRepo.GetImageByPath(pathFromURL)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.Log.Debug(fmt.Sprintf("Image path '%s' not found in database.", pathFromURL))
			return "", errors.New("image not found")
		}
		config.Log.Error("Database error when checking image path: " + err.Error())
		return "", errors.New("internal server error")
	}

	cleanPath := pathFromURL
	if len(cleanPath) > 0 && cleanPath[0] == '/' {
		cleanPath = cleanPath[1:]
	}

	fullPath := filepath.Join(s.basePath, cleanPath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		config.Log.Error("File not found on disk: " + fullPath)
		return "", errors.New("image file not found")
	}

	return fullPath, nil
}
