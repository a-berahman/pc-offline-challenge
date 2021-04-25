package main

// Service is a Translator user.
type Service struct {
	translator Translator
}

// NewService returns new instance of service
func NewService() *Service {

	return &Service{
		translator: NewTranslate(),
	}
}
