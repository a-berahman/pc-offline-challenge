//Package repository is a bunch of repository interface that support needed function for repository actions
package repository

import (
	"github.com/a-berahman/pc-offline-challenge/common"
	"github.com/a-berahman/pc-offline-challenge/repository/cache"
)

// GetRepository returns new repository instance by repository enum
// - solution is implemented by Factory design pattern
func GetRepository(repoConst int) interface{} {
	switch repoConst {
	case common.CacheRepo:
		return cache.New()
	default:
		return nil
	}

}
