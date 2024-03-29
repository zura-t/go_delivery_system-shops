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

func (uc *ShopUseCase) UpdateMenuItem(id int64, req *entity.UpdateMenuItem) (*entity.GetMenuItem, int, error) {
	shop, err := uc.store.GetShopWithMenuItemId(context.Background(), id)
	if shop.UserID != req.UserId {
		err = fmt.Errorf("failed to update menu item: %s", err)
		return nil, http.StatusLocked, err
	}

	arg := db.UpdateMenuItemParams{
		ID: id,
		Name: sql.NullString{
			String: req.Name,
			Valid:  req.Name != "",
		},
		Description: sql.NullString{
			String: req.Description,
			Valid:  req.Description != "",
		},
		Price: sql.NullInt32{
			Int32: req.Price,
			Valid: req.Price != 0,
		},
	}

	menuItemUpdated, err := uc.store.UpdateMenuItem(context.Background(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			err = fmt.Errorf("menu item not found: %s", err)
			return nil, http.StatusNotFound, err
		}
		err = fmt.Errorf("failed to update menu item: %s", err)
		return nil, http.StatusInternalServerError, err
	}

	res := convertMenu(menuItemUpdated)

	return res, http.StatusOK, nil
}
