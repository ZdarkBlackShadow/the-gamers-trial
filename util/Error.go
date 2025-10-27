package util

import "github.com/ZdarkBlackShadow/the-gamers-trial/model/views"

func CreateError(errType string, text string) views.Error {
	return views.Error{
		ErrorType: errType,
		ErrorText: text,
	}
}