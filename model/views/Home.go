package views

type Home struct {
	HasError    bool
	ErrorDetail Error
	IsLogin     bool
	UserInfo    struct {
		Name  string
		Score int
	}
}
