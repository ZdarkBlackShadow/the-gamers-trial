package views

type Profile struct {
	HasError bool
	ErrorDetail    Error
	Pseudo   string
	Email    string
	Score    string
}
