package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/a-berahman/pc-offline-challenge/common"
	"github.com/a-berahman/pc-offline-challenge/services"
	"golang.org/x/text/language"
)

func main() {
	ctx := context.Background()
	rand.Seed(time.Now().UTC().UnixNano())
	translerSerivce := services.GetService(ctx, common.TranlatorService).(services.Translator)

	fmt.Println(translerSerivce.Translate(language.English, language.Japanese, "test"))
}
