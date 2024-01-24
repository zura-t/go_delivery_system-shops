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

func (uc *ShopUseCase) UpdateShop(id int64, req *entity.UpdateShopInfo) (*entity.Shop, int, error) {
	arg := db.UpdateShopParams{
		ID: id,
		Name: sql.NullString{
			String: req.Name,
			Valid:  req.Name != "",
		},
		Description: sql.NullString{
			String: req.Description,
			Valid:  req.Description != "",
		},
		OpenTime: sql.NullTime{
			Time:  req.OpenTime,
			Valid: !req.CloseTime.IsZero(),
		},
		CloseTime: sql.NullTime{
			Time:  req.CloseTime,
			Valid: !req.CloseTime.IsZero(),
		},
		IsClosed: sql.NullBool{
			Bool:  req.IsClosed,
			Valid: true,
		},
	}

	shopUpdated, err := uc.store.UpdateShop(context.Background(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("shop not found: %s", err)
			return nil, http.StatusNotFound, err
		}
		err = fmt.Errorf("failed to update shop: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := convertShop(shopUpdated)

	return res, http.StatusOK, nil
}
