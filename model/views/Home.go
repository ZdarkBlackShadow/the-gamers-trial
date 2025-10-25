package views

type Home struct {
	IsLogin  bool
	UserInfo struct {
		Name  string
		Score int
	}
}
