package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/zura-t/go_delivery_system-shops/internal/entity"
	db "github.com/zura-t/go_delivery_system-shops/pkg/db/sqlc"
)

func convertShop(shop db.Shop) *entity.Shop {
	return &entity.Shop{
		ID:          shop.ID,
		Name:        shop.Name,
		Description: shop.Description.String,
		OpenTime:    shop.OpenTime.Time,
		CloseTime:   shop.CloseTime.Time,
		IsClosed:    shop.IsClosed,
		CreatedAt:   shop.CreatedAt,
	}
}

func (uc *ShopUseCase) CreateShop(req *entity.CreateShop) (*entity.Shop, int, error) {
	arg := db.CreateShopParams{
		Name: req.Name,
		OpenTime: sql.NullTime{
			Valid: true,
			Time:  req.OpenTime,
		},
		CloseTime: sql.NullTime{
			Valid: true,
			Time:  req.CloseTime,
		},
		IsClosed: req.IsClosed,
	}

	shopCreated, err := uc.store.CreateShop(context.Background(), arg)
	if err != nil {
		err = fmt.Errorf("failed to create shop: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := convertShop(shopCreated)

	return res, http.StatusOK, nil
}
