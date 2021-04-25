package main

import (
	"context"
	"testing"

	"golang.org/x/text/language"
)

func Test_Translate(t *testing.T) {
	s := NewService()
	specs := []struct {
		word string
	}{
		{word: "test1"},
		{word: "test2"},
		{word: "test3"},
		{word: "test1"},
		{word: "test1"},
	}
	first, err := s.translator.Translate(context.Background(), language.English, language.Japanese, specs[0].word)
	if err != nil {
		t.Fatalf("expected error to be nil but %v", err)
	}
	for _, spec := range specs {
		got, err := s.translator.Translate(context.Background(), language.English, language.Japanese, spec.word)
		if err != nil {
			t.Fatalf("expected error to be nil but %v", err)
		}
		if spec.word == "test1" && got != first {
			t.Fatalf("expected got to be %v but %v", got, first)
		}

	}

}
