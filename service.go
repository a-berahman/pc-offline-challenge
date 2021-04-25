package main

// Service is a Translator user.
type Service struct {
	translator Translator
}

func NewService() *Service {

	return &Service{
		translator: NewTranslate(),
	}
}
