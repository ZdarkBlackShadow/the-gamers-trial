package views

type Question struct {
	HasError       bool
	ErrorDetail    Error
	State          bool
	ID             int
	Content        string
	HasImage       bool
	ImagePath      string
	Options        []Option
	UserHasCorrect bool
	OldUserScore   int
	NewUserScore   int
	QuestionScore  int
}

type Option struct {
	ID   int
	Text string
}
