package views

type Score struct {
	HasError       bool
	ErrorDetail   Error
	List []UserScore
}

type UserScore struct {
	Pseudo string
	Score int
}
