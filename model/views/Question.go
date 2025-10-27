package views

type Question struct {
	HasError    bool
	ErrorDetail Error
	State       bool
	ID          int
	Content     string
	HasImage    bool
	ImagePath   string
	Options     []struct {
		ID   int
		Text string
	}
	UserHasCorrect bool
	OldUserScore   int
	NewUserScore   int
}
