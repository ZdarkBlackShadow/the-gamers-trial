package service

import (
	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/request"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/views"
	"github.com/ZdarkBlackShadow/the-gamers-trial/repository"
	"github.com/ZdarkBlackShadow/the-gamers-trial/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthentificationService struct {
	userRepo *repository.UserRepository
}

func InitAuthentificationService(db *gorm.DB) *AuthentificationService {
	return &AuthentificationService{
		userRepo: repository.InitUserRepository(db),
	}
}

func (s *AuthentificationService) Register(req request.Register) (bool, views.Error, entity.User) {
	if req.Pseudo == "" {
		return true, util.CreateError(fiber.ErrBadRequest.Message, "pseudo cannot be empty"), entity.User{}
	}

	if len(req.Pseudo) > 20 {
		return true, util.CreateError(fiber.ErrBadRequest.Message, "pseudo is too long(max 20 character)"), entity.User{}
	}

	if !util.IsValidEmail(req.Email) {
		return true, util.CreateError(fiber.ErrBadRequest.Message, "invalid email"), entity.User{}
	}

	if req.Password != req.PasswordConfirmation {
		return true, util.CreateError(fiber.ErrBadRequest.Message, "passwords not match"), entity.User{}
	}

	err := util.IsPasswordStrong(req.Password)
	if err != nil {
		return true, util.CreateError(fiber.ErrBadRequest.Message, err.Error()), entity.User{}
	}

	userID, err := util.GenerateUUID()
	if err != nil {
		config.Log.Error("error when generate UUID:" + err.Error())
		return true, util.CreateError(fiber.ErrInternalServerError.Message, "error when create user"), entity.User{}
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		config.Log.Error("error when trying to hash password:" + err.Error())
		return true, util.CreateError(fiber.ErrInternalServerError.Message, "error when create user"), entity.User{}
	}

	userEntity := entity.User{
		ID:       userID,
		Pseudo:   req.Pseudo,
		Email:    req.Email,
		Score:    0,
		Password: hashedPassword,
	}

	userCreated, err := s.userRepo.AddUser(userEntity)
	if err != nil {
		config.Log.Error("error when create the user in database:" + err.Error())
		return true, util.CreateError(fiber.ErrInternalServerError.Message, "error when create user"), entity.User{}
	}
	return false, views.Error{}, userCreated
}

func (s *AuthentificationService) Login(req request.Login) (bool, views.Error, entity.User) {
	user, err := s.userRepo.GetUserByPseudoOrEmail(req.EmailOrPseudo)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true, util.CreateError(fiber.ErrBadRequest.Message, "email or pseudo not found"), entity.User{}
		} else {
			config.Log.Error("error with database:" + err.Error())
			return true, util.CreateError(fiber.ErrInternalServerError.Message, "error when analysing data"), entity.User{}
		}
	}

	valid, err := util.VerifyPassword(req.Password, user.Password)
	if err != nil {
		config.Log.Error("error when comparing password:" + err.Error())
		return true, util.CreateError(fiber.ErrInternalServerError.Message, "error when analysing data"), entity.User{}
	}

	if !valid {
		return true, util.CreateError(fiber.ErrBadRequest.Message, "incorrect password"), entity.User{}
	}
	return false, views.Error{}, user
}
