package usecase

import (
	"github.com/zura-t/go_delivery_system-shops/config"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

type ShopUseCase struct {
	store      db.Store
	config     *config.Config
}

func New(store db.Store, config *config.Config) *ShopUseCase {
	return &ShopUseCase{
		store:      store,
		config:     config,
	}
}