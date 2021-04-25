package main

import (
	"context"
	"testing"

	"golang.org/x/text/language"
)

func Test_Translate_Success(t *testing.T) {

	curreService := NewTranslate()
	_, err := curreService.Translate(context.Background(), language.English, language.Japanese, "test")
	if err != nil {
		t.Fatalf("expected error to be nil but error is %v", err)
	}

}
