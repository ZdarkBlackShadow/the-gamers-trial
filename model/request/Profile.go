package request

type Profile struct {
	OldPseudo string `json:"old_pseudo"`
	NewPseudo string `json:"new_pseudo"`
}