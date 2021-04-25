//Package services is a bunch of services interface that support needed function for service actions
package services

import (
	"context"

	"github.com/a-berahman/pc-offline-challenge/common"
	"github.com/a-berahman/pc-offline-challenge/services/translate"
)

// GetService returns new service instance by service enum
// - solution is implemented by Factory design pattern
func GetService(ctx context.Context, serviceConst int) interface{} {
	switch serviceConst {
	case common.TranlatorService:
		return translate.New(ctx)
	default:
		return nil
	}
}
