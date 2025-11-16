package service

import (
	"errors"
	"strconv"

	"github.com/ZdarkBlackShadow/the-gamers-trial/config"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/entity"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/request"
	"github.com/ZdarkBlackShadow/the-gamers-trial/model/session"
	"github.com/ZdarkBlackShadow/the-gamers-trial/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type QuestionService struct {
	optionRepo   *repository.OptionRepository
	questionRepo *repository.QuestionRepository
	userRepo     *repository.UserRepository
}

func InitQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{
		optionRepo:   repository.InitOptionRepository(db),
		questionRepo: repository.InitQuestionRepository(db),
		userRepo:     repository.InitUserRepository(db),
	}
}

func (s *QuestionService) ResponseAnalysis(req request.Question, user entity.User) (string, session.UserAnswer, entity.User, error) {
	sessionData := session.UserAnswer{}
	option, err := s.optionRepo.GetOptionByID(uint(req.UserAnswerID))
	if err != nil {
		config.Log.Error("error when analyse option ID: " + err.Error())
		return fiber.ErrBadRequest.Message, session.UserAnswer{}, entity.User{}, err
	}

	if option.QuestionID != req.QuestionID {
		return fiber.ErrBadRequest.Message, session.UserAnswer{}, entity.User{}, errors.New("the id of question and option does not match")
	}

	options, err := s.optionRepo.GetAllOptionsByQuestionId(uint(req.QuestionID))
	if err != nil {
		config.Log.Error("error when getting all options: " + err.Error())
		return fiber.ErrInternalServerError.Message, session.UserAnswer{}, entity.User{}, err
	}
	
	sessionData.OldUserScore = user.Score
	var updatedUser entity.User = user
	
	for _, op := range options {
		if op.Correct {
			if op.ID == option.ID {
				sessionData.UserHasCorrect = true
				sessionData.GoodOption = op.Text
				user.Score += 500
				updatedUser, err = s.userRepo.UpdateUser(user)
				if err != nil {
					config.Log.Error("error when updated user:" + err.Error())
					return fiber.ErrInternalServerError.Message, session.UserAnswer{}, entity.User{}, err
				}
				sessionData.NewUserScore = updatedUser.Score
				return "", sessionData, updatedUser, nil
			}
		}
	}
	
	// Réponse incorrecte
	sessionData.UserHasCorrect = false
	sessionData.NewUserScore = user.Score
	return "", sessionData, updatedUser, nil
}

func (s *QuestionService) GetRandomQuestion() ([]entity.Option, entity.Question, error) {
	numberOfQuestions, err := s.questionRepo.GetNumberOfQuestion()
	if err != nil {
		config.Log.Error("error when trying to get number of question: " + err.Error())
		return []entity.Option{}, entity.Question{}, err
	}
	config.Log.Debug("number of questions: " + strconv.Itoa(numberOfQuestions))

	if numberOfQuestions == 0 {
		config.Log.Error("no question found in the database")
		return []entity.Option{}, entity.Question{}, errors.New("no questions available")
	}

	randomIndex := config.Rand.Intn(numberOfQuestions) + 1

	question, err := s.questionRepo.GetQuestionById(uint(randomIndex))
	if err != nil {
		config.Log.Error("error when trying to get random question: " + err.Error())
		return []entity.Option{}, entity.Question{}, err
	}

	options, err := s.optionRepo.GetAllOptionsByQuestionId(uint(randomIndex))
	if err != nil {
		config.Log.Error("error when trying to get all the options: " + err.Error())
		return []entity.Option{}, entity.Question{}, err
	}

	return options, question, nil
}
