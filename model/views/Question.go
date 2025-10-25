package views

type Question struct {
	State     bool
	ID        int
	Content   string
	HasImage  bool
	ImagePath string
	Options   []struct {
		ID   int
		Text string
	}
	UserHasCorrect bool
	OldUserScore   int
	NewUserScore   int
}
