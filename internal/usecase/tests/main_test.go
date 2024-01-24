package usecase_test

import (
	"testing"

	"github.com/zura-t/go_delivery_system-shops/config"
	"github.com/zura-t/go_delivery_system-shops/internal/usecase"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *usecase.ShopUseCase {
	config := &config.Config{}
	
	server := usecase.New(store, config)

	return server
}