package request

type Register struct {
	Email                string `json:"email"`
	Pseudo               string `json:"pseudo"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type Login struct {
	EmailOrPseudo string `json:"email_or_pseudo"`
	Password      string `json:"password"`
}
