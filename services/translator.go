package services

import "golang.org/x/text/language"

//Translator is implemented by objects that promote Cache repository
type Translator interface {
	//Translate returns translate value as a string
	Translate(from, to language.Tag, data string) (string, error)
}
