package main

import (
	"context"
	"fmt"
	"time"

	"github.com/a-berahman/pc-offline-challenge/common"
	"github.com/a-berahman/pc-offline-challenge/repository"
	"github.com/avast/retry-go"
	"golang.org/x/text/language"
)

// Translate represents basic property for Translate service
type Translate struct {
	translator Translator
	cache      repository.Cacher
}

// NewTranslate returns new instance of Translate service
func NewTranslate() *Translate {
	t := newRandomTranslator(
		100*time.Millisecond,
		500*time.Millisecond,
		0.1,
	)
	return &Translate{translator: t, cache: repository.GetRepository(common.CacheRepo).(repository.Cacher)}

}

//Translate returns translate value as a string
func (t *Translate) Translate(ctx context.Context, from, to language.Tag, data string) (string, error) {
	key := fmt.Sprintf("%v::%v::%v", from, to, data)
	//if it founds the same queries it returns the result from cache
	if cacheRes := t.cache.Get(key); cacheRes != "" {
		return cacheRes, nil
	}
	//otherwise, it sends request to 3rd party
	res, err := t.doTranslate(ctx, from, to, data)
	if err != nil {
		return "", err
	}
	//then, write queries in the cache
	t.cache.Write(key, res)

	return res, nil

}
func (t *Translate) doTranslate(ctx context.Context, from, to language.Tag, data string) (string, error) {
	var res string
	errR := retry.Do(
		func() error {
			var err error
			res, err = t.translator.Translate(ctx, from, to, data)
			if err != nil {
				return err
			}
			return nil
		},
	)
	if errR != nil {
		return res, errR
	}
	return res, nil
}
