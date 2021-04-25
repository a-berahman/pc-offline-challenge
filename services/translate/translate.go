package translate

import (
	"context"
	"time"

	"golang.org/x/text/language"
)

// Translate represents basic property for Translate service
type Translate struct {
	translator Translator
	ctx        context.Context
}

// New returns new instance of Translate service
func New(ctx context.Context) *Translate {

	t := newRandomTranslator(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)
	return &Translate{translator: t, ctx: ctx}

}

//Translate returns translate value as a string
func (t *Translate) Translate(from, to language.Tag, data string) (string, error) {
	return t.translator.Translate(t.ctx, from, to, data)
}
