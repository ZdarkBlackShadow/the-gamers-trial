package session

type UserAnswer struct {
	UserHasCorrect bool
	OldUserScore   int
	NewUserScore   int
	GoodOption     string
	QuestionScore  int
}
