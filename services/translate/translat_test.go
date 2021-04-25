package translate

import (
	"context"
	"testing"

	"golang.org/x/text/language"
)

func Test_Translate_Success(t *testing.T) {

	curreService := New(context.Background())
	got, err := curreService.Translate(language.English, language.Japanese, "test")
	if err != nil {
		t.Fatalf("expected error to be nil but error is %v", err)
	}
	if got != "en -> ja : test -> 6129484611666145821" {
		t.Fatalf("expected got to be en -> ja : test -> 6129484611666145821	but it is %v", got)
	}

}
